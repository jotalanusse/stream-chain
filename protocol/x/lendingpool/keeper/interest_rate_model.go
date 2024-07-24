package keeper

import (
	"math/big"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CalculateBorrowRate(ctx sdk.Context, tokenDenom string, totalLiquidity, availableLiquidity *big.Int) (*big.Int, error) {
	poolParams, found := k.GetPoolParams(ctx, tokenDenom)
	if !found {
		return nil, types.ErrInvalidTokenDenom
	}

	if totalLiquidity.Cmp(big.NewInt(0)) == 0 || totalLiquidity.Cmp(availableLiquidity) < 0 {
		return poolParams.BaseRate, nil
	}

	utilisation := k.GetUtilisation(totalLiquidity, availableLiquidity)

	if utilisation.Cmp(poolParams.OptimalUtilizationRatio) < 0 {
		return k.getBorrowRateUnderOptimalUtilization(poolParams, utilisation), nil
	}

	return k.getBorrowRateAboveOptimalUtilization(poolParams, utilisation), nil
}

func (k Keeper) GetUtilisation(totalLiquidity, availableLiquidity *big.Int) *big.Int {

	//      totalLiquidity - availableLiquidity
	// U = -------------------------------------
	//             totalLiquidity
	utilisation := new(big.Int).Mul(types.EIGHTEEN_DECIMALS, new(big.Int).Sub(totalLiquidity, availableLiquidity))
	utilisation = utilisation.Div(utilisation, totalLiquidity)

	return utilisation
}

func (k Keeper) getBorrowRateUnderOptimalUtilization(poolParams types.InternalPoolParams, utilisation *big.Int) *big.Int {

	//                                    U
	// borrowRate = Rbase + Rslope1 * ----------
	//                                 Uoptimal
	borrowRate := new(big.Int).Mul(poolParams.SlopeOneRate, utilisation)
	borrowRate = borrowRate.Div(borrowRate, poolParams.OptimalUtilizationRatio)
	borrowRate = borrowRate.Add(borrowRate, poolParams.BaseRate)
	return borrowRate
}

func (k Keeper) getBorrowRateAboveOptimalUtilization(poolParams types.InternalPoolParams, utilisation *big.Int) *big.Int {

	//                                           U - Uoptimal
	// borrowRate = Rbase + Rslope1 + Rslope2 * --------------
	//                                           1 - Uoptimal
	borrowRate := new(big.Int).Sub(utilisation, poolParams.OptimalUtilizationRatio)
	borrowRate = borrowRate.Mul(borrowRate, poolParams.SlopeTwoRate)
	borrowRate = borrowRate.Div(borrowRate, new(big.Int).Sub(types.EIGHTEEN_DECIMALS, poolParams.OptimalUtilizationRatio))
	borrowRate = borrowRate.Add(borrowRate, poolParams.SlopeOneRate)
	borrowRate = borrowRate.Add(borrowRate, poolParams.BaseRate)
	return borrowRate
}
