package anpanchain_test

import (
	"testing"

	keepertest "github.com/anpan2val/anpan-chain/testutil/keeper"
	"github.com/anpan2val/anpan-chain/testutil/nullify"
	"github.com/anpan2val/anpan-chain/x/anpanchain"
	"github.com/anpan2val/anpan-chain/x/anpanchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PeopleList: []types.People{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PeopleCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AnpanchainKeeper(t)
	anpanchain.InitGenesis(ctx, *k, genesisState)
	got := anpanchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PeopleList, got.PeopleList)
	require.Equal(t, genesisState.PeopleCount, got.PeopleCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
