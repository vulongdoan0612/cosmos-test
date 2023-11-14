package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "vutest/testutil/keeper"
	"vutest/x/vutest/keeper"
	"vutest/x/vutest/types"
)

func TestBaseDenomMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.VutestKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateBaseDenom{Creator: creator}
	_, err := srv.CreateBaseDenom(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetBaseDenom(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestBaseDenomMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateBaseDenom
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateBaseDenom{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateBaseDenom{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.VutestKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateBaseDenom{Creator: creator}
			_, err := srv.CreateBaseDenom(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateBaseDenom(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetBaseDenom(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestBaseDenomMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteBaseDenom
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteBaseDenom{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteBaseDenom{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.VutestKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateBaseDenom(wctx, &types.MsgCreateBaseDenom{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteBaseDenom(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetBaseDenom(ctx)
				require.False(t, found)
			}
		})
	}
}
