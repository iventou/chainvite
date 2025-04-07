package device_test

import (
	"testing"

	keepertest "github.com/iventou/chainvite/testutil/keeper"
	"github.com/iventou/chainvite/testutil/nullify"
	device "github.com/iventou/chainvite/x/device/module"
	"github.com/iventou/chainvite/x/device/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DeviceList: []types.Device{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DeviceKeeper(t)
	device.InitGenesis(ctx, k, genesisState)
	got := device.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DeviceList, got.DeviceList)
	// this line is used by starport scaffolding # genesis/test/assert
}
