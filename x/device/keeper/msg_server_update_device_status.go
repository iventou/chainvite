package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/iventou/chainvite/x/device/types"
)

func (k msgServer) UpdateDeviceStatus(goCtx context.Context, msg *types.MsgUpdateDeviceStatus) (*types.MsgUpdateDeviceStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateDeviceStatusResponse{}, nil
}
