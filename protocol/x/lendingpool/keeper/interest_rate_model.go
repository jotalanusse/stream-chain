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
		baseRate, err := ConvertStringToBigInt(poolParams.BaseRate)
		if err != nil {
			return nil, err
		}
		return baseRate, nil
	}

	//      expectedLiquidity - availableLiquidity
	// U = -------------------------------------
	//             expectedLiquidity

	utilisation_eighteen := new(big.Int).Mul(big.NewInt(types.EIGHTEEN_DECIMALS), new(big.Int).Sub(expectedLiquidity, availableLiquidity))
	utilisation_eighteen = utilisation_eighteen.Div(utilisation_eighteen, expectedLiquidity)

	U_Optimal_eighteen, _ := big.NewInt(0).SetString(poolParams.OptimalUtilizationRatio, 10)
	R_base_twentyseven, _ := big.NewInt(0).SetString(poolParams.BaseRate, 10)
	R_slope1_twentyseven, _ := big.NewInt(0).SetString(poolParams.SlopeOneRate, 10)
	R_slope2_twentyseven, _ := big.NewInt(0).SetString(poolParams.SlopeTwoRate, 10)
	U_Optimal_inverted_eighteen := new(big.Int).Sub(big.NewInt(types.EIGHTEEN_DECIMALS), U_Optimal_eighteen)

	// if U < Uoptimal:
	//
	//                                    U
	// borrowRate = Rbase + Rslope1 * ----------
	//                                 Uoptimal
	//

	if utilisation_eighteen.Cmp(U_Optimal_eighteen) < 0 {
		// borrowRate = Rbase + Rslope1 * (U / Uoptimal)
		borrowRate := new(big.Int).Mul(R_slope1_twentyseven, utilisation_eighteen)
		borrowRate = borrowRate.Div(borrowRate, U_Optimal_eighteen)
		borrowRate = borrowRate.Add(borrowRate, R_base_twentyseven)
		return borrowRate, nil
	}

	// if U >= Uoptimal:
	//
	//                                           U - Uoptimal
	// borrowRate = Rbase + Rslope1 + Rslope2 * --------------
	//
	borrowRate := new(big.Int).Sub(utilisation_eighteen, U_Optimal_eighteen)
	borrowRate = borrowRate.Mul(borrowRate, R_slope2_twentyseven)
	borrowRate = borrowRate.Div(borrowRate, U_Optimal_inverted_eighteen)
	borrowRate = borrowRate.Add(borrowRate, R_slope1_twentyseven)
	borrowRate = borrowRate.Add(borrowRate, R_base_twentyseven)
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
