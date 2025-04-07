package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/iventou/chainvite/testutil/keeper"
	"github.com/iventou/chainvite/x/device/keeper"
	"github.com/iventou/chainvite/x/device/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDeviceMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.DeviceKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateDevice{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateDevice(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetDevice(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestDeviceMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateDevice
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateDevice{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateDevice{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateDevice{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.DeviceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateDevice{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateDevice(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateDevice(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetDevice(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestDeviceMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteDevice
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteDevice{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteDevice{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteDevice{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.DeviceKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateDevice(ctx, &types.MsgCreateDevice{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteDevice(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetDevice(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
