package keeper

import (
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

/*
	File includes methods related to setting/getting lending accounts.
	- Main methods:
		- CreateLendingAccount
			- Accounts can hold multiple positions on the same asset
		- CreateLendingPosition
			- Creates a new position for a given asset with a uuid
		- UpdateLendingPosition
			- Based on the asset and uuid updates a pre-existing position in the account

	- Helper methods:
		- findAndUpdatePosition
			- Finds and updates a specific position in the account using the given asset and uuid
		- updatePositionBalance
			- Updates the balance of a specific asset in the account after the position has been updated
		- DoesLendingAccountExist
			- Checks if a lending account exists for the given address
		- getPositionsForAsset
			- Returns a list of positions for a given asset within an account
*/

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
func (k Keeper) CreateLendingPosition(ctx sdk.Context, bech32AccAddr string, asset sdk.Coin) (types.LendingAccount, error) {
	// Initialize an empty LendingAccount struct to return in case of errors
	var emptyAccount types.LendingAccount

	// Validate pool existence and sufficient funds
	valid, err := k.TransactionValidation(ctx, bech32AccAddr, asset.Denom, asset)
	if err != nil || !valid {
		return emptyAccount, err // Error from TransactionValidation already formatted
	}

	// Ensure lending account exists
	account, found := k.GetLendingAccount(ctx, bech32AccAddr)
	if !found {
		return emptyAccount, errors.New("account not found")
	}

	// Create a new position if it does not exist
	newPosition := types.NewAccountPosition(asset)

	// Add the new position to account positions
	account.AccountPositions = append(account.AccountPositions, newPosition)

	// Store the updated account
	k.SetLendingAccount(ctx, account)

	// Update the pool deposits
	k.UpdatePoolDeposits(ctx, asset)

	return account, nil
}

// UpdateLendingPosition updates an existing lending position for a given asset and quantity, and returns the updated lending account.
func (k Keeper) UpdateLendingPosition(ctx sdk.Context, bech32AccAddr string, asset sdk.Coin, positionID string) (types.LendingAccount, error) {
	// Initialize an empty LendingAccount struct to return in case of errors
	var emptyAccount types.LendingAccount

	// Validate pool existence and sufficient funds
	valid, err := k.TransactionValidation(ctx, bech32AccAddr, asset.Denom, asset)
	if err != nil || !valid {
		return emptyAccount, err // Error from TransactionValidation already formatted
	}

	// Check if the lending account exists
	account, found := k.GetLendingAccount(ctx, bech32AccAddr)
	if !found {
		return emptyAccount, errors.New("account not found")
	}

	// Find and update the specific account position using the UUID and asset
	err = findAndUpdatePosition(&account, asset, positionID)
	if err != nil {
		return emptyAccount, err
	}

	// Store the updated account
	k.SetLendingAccount(ctx, account)

	// Update the pool deposits
	k.UpdatePoolDeposits(ctx, asset)

	return account, nil
}

// finds and update the position + balance based on the uuid and asset
func findAndUpdatePosition(account *types.LendingAccount, asset sdk.Coin, positionID string) error {
	for _, position := range account.AccountPositions {
		if position.ID == positionID {
			for _, collateral := range position.CollateralAssets {
				if collateral.Denom == asset.Denom {
					collateral.Amount = collateral.Amount.Add(asset.Amount)
					// Account position has been updated; update the balance for that asset as well
					err := updatePositionBalance(position, asset)
					if err != nil {
						return err
					}
					return nil
				}
			}
		}
	}
	return errors.New("position not found")
}

// updates the balance for the given asset in the specified account position.
func updatePositionBalance(position *types.AccountPosition, asset sdk.Coin) error {
	for _, balance := range position.Balance {
		if balance.Denom == asset.Denom {
			balance.Amount = balance.Amount.Add(asset.Amount)
			return nil
		}
	}
	// denomination of that asset within the balance should exist
	return errors.New("asset not found in position balance")
}

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

// returns a list of positions for a given asset within an account
func getPositionsForAsset(account *types.LendingAccount, assetDenom string) ([]*types.AccountPosition, error) {
	var positions []*types.AccountPosition
	for _, position := range account.AccountPositions {
		for _, collateral := range position.CollateralAssets {
			if collateral.Denom == assetDenom {
				positions = append(positions, position)
				break
			}
		}
	}
	if len(positions) == 0 {
		return nil, errors.New("no positions found for the given asset")
	}
	return positions, nil
}
