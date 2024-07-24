package keeper

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CalcBorrowRate returns the borrow rate calculated based on expectedLiquidity and availableLiquidity
// @param expectedLiquidity Expected liquidity in the pool
// @param availableLiquidity Available liquidity in the pool
// @notice In RAY format
func (k Keeper) CalcBorrowRate(ctx sdk.Context, tokenDenom string, expectedLiquidity, availableLiquidity *big.Int) (*big.Int, error) {
	poolParams, found := k.GetPoolParams(ctx, tokenDenom)
	if !found {
		return nil, types.ErrInvalidTokenDenom
	}

	if expectedLiquidity.Cmp(big.NewInt(0)) == 0 || expectedLiquidity.Cmp(availableLiquidity) < 0 {
		return poolParams.BaseRate, nil
	}

	//      expectedLiquidity - availableLiquidity
	// U = -------------------------------------
	//             expectedLiquidity

	utilisation_eighteen := new(big.Int).Mul(big.NewInt(types.EIGHTEEN_DECIMALS), new(big.Int).Sub(expectedLiquidity, availableLiquidity))
	utilisation_eighteen = utilisation_eighteen.Div(utilisation_eighteen, expectedLiquidity)

	U_Optimal_inverted_eighteen := new(big.Int).Sub(big.NewInt(types.EIGHTEEN_DECIMALS), poolParams.OptimalUtilizationRatio)

	// if U < Uoptimal:
	//
	//                                    U
	// borrowRate = Rbase + Rslope1 * ----------
	//                                 Uoptimal
	//

	if utilisation_eighteen.Cmp(poolParams.OptimalUtilizationRatio) < 0 {
		// borrowRate = Rbase + Rslope1 * (U / Uoptimal)
		borrowRate := new(big.Int).Mul(poolParams.SlopeOneRate, utilisation_eighteen)
		borrowRate = borrowRate.Div(borrowRate, poolParams.OptimalUtilizationRatio)
		borrowRate = borrowRate.Add(borrowRate, poolParams.BaseRate)
		return borrowRate, nil
	}

	// if U >= Uoptimal:
	//
	//                                           U - Uoptimal
	// borrowRate = Rbase + Rslope1 + Rslope2 * --------------
	//
	borrowRate := new(big.Int).Sub(utilisation_eighteen, poolParams.OptimalUtilizationRatio)
	borrowRate = borrowRate.Mul(borrowRate, poolParams.SlopeTwoRate)
	borrowRate = borrowRate.Div(borrowRate, U_Optimal_inverted_eighteen)
	borrowRate = borrowRate.Add(borrowRate, poolParams.SlopeOneRate)
	borrowRate = borrowRate.Add(borrowRate, poolParams.BaseRate)
	return borrowRate, nil
}

func ConvertStringToBigInt(str string) (*big.Int, error) {

	bigint, ok := new(big.Int).SetString(str, 10)
	if !ok {
		return nil, errorsmod.Wrap(
			types.ErrUnableToDecodeBigInt,
			"Unable to convert the sDAI conversion rate to a big int",
		)
	}

	return bigint, nil
}
