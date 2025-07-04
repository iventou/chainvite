package types

const (
	// ModuleName defines the module name
	ModuleName = "device"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_device"
)

var (
	ParamsKey = []byte("p_device")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
