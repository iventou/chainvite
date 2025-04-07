package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DeviceKeyPrefix is the prefix to retrieve all Device
	DeviceKeyPrefix = "Device/value/"
)

// DeviceKey returns the store key to retrieve a Device from the index fields
func DeviceKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
