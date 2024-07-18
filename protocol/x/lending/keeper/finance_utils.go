package keeper

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

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

// Helper function to get the net asset price in USD
func (k Keeper) GetNetAssetPriceUSD(ctx sdk.Context, assets []*sdk.Coin) (float64, error) {
	totalPriceInUSD := 0.0

	for _, asset := range assets {
		assetPrice, err := QueryOraclePrice(asset.Denom)
		if err != nil {
			return 0.0, err
		}
		totalPriceInUSD += float64(asset.Amount.Int64()) * assetPrice
	}

	return totalPriceInUSD, nil
}

// QueryOraclePrice queries the oracle for the current price of a given asset
// TODO: Implement the actual logic to query the oracle
func QueryOraclePrice(assetID string) (price float64, err error) {
	// Connect to the oracle service/use the prices module to query the price
	// Return the price and nil if successful, or 0 and error if not
	return 0, nil
}

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
