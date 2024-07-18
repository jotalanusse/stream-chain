package keeper

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

/*
	Functions related to the IRM model and health calculations for lending positions.
	Possibly seperate out further the IRM model code from health code as the IRM code gets expanded more
	- UpdateAccruedInterestOnPosition: Used to update the accrued interest on all lending/borrowing positions of an account
	- GetAccountPositionsHealth: Returns the health of all positions of an account
	- CalculatePositionHealth: Calculated TWV / Liabilities to determine health of a position
	- validateInputParameters: Validates the input parameters for health calculations
	- calculateTotalWeightedValue: (balance of position * price in USD of position) * liquidation threshold)
	- calculateLiabilities: (borrowed amount in USD + interestAccrued + sum of initial collateral amount)
	- CalculateInterestAccruedForPosition: basic interest * rate * time in days to determine interest

*/

// gets the new interest gained for a user's collateral and borrowing over a specified period of time
// Call before doing any health accounting and liquidation checks to ensure up to date values
// TODO: Decide how to handle the time parameter from when the last time accrued interest was calculated for borrowing and lending positions
// TODO: Include this function calls at the start or end of a block to ensure the interest is up to date
func (k Keeper) UpdateAccruedInterestOnPosition(ctx sdk.Context, account types.LendingAccount, time float64) error {

	yearInDays := 365.25
	time = yearInDays

	for _, position := range account.AccountPositions {
		if !position.IsPureLending {
			// Update accrued borrowed amount with accrued interest
			accruedBorrowedInterest, err := k.CalculateInterestAccruedForPosition(ctx, *position.BorrowedAsset, false, time)
			if err != nil {
				return err
			}
			accruedBorrowedInterestDec := sdkmath.LegacyNewDecWithPrec(int64(accruedBorrowedInterest*100), 2)
			position.AccruedBorrowedAsset.Amount = position.AccruedBorrowedAsset.Amount.Add(sdkmath.NewIntFromBigInt(accruedBorrowedInterestDec.BigInt()))
		}

		// Update accured interest for all collateral amounts
		for _, collateral := range position.CollateralAssets {
			accruedLendingInterest, err := k.CalculateInterestAccruedForPosition(ctx, *collateral, true, time)
			if err != nil {
				return err
			}
			accruedLendingInterestDec := sdkmath.LegacyNewDecWithPrec(int64(accruedLendingInterest*100), 2)
			collateral.Amount = collateral.Amount.Add(sdkmath.NewIntFromBigInt(accruedLendingInterestDec.BigInt()))
		}
	}

	//update the balance of the position
	//TODO: implement the logic to update the balance of the position

	// Update the new account values to the kv-store
	k.SetLendingAccount(ctx, account)

	return nil
}

// get the health of all the positions of an account in USD doing: (totalWeightedValue / liabilities)
func (k Keeper) GetAccountPositionsHealth(ctx sdk.Context, bech32AccAddr string) (allPositionsHealth []types.PositionWithHealth, unhealthyPositions []types.PositionWithHealth, err error) {

	// ensure lendingAccount exists and get update principal + accrued interest on all lending & borrow positions
	account, found := k.GetLendingAccount(ctx, bech32AccAddr)
	if !found {
		return nil, nil, errors.New("account not found")
	}

	// Find and update the specific account position of only borrowing positions
	for _, position := range account.AccountPositions {
		if position.IsPureLending {
			continue
		}
		// Calculate the health for a single position
		//positionHealth, err := k.CalculatePositionHealth(ctx, *position.BorrowedAsset, position.Balance, position.CollateralAssets, yearInDays)
		positionHealth, err := k.CalculatePositionHealth(ctx, *position)
		if err != nil {
			return nil, nil, err
		}

		// Create a PositionWithHealth struct to store the position and its health
		positionWithHealth := types.PositionWithHealth{
			Position: *position,
			Health:   positionHealth,
		}

		// Store the health of the position
		allPositionsHealth = append(allPositionsHealth, positionWithHealth)

		// add to unhealthy positions if the health is less than one
		if positionWithHealth.Health.LT(sdkmath.LegacyNewDec(1)) {
			unhealthyPositions = append(unhealthyPositions, positionWithHealth)
		}
	}

	return allPositionsHealth, unhealthyPositions, nil
}

// (balance of position * price in USD of position) * liquidation threshold) / (borrowed amount in USD + interestAccrued + sum of initial collateral amount)
// func (k Keeper) CalculatePositionHealth(ctx sdk.Context, borrowedAsset sdk.Coin, balance []*sdk.Coin, collateralAssets []*sdk.Coin, time float64) (sdkmath.LegacyDec, error)
func (k Keeper) CalculatePositionHealth(ctx sdk.Context, position types.AccountPosition) (sdkmath.LegacyDec, error) {
	// Validate input parameters
	if err := validateInputParameters(*position.BorrowedAsset, position.Balance, position.CollateralAssets); err != nil {
		return sdkmath.LegacyDec{}, err
	}

	// Calculate the numerator
	totalWeightedValue, err := k.calculateTotalWeightedValue(ctx, position)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}

	// Calculate the denominator
	liabilities, err := k.calculateLiabilities(ctx, position)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}

	if liabilities.IsZero() {
		return sdkmath.LegacyDec{}, fmt.Errorf("denominator is zero, cannot divide by zero")
	}

	// Calculate the position health
	positionHealth := totalWeightedValue.Quo(liabilities)
	return positionHealth, nil
}

func validateInputParameters(borrowedAsset sdk.Coin, balance []*sdk.Coin, collateralAssets []*sdk.Coin) error {
	if borrowedAsset.Amount.IsNegative() || len(balance) == 0 || len(collateralAssets) == 0 {
		return fmt.Errorf("invalid input values, all inputs should be non-negative and not empty")
	}
	return nil
}

// numerator for calculating position health: (balance of position * price in USD of position) * liquidation threshold)
func (k Keeper) calculateTotalWeightedValue(ctx sdk.Context, position types.AccountPosition) (sdkmath.LegacyDec, error) {
	//(balance of position * price in USD of position)
	balanceValueInUSD, err := k.GetNetAssetPriceUSD(ctx, position.Balance)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}

	balanceValueDec := sdkmath.LegacyNewDecWithPrec(int64(balanceValueInUSD), 2)

	//get liquidation threshold of what is being borrowed
	borrowedAssetPool, found := k.GetPool(ctx, position.BorrowedAsset.Denom)
	if !found {
		return sdkmath.LegacyDec{}, fmt.Errorf("pool not found")
	}
	borrowedAssetPoolLiquidationThreshold, err := k.GetLiquidationThreshold(ctx, borrowedAssetPool.Params.AssetDenom)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}

	return balanceValueDec.Mul(borrowedAssetPoolLiquidationThreshold), nil
}

// denominator for calculating position health: (borrowed amount in USD + accuredBorrowedInterest + sum of collateral amount (already contains the interest))
func (k Keeper) calculateLiabilities(ctx sdk.Context, position types.AccountPosition) (sdkmath.LegacyDec, error) {
	borrowedValueInUSD, err := QueryOraclePrice(position.BorrowedAsset.Denom)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}
	borrowedValueInUSDDec := sdkmath.LegacyNewDecWithPrec(int64(borrowedValueInUSD*100), 2).Mul(sdkmath.LegacyNewDecFromInt(position.BorrowedAsset.Amount))
	accuredBorrowedInterest := sdkmath.LegacyNewDecFromInt(position.AccruedBorrowedAsset.Amount)

	collateralValueInUSD, err := k.GetNetAssetPriceUSD(ctx, position.CollateralAssets)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}
	collateralValueInUSDDec := sdkmath.LegacyNewDecWithPrec(int64(collateralValueInUSD), 2)

	// Sum up liabilities
	return borrowedValueInUSDDec.Add(accuredBorrowedInterest).Add(collateralValueInUSDDec), nil
}

// calculates the interest gained for a user's position over a specific amount of time, works for lending and borrowing amounts
// TODO: days represents number of days in the year for now, but should change to a different time scale
func (k Keeper) CalculateInterestAccruedForPosition(ctx sdk.Context, asset sdk.Coin, isLendingRate bool, days float64) (float64, error) {
	// Validate inputs
	if asset.Amount.IsNegative() || days < 0 {
		return 0, errors.New("invalid input values, all inputs should be non-negative")
	}

	// Convert asset amount to float64 for interest calculation
	principal := float64(asset.Amount.Int64())

	// Convert days to years for the annual rate calculation
	timeInYears := days / 365.25 // Using 365.25 accounts for leap years

	// Calculate the interest using the simple interest formula: interest = principal * rate * time
	pool, err := k.GetPool(ctx, asset.Denom)
	if !err {
		return 0, fmt.Errorf("pool not found")
	}

	lendingOrBorrowRate := pool.CurBorrowRate
	if isLendingRate {
		lendingOrBorrowRate = pool.CurLendingRate
	}

	// Calculate the interest based on the lending or borrowing rate
	interest := principal * lendingOrBorrowRate * timeInYears
	return interest, nil
}
