package keeper_test

import (
	"testing"

	keepertest "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// generateBech32Address generates a Bech32 address for testing
// TODO: Figure out how other account addresses are generated in dydx
// as creating new private key and public key for each test is not efficient

// Example of a Bech32 address: dydx1zaeeeycequcwxw5vuqnlz3nlax4zljpqsekv4j
// dydx1zaeeeycequcwxw5vuqnlz3nlax4zljpqsekv4j
func generateBech32Address() string {
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()
	addr := sdk.AccAddress(pubKey.Address())
	return addr.String()
}

func TestSetAndGetLendingAccount(t *testing.T) {
	// Setup keeper and context here
	ctx, lendingKeeper, _, _, _, _ := keepertest.LendingKeepers(t, true)

	// Create a test account
	bech32Addr := "dydx1zaeeeycequcwxw5vuqnlz3nlax4zljpqsekv4j"
	account := types.LendingAccount{
		Address:          bech32Addr,
		Nonce:            0,
		AccountPositions: []*types.AccountPosition{},
	}

	// Set the lending account
	lendingKeeper.SetLendingAccount(ctx, account)

	// Retrieve the lending account
	retrievedAccount, exists := lendingKeeper.GetLendingAccount(ctx, bech32Addr)
	require.True(t, exists, "account should exist")
	assert.Equal(t, account, retrievedAccount, "retrieved account should match the set account")
}

func TestGetNonExistentLendingAccount(t *testing.T) {
	// Setup your keeper and context here
	ctx, lendingKeeper, _, _, _, _ := keepertest.LendingKeepers(t, true)

	// Attempt to retrieve a non-existent account
	bech32Addr := "dydx1zaeeeycequcwxw5vuqnlz3nlax4zljpqsekv4j"
	retrievedAccount, exists := lendingKeeper.GetLendingAccount(ctx, bech32Addr)
	require.False(t, exists, "account should not exist")
	assert.Equal(t, types.LendingAccount{}, retrievedAccount, "retrieved account should be empty")
}
