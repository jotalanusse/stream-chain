package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// validates if the user has sufficent funds to open a lending position
// TODO: expand out to also validate for borrow positions
func (k Keeper) HasSufficientFunds(ctx sdk.Context, bech32AccAddr string, amount sdk.Coin) (bool, error) {
	// Fetch the user's wallet balance
	walletBalance, err := k.GetWalletBalance(ctx, bech32AccAddr, amount.Denom)
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

// fetches the balance of a specific asset in the user's wallet.
// TODO: Implement the actual logic to fetch the wallet balance
func (k Keeper) GetWalletBalance(ctx sdk.Context, bech32AccAddr string, denom string) (sdk.Coin, error) {
	// Implementation for fetching the wallet balance for a specific asset
	// This might involve querying the bank module or another component managing balances
	// Return the balance and nil if successful, or sdk.Coin{} and error if not
	return sdk.NewCoin("BTC", sdkmath.NewInt(0)), nil
}

// QueryOraclePrice queries the oracle for the current price of a given asset.
// Parameters:
// - assetID: The unique identifier of the asset.
// Returns:
// - price: The current price of the asset.
// - error: Error if the query fails.
func QueryOraclePrice(assetID string) (price float64, err error) {
	// Connect to the oracle service
	// Query the price using assetID
	// Return the price and nil if successful, or 0 and error if not
	return 0, nil
}

// CalculateInterestRate calculates the interest rate for a given principal over time.
// Parameters:
// - principal: The principal amount.
// - rate: The annual interest rate.
// - time: The time period in years.
// Returns:
// - interest: The calculated interest.
// - error: Error if the calculation fails.
func CalculateInterestRate(principal float64, rate float64, time float64) (interest float64, err error) {
	// Calculate the interest using the given formula (simple or compound)
	// Return the calculated interest and nil if successful, or 0 and error if not
	return 0, nil

}

// CalculateBorrowRate calculates the borrow rate for a given principal, considering the collateral.
// Parameters:
// - principal: The principal amount.
// - rate: The annual interest rate.
// - time: The time period in years.
// - collateral: The collateral amount.
// Returns:
// - borrowRate: The calculated borrow rate.
// - error: Error if the calculation fails.
func CalculateBorrowRate(principal float64, rate float64, time float64, collateral float64) (borrowRate float64, err error) {
	// Consider the collateral in the calculation
	// Use an appropriate formula to calculate the borrow rate
	// Return the calculated borrow rate and nil if successful, or 0 and error if not
	return 0, nil
}
