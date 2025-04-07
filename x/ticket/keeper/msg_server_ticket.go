package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/iventou/chainvite/x/ticket/types"
)

func (k msgServer) CreateTicket(goCtx context.Context, msg *types.MsgCreateTicket) (*types.MsgCreateTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetTicket(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var ticket = types.Ticket{
		Creator:    msg.Creator,
		Index:      msg.Index,
		TicketId:   msg.TicketId,
		EventId:    msg.EventId,
		TicketType: msg.TicketType,
		Valid:      msg.Valid,
		Owner:      msg.Owner,
		Metadata:   msg.Metadata,
	}

	k.SetTicket(
		ctx,
		ticket,
	)
	return &types.MsgCreateTicketResponse{}, nil
}

func (k msgServer) UpdateTicket(goCtx context.Context, msg *types.MsgUpdateTicket) (*types.MsgUpdateTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTicket(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var ticket = types.Ticket{
		Creator:    msg.Creator,
		Index:      msg.Index,
		TicketId:   msg.TicketId,
		EventId:    msg.EventId,
		TicketType: msg.TicketType,
		Valid:      msg.Valid,
		Owner:      msg.Owner,
		Metadata:   msg.Metadata,
	}

	k.SetTicket(ctx, ticket)

	return &types.MsgUpdateTicketResponse{}, nil
}

func (k msgServer) DeleteTicket(goCtx context.Context, msg *types.MsgDeleteTicket) (*types.MsgDeleteTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTicket(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTicket(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteTicketResponse{}, nil
}
