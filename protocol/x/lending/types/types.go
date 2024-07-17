package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewPool creates a new Pool instance.
func NewPool(assetDenom string, params PoolParams) Pool {
	zeroInt := sdkmath.NewInt(1) // Assuming sdk.NewInt exists and can be used to create a zero value.
	return Pool{
		TotalDeposits: &sdk.Coin{Denom: assetDenom, Amount: zeroInt},
		TotalBorrows:  &sdk.Coin{Denom: assetDenom, Amount: zeroInt},
		Params:        &params,
	}
}

// NewAccountPosition creates a new AccountPosition instance.
func NewAccountPosition(amount sdk.Coin) *AccountPosition {
	return &AccountPosition{
		CollateralAmounts: []*sdk.Coin{&amount},
		BorrowedAmounts:   nil, // Set to nil for pure lending
		Balance:           &amount,
		IsPureLending:     true,
	}
}
