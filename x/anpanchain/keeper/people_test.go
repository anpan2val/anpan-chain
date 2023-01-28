package keeper_test

import (
	"testing"

	keepertest "github.com/anpan2val/anpan-chain/testutil/keeper"
	"github.com/anpan2val/anpan-chain/testutil/nullify"
	"github.com/anpan2val/anpan-chain/x/anpanchain/keeper"
	"github.com/anpan2val/anpan-chain/x/anpanchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNPeople(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.People {
	items := make([]types.People, n)
	for i := range items {
		items[i].Id = keeper.AppendPeople(ctx, items[i])
	}
	return items
}

func TestPeopleGet(t *testing.T) {
	keeper, ctx := keepertest.AnpanchainKeeper(t)
	items := createNPeople(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPeople(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPeopleRemove(t *testing.T) {
	keeper, ctx := keepertest.AnpanchainKeeper(t)
	items := createNPeople(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePeople(ctx, item.Id)
		_, found := keeper.GetPeople(ctx, item.Id)
		require.False(t, found)
	}
}

func TestPeopleGetAll(t *testing.T) {
	keeper, ctx := keepertest.AnpanchainKeeper(t)
	items := createNPeople(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPeople(ctx)),
	)
}

func TestPeopleCount(t *testing.T) {
	keeper, ctx := keepertest.AnpanchainKeeper(t)
	items := createNPeople(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPeopleCount(ctx))
}
