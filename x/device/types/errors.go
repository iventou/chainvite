package types

import (
	"cosmossdk.io/errors"
)

// Device module error codes
const (
	ErrorCodeDeviceNotFound      = uint32(1)
	ErrorCodeInvalidDeviceStatus = uint32(2)
	ErrorCodeUnauthorizedAccess  = uint32(3)
)

// Device module errors
var (
	ErrDeviceNotFound      = errors.Register(ModuleName, ErrorCodeDeviceNotFound, "device not found")
	ErrInvalidDeviceStatus = errors.Register(ModuleName, ErrorCodeInvalidDeviceStatus, "invalid device status")
	ErrUnauthorizedAccess  = errors.Register(ModuleName, ErrorCodeUnauthorizedAccess, "unauthorized access")
)

