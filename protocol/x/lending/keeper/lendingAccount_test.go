package keeper_test

import (
	"reflect"
	"testing"

	sdkmath "cosmossdk.io/math"

	keepertest "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createPool(ctx sdk.Context, lendingKeeper *keeper.Keeper, assetDenom string) {
	// Create pool parameters
	params := types.PoolParams{
		AssetDenom: assetDenom,
		InterestRateModel: &types.InterestRateModel{
			BaseRate:        0.02,
			Multiplier:      0.1,
			JumpMultiplier:  0.5,
			TargetThreshold: 0.8,
		},
	}

	// Create the pool
	_, err := lendingKeeper.CreatePool(ctx, assetDenom, params)
	if err != nil {
		panic(err) // Handle the error as appropriate for your tests
	}
}

func TestCheckLendingAccountExists(t *testing.T) {
	// Setup keeper and context here
	ctx, lendingKeeper, _, _, _, _ := keepertest.LendingKeepers(t, true)

	// Create a test account
	bech32Addr := generateBech32Address()
	account := types.LendingAccount{
		Address:          bech32Addr,
		Nonce:            0,
		AccountPositions: []*types.AccountPosition{},
	}

	// Ensure the account doesn't exist
	nonExistentAddress, err := lendingKeeper.DoesLendingAccountExist(ctx, bech32Addr)
	require.NoError(t, err)
	assert.False(t, nonExistentAddress, "account should not exist")

	// Set the lending account
	lendingKeeper.SetLendingAccount(ctx, account)

	// Ensure the account does exist
	exists, err := lendingKeeper.DoesLendingAccountExist(ctx, bech32Addr)
	require.NoError(t, err)
	assert.True(t, exists, "account should exist")
}

func TestCreateLendingAccount(t *testing.T) {
	// Setup keeper and context here
	ctx, lendingKeeper, _, _, _, _ := keepertest.LendingKeepers(t, true)

	// Create a new lending account
	bech32Addr := generateBech32Address()
	account, err := lendingKeeper.CreateLendingAccount(ctx, bech32Addr)
	require.NoError(t, err, "failed to create lending account")
	assert.Equal(t, bech32Addr, account.Address, "account address should match")

	// Try to create the same account again
	account, err = lendingKeeper.CreateLendingAccount(ctx, bech32Addr)
	require.Error(t, err, "should fail, account already exists")
	assert.Equal(t, "account already exists", err.Error(), "error message should match")
}

func TestCreateLendingPosition(t *testing.T) {
	// Setup keeper and context here
	ctx, lendingKeeper, _, _, _, _ := keepertest.LendingKeepers(t, true)

	// Create a new lending account
	bech32Addr := generateBech32Address()
	account, err := lendingKeeper.CreateLendingAccount(ctx, bech32Addr)
	require.NoError(t, err, "failed to create lending account")

	// Create the pool for ETH
	createPool(ctx, lendingKeeper, "ETH")

	// Open a lending position
	amount := sdk.NewCoin("ETH", sdkmath.NewInt(100))
	updatedAccount, err := lendingKeeper.CreateLendingPosition(ctx, account.Address, amount)
	require.NoError(t, err, "failed to open lending position")

	// Check if the lending position was added
	retrievedAccount, exists := lendingKeeper.GetLendingAccount(ctx, updatedAccount.Address)
	require.True(t, exists, "account should exist")
	assert.Len(t, retrievedAccount.AccountPositions, 1, "there should be one account position")
	assert.Equal(t, []*sdk.Coin{&amount}, retrievedAccount.AccountPositions[0].LendingAssets, "collateral amounts should match")
	assert.Equal(t, []*sdk.Coin{&amount}, retrievedAccount.AccountPositions[0].Balance, "balance should match") // Updated to match slice structure

	// Ensure all account instances are equivalent
	t.Logf("Created Account: %+v\n", account)
	t.Logf("Retrieved Account: %+v\n", retrievedAccount)

	assert.True(t, reflect.DeepEqual(updatedAccount, retrievedAccount), "created and retrieved accounts should have the same contents")
}

func TestAddMultipleAssetsToLendingPosition(t *testing.T) {
	// Setup keeper and context here
	ctx, lendingKeeper, _, _, _, _ := keepertest.LendingKeepers(t, true)

	// Create a new lending account
	bech32Addr := generateBech32Address()
	account, err := lendingKeeper.CreateLendingAccount(ctx, bech32Addr)
	require.NoError(t, err, "failed to create lending account")

	// Create the pools for BTC, ETH, and SOL
	createPool(ctx, lendingKeeper, "BTC")
	createPool(ctx, lendingKeeper, "ETH")
	createPool(ctx, lendingKeeper, "SOL")

	// Open a 1 BTC lending position
	btcAmount := sdk.NewCoin("BTC", sdkmath.NewInt(1))
	_, err = lendingKeeper.CreateLendingPosition(ctx, account.Address, btcAmount)
	require.NoError(t, err, "failed to open BTC lending position")

	// Open an 10 ETH lending position
	ethAmount := sdk.NewCoin("ETH", sdkmath.NewInt(10))
	_, err = lendingKeeper.CreateLendingPosition(ctx, account.Address, ethAmount)
	require.NoError(t, err, "failed to open ETH lending position")

	// Open a 100 SOL lending position
	solAmount := sdk.NewCoin("SOL", sdkmath.NewInt(100))
	_, err = lendingKeeper.CreateLendingPosition(ctx, account.Address, solAmount)
	require.NoError(t, err, "failed to open SOL lending position")

	// Retrieve the updated account to verify positions
	updatedAccount, exists := lendingKeeper.GetLendingAccount(ctx, account.Address)
	require.True(t, exists, "account should exist")
	assert.Len(t, updatedAccount.AccountPositions, 3, "there should be three account positions")

	// Verify each position
	foundBTC, foundETH, foundSOL := false, false, false
	for _, position := range updatedAccount.AccountPositions {
		for _, collateral := range position.LendingAssets {
			switch collateral.Denom {
			case "BTC":
				assert.Equal(t, btcAmount, *collateral, "BTC position should match")
				foundBTC = true
			case "ETH":
				assert.Equal(t, ethAmount, *collateral, "ETH position should match")
				foundETH = true
			case "SOL":
				assert.Equal(t, solAmount, *collateral, "SOL position should match")
				foundSOL = true
			}
		}
	}

	// Ensure all positions are found
	assert.True(t, foundBTC, "BTC position not found")
	assert.True(t, foundETH, "ETH position not found")
	assert.True(t, foundSOL, "SOL position not found")
}
