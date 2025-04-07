package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/iventou/chainvite/x/ticket/types"
)

func (k msgServer) TransferTicket(goCtx context.Context, msg *types.MsgTransferTicket) (*types.MsgTransferTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the ticket exists
	ticket, found := k.GetTicket(ctx, msg.TicketId)
	if !found {
		return nil, errors.Wrapf(types.ErrTicketNotFound, "ticket %s not found", msg.TicketId)
	}

	// Check if the ticket is valid
	if !ticket.Valid {
		return nil, errors.Wrapf(types.ErrInvalidTicketValidity, "ticket %s is not valid", msg.TicketId)
	}

	// Check if the message creator is the current owner of the ticket
	if ticket.Owner != msg.Creator {
		return nil, errors.Wrapf(types.ErrUnauthorizedAccess, "only the current owner can transfer the ticket")
	}

	// Update the ticket owner
	ticket.Owner = msg.NewOwner
	k.SetTicket(ctx, ticket)

	// Emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeTransferTicket,
			sdk.NewAttribute("ticket_id", msg.TicketId),
			sdk.NewAttribute("previous_owner", msg.Creator),
			sdk.NewAttribute("new_owner", msg.NewOwner),
		),
	})

	return &types.MsgTransferTicketResponse{}, nil
}

