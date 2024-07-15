package keeper

import (
	//errorsmod "cosmossdk.io/errors"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
	File includes methods related to setting/getting lending accounts.
	File methods include:
	- GetLendingAccount
	- checkLendingAccountExists
	- SetLendingAccount
	- CreateLendingAccount
*/

// retrieves a lending account if it exists.
func (k Keeper) GetLendingAccount(ctx sdk.Context, address string) (types.LendingAccount, bool) {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(address)) {
		return types.LendingAccount{}, false
	}
	bz := store.Get([]byte(address))
	var account types.LendingAccount
	k.cdc.MustUnmarshal(bz, &account)

	// Ensure the Balance field is initialized as an empty slice if it's nil (go treats nil and empty slice differently)
	if account.Balance == nil {
		account.Balance = []*sdk.Coin{}
	}

	return account, true
}

// checks if a lending account exists for the given address.
func (k Keeper) checkLendingAccountExists(ctx sdk.Context, address string) (bool, error) {
	_, accountExists := k.GetLendingAccount(ctx, address)
	if accountExists {
		// Commented out the specific error type
		// return true, errorsmod.Wrapf(
		// 	types.ErrAccountAlreadyExists,
		// 	"account with address %v already exists",
		// 	address,
		// )
		println("account with address", address, "already exists")
		return true, nil
	}
	return false, nil
}

// marshals and stores the lending account in kv-store
func (k Keeper) SetLendingAccount(ctx sdk.Context, account types.LendingAccount) {
	store := ctx.KVStore(k.storeKey)
	accountKey := []byte(account.Address)

	bz := k.cdc.MustMarshal(&account)
	store.Set(accountKey, bz)
}

// checks for existence, validates inputs, creates, and stores a lending account.
func (k Keeper) CreateLendingAccount(ctx sdk.Context, address string) (types.LendingAccount, error) {
	// check if the lending account exists, if it does return error
	accountExists, err := k.checkLendingAccountExists(ctx, address)
	if accountExists || err != nil {
		// type specific error handling
		// return types.LendingAccount{}, err
		println("Error creating lending account for address", address)
		return types.LendingAccount{}, nil
	}

	// Create and store the new lending account
	account := types.LendingAccount{
		Address: address,
		Balance: []*sdk.Coin{},
		Nonce:   0,
	}

	k.SetLendingAccount(ctx, account)

	return account, nil
}

// Updates the values of an existing lending account.
func (k Keeper) UpdateLendingAccount(ctx sdk.Context, address string, newBalance []*sdk.Coin, newNonce uint64) error {
	// Check if the lending account exists
	account, exists := k.GetLendingAccount(ctx, address)
	if !exists {
		println("Lending account does not exist for address", address)
		return nil // Replace with a specific error when reintroducing error handling
	}

	// Update the account's balance and nonce
	account.Balance = newBalance
	account.Nonce = newNonce

	// Marshal and store the updated account
	k.SetLendingAccount(ctx, account)

	return nil
}

// func someFunctionName2() int {
// 	return 1
// }
