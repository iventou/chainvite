package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/iventou/chainvite/x/device/types"
)

func (k msgServer) RegisterDevice(goCtx context.Context, msg *types.MsgRegisterDevice) (*types.MsgRegisterDeviceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterDeviceResponse{}, nil
}
