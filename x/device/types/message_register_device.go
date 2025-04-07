package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegisterDevice{}

func NewMsgRegisterDevice(creator string, manufacturer string, location string, eventId string) *MsgRegisterDevice {
	return &MsgRegisterDevice{
		Creator:      creator,
		Manufacturer: manufacturer,
		Location:     location,
		EventId:      eventId,
	}
}

func (msg *MsgRegisterDevice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
