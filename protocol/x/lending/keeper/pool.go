package keeper

import (
	sdkmath "cosmossdk.io/math"
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

// Checks if the pool has enough liquidity to fulfill a borrow request.
func (k Keeper) HasPoolLiquidityForBorrow(ctx sdk.Context, assetDenom string, amount sdk.Coin) (bool, error) {
	pool, found := k.GetPool(ctx, assetDenom)
	if !found {
		return false, errors.New("pool not found")
	}

	// Check if the pool has enough deposits to fulfill the borrow request
	availableLiquidity := pool.TotalDeposits.Sub(*pool.TotalBorrows)
	if availableLiquidity.IsLT(amount) {
		return false, errors.New("insufficient liquidity in the pool")
	}

	return true, nil
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

	// Retrieve the most recent updated pool and call updateRatesForPool
	updatedPool, found := k.GetPool(ctx, assetDenom)
	if !found {
		panic("pool not found after setting")
	}
	k.UpdateRatesForPool(ctx, &updatedPool)
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

	// Retrieve the most recent updated pool and call updateRatesForPool
	updatedPool, found := k.GetPool(ctx, assetDenom)
	if !found {
		panic("pool not found after setting")
	}
	k.UpdateRatesForPool(ctx, &updatedPool)
}

func (k Keeper) UpdateRatesForPool(ctx sdk.Context, pool *types.Pool) {
	// utilization rate = total quantity of asset borrowed / total quantity of asset deposited
	utilizationRate := sdkmath.LegacyNewDecFromInt(pool.TotalBorrows.Amount).Quo(sdkmath.LegacyNewDecFromInt(pool.TotalDeposits.Amount))

	// Fetch interest rate model parameters from the pool
	//Ensure values are calculated everytime with the correct precision
	//Suppose BaseRate is 0.02 (2%):
	//int64(pool.Params.InterestRateModel.BaseRate * 100) converts 0.02 to 2.
	//sdkmath.LegacyNewDecWithPrec(2, 2) converts 2 to 0.02 with a precision of 2 decimal places.
	baseRate := sdkmath.LegacyNewDecWithPrec(int64(pool.Params.InterestRateModel.BaseRate*100), 2)
	multiplier := sdkmath.LegacyNewDecWithPrec(int64(pool.Params.InterestRateModel.Multiplier*100), 2)
	jumpMultiplier := sdkmath.LegacyNewDecWithPrec(int64(pool.Params.InterestRateModel.JumpMultiplier*100), 2)
	targetThreshold := sdkmath.LegacyNewDecWithPrec(int64(pool.Params.InterestRateModel.TargetThreshold*100), 2)

	// Calculate the lending rate based on utilization
	var lendingRate sdkmath.LegacyDec
	if utilizationRate.LTE(targetThreshold) {
		lendingRate = baseRate.Add(utilizationRate.Mul(multiplier))
	} else {
		excessUtilization := utilizationRate.Sub(targetThreshold)
		lendingRate = baseRate.Add(targetThreshold.Mul(multiplier)).Add(excessUtilization.Mul(jumpMultiplier))
	}

	// Update the pool's current lending rate
	pool.CurLendingRate = lendingRate.MustFloat64()

	// Calculate the borrow rate, assuming a spread of 0.02 (2%)
	//TODO: Implement the actual logic to calculate the borrow rate
	spread := sdkmath.LegacyNewDecWithPrec(2, 2)
	pool.CurBorrowRate = lendingRate.Add(spread).MustFloat64()

	// Save the updated pool
	k.SetPool(ctx, *pool)
}
