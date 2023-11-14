package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"vutest/x/vutest/types"
)

func (k Keeper) BaseDenom(goCtx context.Context, req *types.QueryGetBaseDenomRequest) (*types.QueryGetBaseDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetBaseDenom(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBaseDenomResponse{BaseDenom: val}, nil
}
