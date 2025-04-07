package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/iventou/chainvite/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgMintTicket_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgMintTicket
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgMintTicket{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgMintTicket{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
