package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/iventou/chainvite/testutil/keeper"
	"github.com/iventou/chainvite/testutil/nullify"
	"github.com/iventou/chainvite/x/device/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDeviceQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DeviceKeeper(t)
	msgs := createNDevice(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetDeviceRequest
		response *types.QueryGetDeviceResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetDeviceRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetDeviceResponse{Device: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetDeviceRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetDeviceResponse{Device: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetDeviceRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Device(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestDeviceQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DeviceKeeper(t)
	msgs := createNDevice(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllDeviceRequest {
		return &types.QueryAllDeviceRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DeviceAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Device), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Device),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DeviceAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Device), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Device),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.DeviceAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Device),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.DeviceAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
