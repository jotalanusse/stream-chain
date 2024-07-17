package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Module name and store keys
const (
	// ModuleName defines the module name
	ModuleName = "lending"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName
)

// State
const (
	//used to store and access lending account data
	LendingAccountStoreKeyPrefix = "LendingAccount/"
)

// returns the store key for a specific lending account
func GetLendingAccountStoreKey(bech32AccAddress string) []byte {
	accAddress, err := sdk.AccAddressFromBech32(bech32AccAddress)
	if err != nil {
		panic(err) //TODO: handle error gracefully
	}
	return []byte(LendingAccountStoreKeyPrefix + accAddress.String())
}
