package vutest_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "vutest/testutil/keeper"
	"vutest/testutil/nullify"
	"vutest/x/vutest"
	"vutest/x/vutest/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VutestKeeper(t)
	vutest.InitGenesis(ctx, *k, genesisState)
	got := vutest.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
