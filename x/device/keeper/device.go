package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/iventou/chainvite/x/device/types"
)

// SetDevice set a specific device in the store from its index
func (k Keeper) SetDevice(ctx context.Context, device types.Device) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DeviceKeyPrefix))
	b := k.cdc.MustMarshal(&device)
	store.Set(types.DeviceKey(
		device.Index,
	), b)
}

// GetDevice returns a device from its index
func (k Keeper) GetDevice(
	ctx context.Context,
	index string,

) (val types.Device, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DeviceKeyPrefix))

	b := store.Get(types.DeviceKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDevice removes a device from the store
func (k Keeper) RemoveDevice(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DeviceKeyPrefix))
	store.Delete(types.DeviceKey(
		index,
	))
}

// GetAllDevice returns all device
func (k Keeper) GetAllDevice(ctx context.Context) (list []types.Device) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DeviceKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Device
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
