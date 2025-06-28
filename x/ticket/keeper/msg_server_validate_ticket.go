package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/iventou/chainvite/x/ticket/types"
)

func (k msgServer) ValidateTicket(goCtx context.Context, msg *types.MsgValidateTicket) (*types.MsgValidateTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the ticket exists
	ticket, found := k.GetTicket(ctx, msg.TicketId)
	if !found {
		return nil, errors.Wrapf(types.ErrTicketNotFound, "ticket %s not found", msg.TicketId)
	}

	// Check if the ticket is already used
	// if !ticket.Valid {
	// 	return nil, errors.Wrapf(types.ErrInvalidTicketValidity, "ticket %s is not valid", msg.TicketId)
	// }

	// Verify that the message creator is the ticket owner
	if msg.Creator != ticket.Owner {
		return nil, errors.Wrapf(types.ErrUnauthorizedAccess, "only the ticket owner can validate this ticket")
	}

	// Verify that the device exists and is active
	// if !k.deviceKeeper.IsDeviceActive(ctx, msg.DeviceId) {
	// 	return nil, errors.Wrapf(types.ErrUnauthorizedAccess, "device %s is not active", msg.DeviceId)
	// }

	// Verify that the device is authorized for this event
	// if !k.deviceKeeper.IsDeviceAuthorizedForEvent(ctx, msg.DeviceId, ticket.EventId) {
	// 	return nil, errors.Wrapf(types.ErrUnauthorizedAccess, "device %s is not authorized for event %s", msg.DeviceId, ticket.EventId)
	// }

	// Update the ticket status to "used" (not valid anymore)
	ticket.Valid = false
	k.SetTicket(ctx, ticket)

	payingReserveAddr, err := k.Keeper.addressCodec.StringToBytes("cosmos1dhdtjvlr672zverhsxhyztev3jr96kxt4vdz3m")
	if err != nil {
		return nil, errors.Wrap(err, "invalid address")
	}
	payingCoinDenom := "pedro"
	spendableCoins := k.Keeper.bankKeeper.SpendableCoins(ctx, payingReserveAddr)
	reserveCoin := sdk.NewCoin(payingCoinDenom, spendableCoins.AmountOf(payingCoinDenom))
	destionationAddr, err := k.Keeper.addressCodec.StringToBytes("cosmos10xw7klsm7xfht6guxfken9f9rpsqz5hf95n0wn")
	if err != nil {
		return nil, errors.Wrap(err, "invalid address")
	}
	if err := k.Keeper.bankKeeper.SendCoins(
		ctx,
		payingReserveAddr,
		destionationAddr,
		sdk.NewCoins(reserveCoin),
	); err != nil {
		return nil, errors.Wrapf(types.ErrFailedTicketValidationTransfer, "Failed to transfer coins on ticket validation %s", err.Error())
	}

	println("Send")

	// Emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeValidateTicket,
			sdk.NewAttribute("ticket_id", msg.TicketId),
			sdk.NewAttribute("device_id", msg.DeviceId),
			sdk.NewAttribute("validated", "true"),
		),
	})

	return &types.MsgValidateTicketResponse{}, nil
}
