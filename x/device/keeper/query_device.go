package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/iventou/chainvite/x/device/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DeviceAll(ctx context.Context, req *types.QueryAllDeviceRequest) (*types.QueryAllDeviceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var devices []types.Device

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	deviceStore := prefix.NewStore(store, types.KeyPrefix(types.DeviceKeyPrefix))

	pageRes, err := query.Paginate(deviceStore, req.Pagination, func(key []byte, value []byte) error {
		var device types.Device
		if err := k.cdc.Unmarshal(value, &device); err != nil {
			return err
		}

		devices = append(devices, device)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDeviceResponse{Device: devices, Pagination: pageRes}, nil
}

func (k Keeper) Device(ctx context.Context, req *types.QueryGetDeviceRequest) (*types.QueryGetDeviceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetDevice(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDeviceResponse{Device: val}, nil
}
