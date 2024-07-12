package types

import (
	"github.com/cosmos/cosmos-sdk/types"
)

// Account defines a blockchain account structure
type Account struct {
	Address  types.AccAddress // Unique address of the account
	Balance  types.Coins      // Balance of various coins
	Nonce    uint64           // Used to ensure transaction order
	Metadata map[string]any   // Additional account information
}
