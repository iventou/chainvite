package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/iventou/chainvite/testutil/keeper"
	"github.com/iventou/chainvite/x/ticket/keeper"
	"github.com/iventou/chainvite/x/ticket/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestTicketMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.TicketKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateTicket{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateTicket(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetTicket(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestTicketMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateTicket
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateTicket{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateTicket{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateTicket{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.TicketKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateTicket{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateTicket(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateTicket(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetTicket(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestTicketMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteTicket
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteTicket{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteTicket{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteTicket{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.TicketKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateTicket(ctx, &types.MsgCreateTicket{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteTicket(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetTicket(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
