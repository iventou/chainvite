package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/iventou/chainvite/x/ticket/keeper"
	"github.com/iventou/chainvite/x/ticket/types"
)

func SimulateMsgTransferTicket(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTransferTicket{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the TransferTicket simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "TransferTicket simulation not implemented"), nil, nil
	}
}
