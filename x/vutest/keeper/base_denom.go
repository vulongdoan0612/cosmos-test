package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"vutest/x/vutest/types"
)

// SetBaseDenom set baseDenom in the store
func (k Keeper) SetBaseDenom(ctx sdk.Context, baseDenom types.BaseDenom) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseDenomKey))
	b := k.cdc.MustMarshal(&baseDenom)
	store.Set([]byte{0}, b)
}

// GetBaseDenom returns baseDenom
func (k Keeper) GetBaseDenom(ctx sdk.Context) (val types.BaseDenom, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseDenomKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBaseDenom removes baseDenom from the store
func (k Keeper) RemoveBaseDenom(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BaseDenomKey))
	store.Delete([]byte{0})
}
