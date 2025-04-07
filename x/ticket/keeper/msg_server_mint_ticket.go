package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/iventou/chainvite/x/ticket/types"
)

func (k msgServer) MintTicket(goCtx context.Context, msg *types.MsgMintTicket) (*types.MsgMintTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Generate a unique ID for the ticket using a hash based on ticket information
	ticketDataForHash := fmt.Sprintf("%s-%s-%s-%d", msg.EventId, msg.TicketType, msg.Owner, time.Now().UnixNano())
	hash := sha256.Sum256([]byte(ticketDataForHash))
	ticketId := hex.EncodeToString(hash[:])

	// Create the ticket with initial status as valid
	ticket := types.Ticket{
		Creator:    msg.Creator,
		Index:      ticketId, // Use ticketId as index
		TicketId:   ticketId,
		EventId:    msg.EventId,
		TicketType: msg.TicketType,
		Valid:      true, // Set to valid (instead of 'status: active')
		Owner:      msg.Owner,
		Metadata:   msg.Metadata,
	}

	// Store the ticket in the KVStore
	k.SetTicket(ctx, ticket)

	// Emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMintTicket,
			sdk.NewAttribute("ticket_id", ticketId),
			sdk.NewAttribute("event_id", msg.EventId),
			sdk.NewAttribute("owner", msg.Owner),
			sdk.NewAttribute("ticket_type", msg.TicketType),
		),
	})

	// Need to see what fields are available in the response
	return &types.MsgMintTicketResponse{}, nil
}

