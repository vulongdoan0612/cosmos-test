package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"vutest/x/vutest/types"
)

func (k msgServer) CreateBaseDenom(goCtx context.Context, msg *types.MsgCreateBaseDenom) (*types.MsgCreateBaseDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetBaseDenom(ctx)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var baseDenom = types.BaseDenom{
    	Denom:   sdk.Coin{Denom: "stake", Amount: sdk.NewInt(1)}, // Giả sử số lượng là 1
		Creator: msg.Creator,
	}

	k.SetBaseDenom(
		ctx,
		baseDenom,
	)
	return &types.MsgCreateBaseDenomResponse{}, nil
}

func (k msgServer) UpdateBaseDenom(goCtx context.Context, msg *types.MsgUpdateBaseDenom) (*types.MsgUpdateBaseDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
		valFound, isFound := k.GetBaseDenom(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}
																		
	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var baseDenom = types.BaseDenom{
    	Denom:   sdk.Coin{Denom: "stake", Amount: sdk.NewInt(1)}, // Giả sử số lượng là 1
		Creator: msg.Creator,
	}

	k.SetBaseDenom(ctx, baseDenom)

	return &types.MsgUpdateBaseDenomResponse{}, nil
}

func (k msgServer) DeleteBaseDenom(goCtx context.Context, msg *types.MsgDeleteBaseDenom) (*types.MsgDeleteBaseDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetBaseDenom(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveBaseDenom(ctx)

	return &types.MsgDeleteBaseDenomResponse{}, nil
}
