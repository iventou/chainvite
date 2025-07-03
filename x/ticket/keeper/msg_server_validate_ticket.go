package keeper

import (
	"context"
	gerrors "errors"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/iventou/chainvite/x/ticket/types"
)

func (k msgServer) splitPayment(ctx sdk.Context, fromAddress sdk.AccAddress, addresses [][]byte, amount math.Int) error {
	payingCoinDenom := "pedro"
	spendableCoins := k.Keeper.bankKeeper.SpendableCoins(ctx, fromAddress)
	fromAddressTotalAmount := sdk.NewCoin(payingCoinDenom, spendableCoins.AmountOf(payingCoinDenom))

	reserveCoin := sdk.NewCoin(payingCoinDenom, math.NewInt(10))
	totalAmount := int64(len(addresses) * 10)
	if fromAddressTotalAmount.Amount.LT(reserveCoin.Amount.Mul(math.NewInt(totalAmount))) {
		return errors.Wrap(types.ErrFailedTicketValidationTransfer, "insuficient funds on payment account")
	}

	var mError error

	for _, addr := range addresses {
		if err := k.Keeper.bankKeeper.SendCoins(
			ctx,
			fromAddress,
			addr,
			sdk.NewCoins(reserveCoin),
		); err != nil {
			gerrors.Join(err, mError)
			return errors.Wrapf(types.ErrFailedTicketValidationTransfer, "Failed to transfer coins on ticket validation %s", err.Error())
		}
	}

	if gerrors.Unwrap(mError) != nil {
		return errors.Wrapf(mError, "Failed transfers")
	}

	return nil
}

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
	//

	// Update the ticket status to "used" (not valid anymore)
	ticket.Valid = false
	k.SetTicket(ctx, ticket)

	payingReserveAddr, err := k.Keeper.addressCodec.StringToBytes("cosmos1k9vyqkrc6z40gt3fcfplxgxtjcj4l5k355yv6d")
	if err != nil {
		return nil, errors.Wrap(err, "invalid address for payingReserveAddr")
	}
	destionationAddr01, err := k.Keeper.addressCodec.StringToBytes("cosmos1sr3cky55tj7kx4k8ch9vpares4wpkqvtmqps9n")
	if err != nil {
		return nil, errors.Wrap(err, "invalid address for destionationAddr01")
	}

	destionationAddr02, err := k.Keeper.addressCodec.StringToBytes("cosmos18rtlpdvmjv7erkvlxfcxqm86xcr9fchwcwjfk9")
	if err != nil {
		return nil, errors.Wrap(err, "invalid address for destionationAddr02")
	}

	k.splitPayment(ctx, payingReserveAddr, [][]byte{
		destionationAddr01,
		destionationAddr02,
	}, math.NewInt(10))

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
