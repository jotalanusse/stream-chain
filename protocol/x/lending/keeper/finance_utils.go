package keeper

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

// fetches the balance of a specific asset from the user's account in the bank module.
// TODO: Implement the actual logic to fetch the bank balance
func (k Keeper) GetBankBalance(ctx sdk.Context, bech32AccAddr string, denom string) (sdk.Coin, error) {
	// Implementation for fetching the account balance for a specific asset
	// This might involve querying the bank module or another component managing balances
	// Return the balance and nil if successful, or sdk.Coin{} and error if not
	return sdk.NewCoin("BTC", sdkmath.NewInt(0)), nil
}

// validates if the user has sufficent funds to open a lending position
// TODO: expand out to also validate for borrow positions
func (k Keeper) HasSufficientFunds(ctx sdk.Context, bech32AccAddr string, amount sdk.Coin) (bool, error) {
	// Fetch the user's wallet balance
	walletBalance, err := k.GetBankBalance(ctx, bech32AccAddr, amount.Denom)
	if err != nil || !walletBalance.Amount.IsZero() {
		return false, err
	}

	// Check if the wallet balance is greater than or equal to the amount
	//Comment out for now since getWalletBalance isn't implemented
	// if walletBalance.IsLT(amount) {
	// 	return false, nil // User does not have sufficient funds
	// }

	return true, nil // User has sufficient funds
}

// transactionValidation validates that the user has sufficient funds and the pool exists.
func (k Keeper) TransactionValidation(ctx sdk.Context, bech32AccAddr string, assetDenom string, amount sdk.Coin) (bool, error) {
	// validate that the pool exists
	poolExists, err := k.DoesPoolExist(ctx, assetDenom)
	if err != nil {
		return false, err
	}
	if !poolExists {
		return false, errors.New(fmt.Sprintf("pool with asset denomination %s does not exist", assetDenom))
	}

	// validate that the account has sufficient funds
	hasFunds, err := k.HasSufficientFunds(ctx, bech32AccAddr, amount)
	if err != nil {
		return false, err
	}
	if !hasFunds {
		return false, errors.New("account does not have sufficient funds")
	}

	return true, nil
}

// QueryOraclePrice queries the oracle for the current price of a given asset
// TODO: Implement the actual logic to query the oracle
func QueryOraclePrice(assetID string) (price float64, err error) {
	// Connect to the oracle service/use the prices module to query the price
	// Return the price and nil if successful, or 0 and error if not
	return 0, nil
}

// calculates the interest gained for a user's position over a specific amount of time, works for lending and borrowing amounts
// TODO: days represents number of days in the year for now, but should change to a different time scale
func (k Keeper) CalculateInterestAccruedForPosition(ctx sdk.Context, asset sdk.Coin, rate float64, days float64) (float64, error) {
	// Validate inputs
	if asset.Amount.IsNegative() || rate < 0 || days < 0 {
		return 0, errors.New("invalid input values, all inputs should be non-negative")
	}

	// Convert asset amount to float64 for interest calculation
	principal := float64(asset.Amount.Int64())

	// Convert days to years for the annual rate calculation
	timeInYears := days / 365.25 // Using 365.25 accounts for leap years

	// Calculate the interest using the simple interest formula: interest = principal * rate * time
	interest := principal * rate * timeInYears
	return interest, nil
}

// get the total price of all the borrwed/collateral assets in the account in USD
func getNetAssetPriceUSD(ctx sdk.Context, assets []sdk.Coin) (float64, error) {
	totalPriceInUSD := 0.0

	for _, asset := range assets {
		assetPriceInUSD, err := QueryOraclePrice(asset.Denom)
		if err != nil {
			return 0.0, err
		}
		totalPriceInUSD += float64(asset.Amount.Int64()) * assetPriceInUSD
	}

	return totalPriceInUSD, nil
}

// CalculatePositionHealth calculates the health of a position
// totalWeightedValue / liabilities
// (balance of position * price in USD of position) * liquidation threshold) / (borrowed amount in USD + interestAccrued + sum of initial collateral amount)
func (k Keeper) CalculatePositionHealth(ctx sdk.Context, borrowedAmount sdk.Coin, balance []*sdk.Coin, collateralAssets []*sdk.Coin, time float64) (sdkmath.LegacyDec, error) {
	// Validate input parameters
	if borrowedAmount.Amount.IsNegative() || len(balance) == 0 || len(collateralAssets) == 0 {
		return sdkmath.LegacyDec{}, fmt.Errorf("invalid input values, all inputs should be non-negative and not empty")
	}

	// Calculate the numerator: (balance of position * price in USD of position) * liquidation threshold
	balanceValueInUSD, err := getNetAssetPriceUSD(ctx, balance)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}

	borrowedAssetPool, found := k.GetPool(ctx, borrowedAmount.Denom)
	if !found {
		return sdkmath.LegacyDec{}, fmt.Errorf("pool not found")
	}
	borrowedAssetPoolLiquidationThreshold := sdkmath.LegacyNewDecWithPrec(int64(borrowedAssetPool.Params.LiquidationThreshold*100), 2)

	totalWeightedValue := balanceValueInUSD.Mul(borrowedAssetPoolLiquidationThreshold)

	// Calculate the denominator: borrowed amount in USD + interestAccrued + sum of initial collateral amount
	borrowedValueInUSD, err := QueryOraclePrice(borrowedAmount.Denom)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}

	borrowedValueInUSDDec := sdkmath.LegacyNewDecWithPrec(int64(borrowedValueInUSD*100), 2).Mul(sdkmath.LegacyNewDecFromInt(borrowedAmount.Amount))

	// Calculate interest accrued
	interestAccrued, err := k.CalculateInterestAccruedForPosition(ctx, borrowedAmount, borrowedAssetPool.CurLendingRate, time)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}

	interestAccruedDec := sdkmath.LegacyNewDecWithPrec(int64(interestAccrued*100), 2)

	collateralValueInUSD, err := getNetAssetPriceUSD(ctx, collateralAssets)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}

	// Sum up liabilities
	liabilities := borrowedValueInUSDDec.Add(interestAccruedDec).Add(collateralValueInUSD)

	if liabilities.IsZero() {
		return sdkmath.LegacyDec{}, fmt.Errorf("denominator is zero, cannot divide by zero")
	}

	// Calculate the position health
	positionHealth := totalWeightedValue.Quo(liabilities)
	return positionHealth, nil
}

// get the health of all the positions of an account (calculations done in USDC to standardize numbers across assets)
func (k Keeper) GetAccountHealth(ctx sdk.Context, bech32AccAddr string) (allPositionsHealth []sdkmath.LegacyDec, unhealthyPositions []types.AccountPosition, err error) {

	// Check if the lending account exists
	account, found := k.GetLendingAccount(ctx, bech32AccAddr)
	if !found {
		return nil, nil, errors.New("account not found")
	}

	// Find and update the specific account position of only borrowing positions
	for _, position := range account.AccountPositions {
		if position.IsPureLending {
			continue
		}

		yearInDays := 365.25

		// Calculate the health for a single position
		positionHealth, err := k.CalculatePositionHealth(ctx, *position.BorrowedAsset, position.Balance, position.CollateralAssets, yearInDays)
		if err != nil {
			return nil, nil, err
		}

		// Store the health of the position
		allPositionsHealth = append(allPositionsHealth, positionHealth)

		// Check if the position health is less than 1 and add to unhealthyPositions if so
		if positionHealth.LT(sdkmath.LegacyNewDec(1)) {
			unhealthyPositions = append(unhealthyPositions, *position)
		}
	}

	return allPositionsHealth, unhealthyPositions, nil
}
