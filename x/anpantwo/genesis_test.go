package anpantwo_test

import (
	"testing"

	keepertest "github.com/anpan2val/anpan-chain/testutil/keeper"
	"github.com/anpan2val/anpan-chain/testutil/nullify"
	"github.com/anpan2val/anpan-chain/x/anpantwo"
	"github.com/anpan2val/anpan-chain/x/anpantwo/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TaskList: []types.Task{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TaskCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AnpantwoKeeper(t)
	anpantwo.InitGenesis(ctx, *k, genesisState)
	got := anpantwo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TaskList, got.TaskList)
	require.Equal(t, genesisState.TaskCount, got.TaskCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
