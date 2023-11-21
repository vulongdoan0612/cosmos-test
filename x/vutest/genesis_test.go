package vutest_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "vutest/testutil/keeper"
	"vutest/testutil/nullify"
	"vutest/x/vutest"
	"vutest/x/vutest/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		BaseDenom: &types.BaseDenom{
			Denom:   sdk.Coin{Denom: "stake"}, // Giả sử số lượng là 1
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VutestKeeper(t)
	vutest.InitGenesis(ctx, *k, genesisState)
	got := vutest.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.BaseDenom, got.BaseDenom)
	// this line is used by starport scaffolding # genesis/test/assert
}
