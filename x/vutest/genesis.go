package vutest

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"vutest/x/vutest/keeper"
	"vutest/x/vutest/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	var baseDenom = types.BaseDenom{
    	Denom:   sdk.Coin{Denom: "stake", Amount: sdk.NewInt(1)}, // Giả sử số lượng là 1
	}
	if genState.BaseDenom != nil {
		k.SetBaseDenom(ctx, baseDenom)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all baseDenom
	baseDenom, found := k.GetBaseDenom(ctx)
	if found {
		genesis.BaseDenom = &baseDenom
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
