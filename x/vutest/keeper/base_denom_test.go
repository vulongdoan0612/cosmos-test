package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "vutest/testutil/keeper"
	"vutest/testutil/nullify"
	"vutest/x/vutest/keeper"
	"vutest/x/vutest/types"
)

func createTestBaseDenom(keeper *keeper.Keeper, ctx sdk.Context) types.BaseDenom {
	item := types.BaseDenom{}
	keeper.SetBaseDenom(ctx, item)
	return item
}

func TestBaseDenomGet(t *testing.T) {
	keeper, ctx := keepertest.VutestKeeper(t)
	item := createTestBaseDenom(keeper, ctx)
	rst, found := keeper.GetBaseDenom(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestBaseDenomRemove(t *testing.T) {
	keeper, ctx := keepertest.VutestKeeper(t)
	createTestBaseDenom(keeper, ctx)
	keeper.RemoveBaseDenom(ctx)
	_, found := keeper.GetBaseDenom(ctx)
	require.False(t, found)
}
