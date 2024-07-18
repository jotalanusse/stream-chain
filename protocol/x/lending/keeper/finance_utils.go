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
	pool.CurrentLendingRate = lendingRate.MustFloat64()

	// Calculate the borrow rate, assuming a spread of 0.02 (2%)
	//TODO: Implement the actual logic to calculate the borrow rate
	spread := sdkmath.LegacyNewDecWithPrec(2, 2)
	pool.CurrentBorrowRate = lendingRate.Add(spread).MustFloat64()

	// Save the updated pool
	k.SetPool(ctx, *pool)
}

// CalculateLendingRate calculates the interest rate for a given principal over time.
// Parameters:
// - principal: The principal amount.
// - rate: The annual interest rate.
// - time: The time period in years.
func CalculateLendingRateForPosition(principal float64, rate float64, time float64) (interest float64, err error) {
	// Check if the inputs are valid
	if principal < 0 || rate < 0 || time < 0 {
		return 0, errors.New("invalid input values, all inputs should be non-negative")
	}

	// Calculate the interest using the simple interest formula: interest = principal * rate * time
	interest = principal * rate * time
	return interest, nil
}

// CalculateBorrowRate calculates the borrow rate for a given principal, considering the collateral.
// Parameters:
// - principal: The principal amount.
// - rate: The annual interest rate.
// - time: The time period in years.
// - collateral: The collateral amount.
func CalculateBorrowRate(principal float64, rate float64, time float64, collateral float64) (borrowRate float64, err error) {
	// Consider the collateral in the calculation the amount of interest accured from the loan so far
	// Use an appropriate formula to calculate the borrow rate
	// Return the calculated borrow rate and nil if successful, or 0 and error if not
	return 0, nil
}
