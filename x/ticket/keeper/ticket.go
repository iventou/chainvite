package keeper

import (
	"context"

	"cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/iventou/chainvite/x/ticket/types"
)

// SetTicket set a specific ticket in the store from its index
func (k Keeper) SetTicket(ctx context.Context, ticket types.Ticket) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TicketKeyPrefix))
	b := k.cdc.MustMarshal(&ticket)
	store.Set(types.TicketKey(
		ticket.Index,
	), b)
}

// GetTicket returns a ticket from its index
func (k Keeper) GetTicket(
	ctx context.Context,
	index string,
) (val types.Ticket, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TicketKeyPrefix))

	b := store.Get(types.TicketKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTicket removes a ticket from the store
func (k Keeper) RemoveTicket(
	ctx context.Context,
	index string,
) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TicketKeyPrefix))
	store.Delete(types.TicketKey(
		index,
	))
}

// GetAllTicket returns all ticket
func (k Keeper) GetAllTicket(ctx context.Context) (list []types.Ticket) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TicketKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Ticket
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Update ticket validity
func (k Keeper) UpdateTicketValidity(ctx sdk.Context, ticketId string, valid bool) error {
	ticket, found := k.GetTicket(ctx, ticketId)
	if !found {
		return errors.Wrapf(types.ErrTicketNotFound, "ticket %s not found", ticketId)
	}

	ticket.Valid = valid
	k.SetTicket(ctx, ticket)
	return nil
}
