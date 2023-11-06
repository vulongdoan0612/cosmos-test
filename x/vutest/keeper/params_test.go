package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "vutest/testutil/keeper"
	"vutest/x/vutest/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VutestKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
