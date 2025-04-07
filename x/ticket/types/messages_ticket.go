package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateTicket{}

func NewMsgCreateTicket(
	creator string,
	index string,
	ticketId string,
	eventId string,
	ticketType string,
	valid bool,
	owner string,
	metadata string,

) *MsgCreateTicket {
	return &MsgCreateTicket{
		Creator:    creator,
		Index:      index,
		TicketId:   ticketId,
		EventId:    eventId,
		TicketType: ticketType,
		Valid:      valid,
		Owner:      owner,
		Metadata:   metadata,
	}
}

func (msg *MsgCreateTicket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTicket{}

func NewMsgUpdateTicket(
	creator string,
	index string,
	ticketId string,
	eventId string,
	ticketType string,
	valid bool,
	owner string,
	metadata string,

) *MsgUpdateTicket {
	return &MsgUpdateTicket{
		Creator:    creator,
		Index:      index,
		TicketId:   ticketId,
		EventId:    eventId,
		TicketType: ticketType,
		Valid:      valid,
		Owner:      owner,
		Metadata:   metadata,
	}
}

func (msg *MsgUpdateTicket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTicket{}

func NewMsgDeleteTicket(
	creator string,
	index string,

) *MsgDeleteTicket {
	return &MsgDeleteTicket{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteTicket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
