package keeper

import (
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

// Checks if a pool exists for the given asset denomination.
func (k Keeper) DoesPoolExist(ctx sdk.Context, assetDenom string) (bool, error) {
	_, poolExists := k.GetPool(ctx, assetDenom)
	if poolExists {
		println("pool with asset denomination", assetDenom, "already exists")
		return true, nil
	}
	return false, nil
}

// Checks for existence, validates inputs, creates, and stores a new pool.
func (k Keeper) CreatePool(ctx sdk.Context, assetDenom string, params types.PoolParams) (types.Pool, error) {
	// Check if the pool exists, if it does return error
	poolExists, err := k.DoesPoolExist(ctx, assetDenom)
	if err != nil || poolExists {
		return types.Pool{}, errors.New("pool already exists")
	}

	// Create and store the new pool
	pool := types.NewPool(assetDenom, params)

	k.SetPool(ctx, pool)

	return pool, nil
}

func (k Keeper) UpdatePoolDeposits(ctx sdk.Context, assetDenom string, amount sdk.Coin) {
	pool, found := k.GetPool(ctx, assetDenom)
	if !found {
		panic("pool not found")
	}

	// Create a new Coin that is the sum of the existing TotalDeposits and the amount to add.
	// Then take the address of this new Coin to match the expected *Coin type.
	sum := pool.TotalDeposits.Add(amount)
	pool.TotalDeposits = &sum
	k.SetPool(ctx, pool)
}

func (k Keeper) UpdatePoolBorrows(ctx sdk.Context, assetDenom string, amount sdk.Coin) {
	pool, found := k.GetPool(ctx, assetDenom)
	if !found {
		panic("pool not found")
	}
	// Create a new Coin that is the sum of the existing TotalBorrows and the amount to add.
	// Then take the address of this new Coin to match the expected *Coin type.
	sum := pool.TotalBorrows.Add(amount)
	pool.TotalBorrows = &sum
	k.SetPool(ctx, pool)
}
