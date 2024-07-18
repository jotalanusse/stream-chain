package keeper

import (
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

/*
	File includes methods related to setting/getting lending accounts.
	File methods include:
	- DoesLendingAccountExist
	- CreateLendingAccount
	- OpenLendingPosition
	- CreateLendingAccount
*/

// checks if a lending account exists for the given address.
func (k Keeper) DoesLendingAccountExist(ctx sdk.Context, bech32AccAddr string) (bool, error) {
	_, accountExists := k.GetLendingAccount(ctx, bech32AccAddr)
	if accountExists {
		// Commented out the specific error type
		// return true, errorsmod.Wrapf(
		// 	types.ErrAccountAlreadyExists,
		// 	"account with address %v already exists",
		// 	address,
		// )
		println("account with address", bech32AccAddr, "already exists")
		return true, nil
	}
	return false, nil
}

// Checks for existence, validates inputs, creates, and stores a lending account.
func (k Keeper) CreateLendingAccount(ctx sdk.Context, bech32AccAddr string) (types.LendingAccount, error) {
	// Check if the lending account exists, if it does return error
	accountExists, err := k.DoesLendingAccountExist(ctx, bech32AccAddr)
	if err != nil || accountExists {
		return types.LendingAccount{}, errors.New("account already exists")
	}

	// Create and store the new lending account
	account := types.LendingAccount{
		Address:          bech32AccAddr,
		Nonce:            0,
		AccountPositions: []*types.AccountPosition{},
	}

	k.SetLendingAccount(ctx, account)

	return account, nil
}

// Opens a new lending position for a given asset and quantity and returns the updated lending account.
func (k Keeper) OpenLendingPosition(ctx sdk.Context, bech32AccAddr string, amount sdk.Coin) (types.LendingAccount, error) {
	// Initialize an empty LendingAccount struct to return in case of errors
	var emptyAccount types.LendingAccount

	// validate pool existence and sufficient funds
	valid, err := k.TransactionValidation(ctx, bech32AccAddr, amount.Denom, amount)
	if err != nil || !valid {
		return emptyAccount, err // Error from TransactionValidation already formatted
	}

	// Check if the lending account exists
	account, found := k.GetLendingAccount(ctx, bech32AccAddr)
	if !found {
		return emptyAccount, errors.New("account not found")
	}

	newPosition := types.NewAccountPosition(amount)

	// Add the new position to account positions
	account.AccountPositions = append(account.AccountPositions, newPosition)

	// Store the updated account
	k.SetLendingAccount(ctx, account)

	// Update the pool deposits
	k.UpdatePoolDeposits(ctx, amount.Denom, amount)

	return account, nil
}

// updates an existing lending position for a given asset and quantity, and returns the updated lending account.
func (k Keeper) UpdateLendingPosition(ctx sdk.Context, bech32AccAddr string, amount sdk.Coin) (types.LendingAccount, error) {
	// Initialize an empty LendingAccount struct to return in case of errors
	var emptyAccount types.LendingAccount

	// validate pool existence and sufficient funds
	valid, err := k.TransactionValidation(ctx, bech32AccAddr, amount.Denom, amount)
	if err != nil || !valid {
		return emptyAccount, err // Error from TransactionValidation already formatted
	}

	// Check if the lending account exists
	account, found := k.GetLendingAccount(ctx, bech32AccAddr)
	if !found {
		return emptyAccount, errors.New("account not found")
	}

	// Find and update the specific account position for the given asset
	updated := false
	for _, position := range account.AccountPositions {
		for _, collateral := range position.CollateralAssets {
			if collateral.Denom == amount.Denom {
				collateral.Amount = collateral.Amount.Add(amount.Amount)
				//account position has been update the balance for that asset as well
				for _, balance := range position.Balance {
					if balance.Denom == amount.Denom {
						balance.Amount = balance.Amount.Add(amount.Amount)
						break
					}
				}
				updated = true
				break
			}
		}
		if updated {
			break
		}
	}

	// If the position was not found, return an error
	if !updated {
		return emptyAccount, errors.New("position not found")
	}

	// Store the updated account
	k.SetLendingAccount(ctx, account)

	// Update the pool deposits
	k.UpdatePoolDeposits(ctx, amount.Denom, amount)

	return account, nil
}
