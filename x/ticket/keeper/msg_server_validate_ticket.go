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
	if !ticket.Valid {
		return nil, errors.Wrapf(types.ErrInvalidTicketValidity, "ticket %s is not valid", msg.TicketId)
	}

	// Verify that the message creator is the ticket owner
	if msg.Creator != ticket.Owner {
		return nil, errors.Wrapf(types.ErrUnauthorizedAccess, "only the ticket owner can validate this ticket")
	}

	// Update the ticket status to "used" (not valid anymore)
	ticket.Valid = false
	k.SetTicket(ctx, ticket)

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
