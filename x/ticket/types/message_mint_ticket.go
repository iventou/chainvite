package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMintTicket{}

func NewMsgMintTicket(creator string, eventId string, ticketType string, owner string, metadata string) *MsgMintTicket {
	return &MsgMintTicket{
		Creator:    creator,
		EventId:    eventId,
		TicketType: ticketType,
		Owner:      owner,
		Metadata:   metadata,
	}
}

func (msg *MsgMintTicket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
