package types

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgWithdrawLiquidityFromPool{}

// MsgWithdrawLiquidityFromPool constructs a `MsgWithdrawLiquidityFromPool` from an
// a liquidity provider, a token denom, and an amount.
func NewMsgWithdrawLiquidityFromPool(
	liquidityProvider string,
	tokenDenom string,
	amount string,
) *MsgWithdrawLiquidityFromPool {
	return &MsgWithdrawLiquidityFromPool{
		LiquidityProvider: liquidityProvider,
		TokenDenom:        tokenDenom,
		Amount:            amount,
	}
}

// ValidateBasic runs validation on the fields of a MsgWithdrawLiquidityFromPool.
func (msg *MsgWithdrawLiquidityFromPool) ValidateBasic() error {
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
		return ErrInvalidWithdrawAmount
	}

	if bigAmount.Cmp(big.NewInt(0)) <= 0 {
		return ErrInvalidDepositAmount
	}
	return nil
}
