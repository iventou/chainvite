package ticket_test

import (
	"testing"

	keepertest "github.com/iventou/chainvite/testutil/keeper"
	"github.com/iventou/chainvite/testutil/nullify"
	ticket "github.com/iventou/chainvite/x/ticket/module"
	"github.com/iventou/chainvite/x/ticket/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TicketList: []types.Ticket{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TicketKeeper(t)
	ticket.InitGenesis(ctx, k, genesisState)
	got := ticket.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TicketList, got.TicketList)
	// this line is used by starport scaffolding # genesis/test/assert
}
