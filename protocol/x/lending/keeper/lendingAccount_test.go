package keeper_test

import (
	"testing"

	keepertest "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetAndGetLendingAccount(t *testing.T) {
	// Setup your keeper and context here
	ctx, lendingKeeper, _, _, _, _ := keepertest.LendingKeepers(t, true)

	// Create a test account
	address := "streamTestAccount1"
	account := types.LendingAccount{
		Address: address,
		Balance: []*sdk.Coin{},
		Nonce:   0,
	}

	// Set the lending account
	lendingKeeper.SetLendingAccount(ctx, account)

	// Retrieve the lending account
	retrievedAccount, exists := lendingKeeper.GetLendingAccount(ctx, address)
	require.True(t, exists, "account should exist")
	assert.Equal(t, account, retrievedAccount, "retrieved account should match the set account")
}

func TestGetNonExistentLendingAccount(t *testing.T) {
	// Setup your keeper and context here
	ctx, lendingKeeper, _, _, _, _ := keepertest.LendingKeepers(t, true)

	// Attempt to retrieve a non-existent account
	address := "nonExistentAddress"
	retrievedAccount, exists := lendingKeeper.GetLendingAccount(ctx, address)
	require.False(t, exists, "account should not exist")
	assert.Equal(t, types.LendingAccount{}, retrievedAccount, "retrieved account should be empty")
}
