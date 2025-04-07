package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// IsDeviceActive checks if a device is in "active" status
func (k Keeper) IsDeviceActive(ctx sdk.Context, deviceId string) bool {
	device, found := k.GetDevice(ctx, deviceId)
	return found && device.Status == "active"
}

// IsDeviceAuthorizedForEvent checks if a device is authorized for a specific event
func (k Keeper) IsDeviceAuthorizedForEvent(ctx sdk.Context, deviceId string, eventId string) bool {
	device, found := k.GetDevice(ctx, deviceId)
	return found && device.EventId == eventId && device.Status == "active"
}
