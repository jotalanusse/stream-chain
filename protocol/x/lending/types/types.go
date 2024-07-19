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
		CollateralAssets: []*sdk.Coin{&amount},
		BorrowedAsset:    nil, // Set to nil for pure lending
		Balance:          []*sdk.Coin{&amount},
		IsPureLending:    true,
	}
}

// Define a new struct to hold the AccountPosition and its health value
type PositionWithHealth struct {
	Position AccountPosition
	Health   sdkmath.LegacyDec
}

//Define the lending keeper's interface once implementation is more fleshed
//out with how we want the various function to be defined
//example of interface implementation:
// /home/ec2-user/stream-chain/protocol/x/perpetuals/types/types.go
