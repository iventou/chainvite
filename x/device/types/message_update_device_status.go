package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateDeviceStatus{}

func NewMsgUpdateDeviceStatus(creator string, deviceId string, status string) *MsgUpdateDeviceStatus {
	return &MsgUpdateDeviceStatus{
		Creator:  creator,
		DeviceId: deviceId,
		Status:   status,
	}
}

func (msg *MsgUpdateDeviceStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
