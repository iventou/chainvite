package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/iventou/chainvite/testutil/keeper"
	"github.com/iventou/chainvite/testutil/nullify"
	"github.com/iventou/chainvite/x/device/keeper"
	"github.com/iventou/chainvite/x/device/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDevice(keeper keeper.Keeper, ctx context.Context, n int) []types.Device {
	items := make([]types.Device, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetDevice(ctx, items[i])
	}
	return items
}

func TestDeviceGet(t *testing.T) {
	keeper, ctx := keepertest.DeviceKeeper(t)
	items := createNDevice(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDevice(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDeviceRemove(t *testing.T) {
	keeper, ctx := keepertest.DeviceKeeper(t)
	items := createNDevice(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDevice(ctx,
			item.Index,
		)
		_, found := keeper.GetDevice(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestDeviceGetAll(t *testing.T) {
	keeper, ctx := keepertest.DeviceKeeper(t)
	items := createNDevice(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDevice(ctx)),
	)
}
