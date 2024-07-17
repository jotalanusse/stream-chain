package keeper_test

import (
	"reflect"
	"testing"

	sdkmath "cosmossdk.io/math"

	keepertest "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckLendingAccountExists(t *testing.T) {
	// Setup keeper and context here
	ctx, lendingKeeper, _, _, _, _ := keepertest.LendingKeepers(t, true)

	// Create a test account
	bech32Addr := generateBech32Address()
	account := types.LendingAccount{
		Address:            bech32Addr,
		Nonce:              0,
		LendingPositions:   []*sdk.Coin{},
		BorrowingPositions: []*types.Loan{},
	}

	//ensure the account doesn't exist
	nonExistentAddress, err := lendingKeeper.DoesLendingAccountExist(ctx, bech32Addr)
	require.NoError(t, err)
	assert.False(t, nonExistentAddress, "account should not exist")

	// Set the lending account
	lendingKeeper.SetLendingAccount(ctx, account)

	// ensure the account does exist
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

func TestOpenLendingPosition(t *testing.T) {
	// Setup keeper and context here
	ctx, lendingKeeper, _, _, _, _ := keepertest.LendingKeepers(t, true)

	// Create a new lending account
	bech32Addr := generateBech32Address()
	account, err := lendingKeeper.CreateLendingAccount(ctx, bech32Addr)
	require.NoError(t, err, "failed to create lending account")

	// Open a lending position
	amount := sdk.NewCoin("ETH", sdkmath.NewInt(100))
	updatedAccount, err := lendingKeeper.OpenLendingPosition(ctx, account.Address, amount)
	require.NoError(t, err, "failed to open lending position")

	// Check if the lending position was added
	retrievedAccount, exists := lendingKeeper.GetLendingAccount(ctx, updatedAccount.Address)
	require.True(t, exists, "account should exist")
	assert.Len(t, retrievedAccount.LendingPositions, 1, "there should be one lending position")
	assert.Equal(t, amount, *retrievedAccount.LendingPositions[0], "lending position should match")

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

	// Open a 1 BTC lending position
	btcAmount := sdk.NewCoin("BTC", sdkmath.NewInt(1))
	_, err = lendingKeeper.OpenLendingPosition(ctx, account.Address, btcAmount)
	require.NoError(t, err, "failed to open BTC lending position")

	// Open an 10 ETH lending position
	ethAmount := sdk.NewCoin("ETH", sdkmath.NewInt(10))
	_, err = lendingKeeper.OpenLendingPosition(ctx, account.Address, ethAmount)
	require.NoError(t, err, "failed to open ETH lending position")

	// Open a 100 SOL lending position
	solAmount := sdk.NewCoin("SOL", sdkmath.NewInt(100))
	_, err = lendingKeeper.OpenLendingPosition(ctx, account.Address, solAmount)
	require.NoError(t, err, "failed to open SOL lending position")

	// Retrieve the updated account to verify positions
	updatedAccount, exists := lendingKeeper.GetLendingAccount(ctx, account.Address)
	require.True(t, exists, "account should exist")
	assert.Len(t, updatedAccount.LendingPositions, 3, "there should be three lending positions")

	// Verify each position
	foundBTC, foundETH, foundSOL := false, false, false
	for _, position := range updatedAccount.LendingPositions {
		switch position.Denom {
		case "BTC":
			assert.Equal(t, btcAmount, *position, "BTC position should match")
			foundBTC = true
		case "ETH":
			assert.Equal(t, ethAmount, *position, "ETH position should match")
			foundETH = true
		case "SOL":
			assert.Equal(t, solAmount, *position, "SOL position should match")
			foundSOL = true
		}
	}

	// Ensure all positions are found
	assert.True(t, foundBTC, "BTC position not found")
	assert.True(t, foundETH, "ETH position not found")
	assert.True(t, foundSOL, "SOL position not found")
}
