package types

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgDepositLiquidityIntoPool{}

// MsgDepositLiquidityIntoPool constructs a `MsgDepositLiquidityIntoPool` from an
// a liquidity provider, a token denom, and an amount.
func NewMsgDepositLiquidityIntoPool(
	liquidityProvider string,
	tokenDenom string,
	amount string,
) *MsgDepositLiquidityIntoPool {
	return &MsgDepositLiquidityIntoPool{
		LiquidityProvider: liquidityProvider,
		TokenDenom:        tokenDenom,
		Amount:            amount,
	}
}

// ValidateBasic runs validation on the fields of a MsgDepositLiquidityIntoPool.
func (msg *MsgDepositLiquidityIntoPool) ValidateBasic() error {
	// Validate account sender.
	_, err := sdk.AccAddressFromBech32(msg.LiquidityProvider)
	if err != nil {
		return ErrInvalidAccountAddress
	}

	if msg.TokenDenom == "" {
		return ErrInvalidTokenDenom
	}

	bigAmount, err := ConvertStringToBigInt(msg.Amount)
	if err != nil {
		return ErrInvalidDepositAmount
	}

	if bigAmount.Cmp(big.NewInt(0)) <= 0 {
		return ErrInvalidDepositAmount
	}
	return nil
}
