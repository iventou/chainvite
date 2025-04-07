package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDevice{}

func NewMsgCreateDevice(
	creator string,
	index string,
	deviceId string,
	manufacturer string,
	location string,
	eventId string,
	status string,

) *MsgCreateDevice {
	return &MsgCreateDevice{
		Creator:      creator,
		Index:        index,
		DeviceId:     deviceId,
		Manufacturer: manufacturer,
		Location:     location,
		EventId:      eventId,
		Status:       status,
	}
}

func (msg *MsgCreateDevice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDevice{}

func NewMsgUpdateDevice(
	creator string,
	index string,
	deviceId string,
	manufacturer string,
	location string,
	eventId string,
	status string,

) *MsgUpdateDevice {
	return &MsgUpdateDevice{
		Creator:      creator,
		Index:        index,
		DeviceId:     deviceId,
		Manufacturer: manufacturer,
		Location:     location,
		EventId:      eventId,
		Status:       status,
	}
}

func (msg *MsgUpdateDevice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteDevice{}

func NewMsgDeleteDevice(
	creator string,
	index string,

) *MsgDeleteDevice {
	return &MsgDeleteDevice{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteDevice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
