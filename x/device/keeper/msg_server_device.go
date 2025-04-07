package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/iventou/chainvite/x/device/types"
)

func (k msgServer) CreateDevice(goCtx context.Context, msg *types.MsgCreateDevice) (*types.MsgCreateDeviceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDevice(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var device = types.Device{
		Creator:      msg.Creator,
		Index:        msg.Index,
		DeviceId:     msg.DeviceId,
		Manufacturer: msg.Manufacturer,
		Location:     msg.Location,
		EventId:      msg.EventId,
		Status:       msg.Status,
	}

	k.SetDevice(
		ctx,
		device,
	)
	return &types.MsgCreateDeviceResponse{}, nil
}

func (k msgServer) UpdateDevice(goCtx context.Context, msg *types.MsgUpdateDevice) (*types.MsgUpdateDeviceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDevice(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var device = types.Device{
		Creator:      msg.Creator,
		Index:        msg.Index,
		DeviceId:     msg.DeviceId,
		Manufacturer: msg.Manufacturer,
		Location:     msg.Location,
		EventId:      msg.EventId,
		Status:       msg.Status,
	}

	k.SetDevice(ctx, device)

	return &types.MsgUpdateDeviceResponse{}, nil
}

func (k msgServer) DeleteDevice(goCtx context.Context, msg *types.MsgDeleteDevice) (*types.MsgDeleteDeviceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDevice(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDevice(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteDeviceResponse{}, nil
}
