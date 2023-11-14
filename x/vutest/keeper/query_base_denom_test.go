package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "vutest/testutil/keeper"
	"vutest/testutil/nullify"
	"vutest/x/vutest/types"
)

func TestBaseDenomQuery(t *testing.T) {
	keeper, ctx := keepertest.VutestKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestBaseDenom(keeper, ctx)
	tests := []struct {
		desc     string
		request  *types.QueryGetBaseDenomRequest
		response *types.QueryGetBaseDenomResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetBaseDenomRequest{},
			response: &types.QueryGetBaseDenomResponse{BaseDenom: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.BaseDenom(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
