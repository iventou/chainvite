package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransferTicket{}

func NewMsgTransferTicket(creator string, ticketId string, newOwner string) *MsgTransferTicket {
	return &MsgTransferTicket{
		Creator:  creator,
		TicketId: ticketId,
		NewOwner: newOwner,
	}
}

func (msg *MsgTransferTicket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
