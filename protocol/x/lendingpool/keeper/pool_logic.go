package keeper

import (
	"math/big"
	"slices"

	errorsmod "cosmossdk.io/errors"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (k Keeper) validateCanAddLiquidity(ctx sdk.Context, amount *big.Int, tokenDenom string) error {

	poolParams, ok := k.GetPoolParams(ctx, tokenDenom)
	if !ok {
		return errorsmod.Wrap(types.ErrInvalidTokenDenom, "pool params not found")
	}

	err := k.CheckPoolLiquidityDoesNotExceedLimit(ctx, tokenDenom, amount, poolParams.MaxPoolLiquidity)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) CheckPoolLiquidityDoesNotExceedLimit(ctx sdk.Context, tokenDenom string, amount *big.Int, totalLiquidityLimit *big.Int) error {

	totalLiquidity, err := k.CalculateTotalLiquidityIncludingAccruedInterest(ctx, tokenDenom)
	if err != nil {
		return err
	}

	if new(big.Int).Add(totalLiquidity, amount).Cmp(totalLiquidityLimit) > 0 {
		return errorsmod.Wrap(types.ErrPoolMoreThanMaxLiquidityLimit, "total liquidity would exceed the limit if deposit")
	}

	return nil
}

func (k Keeper) CalculateTotalLiquidityIncludingAccruedInterest(ctx sdk.Context, tokenDenom string) (totalLiquidity *big.Int, err error) {

	timeSinceYieldLastClaimed, err := k.GetTimeSinceYieldLastClaimed(ctx, tokenDenom)
	if err != nil {
		return nil, err
	}

	totalBorrowed, found := k.GetTotalBorrowed(ctx, tokenDenom)
	if !found {
		return nil, errorsmod.Wrap(types.ErrInvalidTokenDenom, "total borrowed not found")
	}

	borrowAPY, found := k.GetCurrentBorrowAPY(ctx, tokenDenom)
	if !found {
		return nil, errorsmod.Wrap(types.ErrInvalidTokenDenom, "current borrow APY not found")
	}

	lastUpdatedTotalLiquidity, found := k.GetLastUpdatedTotalLiquidity(ctx, tokenDenom)
	if !found {
		return nil, errorsmod.Wrap(types.ErrInvalidTokenDenom, "expected liquidity last updated not found")
	}

	interestAccrued := CalculateInterestAccrued(totalBorrowed, borrowAPY, timeSinceYieldLastClaimed)

	return new(big.Int).Add(lastUpdatedTotalLiquidity, interestAccrued), nil
}

func CalculateInterestAccrued(totalBorrowed, borrowAPY *big.Int, timeDifference *big.Int) (interestAccrued *big.Int) {

	//                                    currentBorrowRate * timeDifference
	//  interestAccrued = totalBorrow *  ------------------------------------
	//                                             SECONDS_PER_YEAR
	interestAccrued = new(big.Int).Mul(totalBorrowed, borrowAPY)
	interestAccrued = interestAccrued.Mul(interestAccrued, timeDifference)
	interestAccrued = interestAccrued.Div(interestAccrued, types.TWENTY_SEVEN_DECIMALS)
	return interestAccrued.Div(interestAccrued, types.SECONDS_PER_YEAR)
}

func (k Keeper) GetTimeSinceYieldLastClaimed(ctx sdk.Context, tokenDenom string) (timeSinceYieldLastClaimed *big.Int, err error) {

	lastUpdatedTime, found := k.GetLastUpdatedTime(ctx, tokenDenom)
	if !found {
		return nil, errorsmod.Wrap(types.ErrInvalidTokenDenom, "last updated time not found")
	}

	return new(big.Int).Sub(big.NewInt(ctx.BlockTime().Unix()), lastUpdatedTime), nil
}

func (k Keeper) ConvertBaseToLendingToken(ctx sdk.Context, baseAmount *big.Int, denom string) (lendingAmount *big.Int, err error) {
	exchangeRate, err := k.GetLendingTokenExchangeRate(ctx, denom)
	if err != nil {
		return nil, err
	}

	lendingAmount = new(big.Int).Mul(baseAmount, types.TWENTY_SEVEN_DECIMALS)
	lendingAmount = lendingAmount.Div(lendingAmount, exchangeRate)
	return lendingAmount, nil
}

func (k Keeper) ConvertLendingTokenToBase(ctx sdk.Context, lendingAmount *big.Int, denom string) (baseAmount *big.Int, err error) {
	exchangeRate, err := k.GetLendingTokenExchangeRate(ctx, denom)
	if err != nil {
		return nil, err
	}
	baseAmount = new(big.Int).Mul(lendingAmount, exchangeRate)
	baseAmount = baseAmount.Div(baseAmount, types.TWENTY_SEVEN_DECIMALS)
	return baseAmount, nil
}

func (k Keeper) GetLendingTokenExchangeRate(ctx sdk.Context, denom string) (exchangeRate *big.Int, err error) {

	lendingTokenSupply := k.bankKeeper.GetSupply(ctx, types.GetLendingTokenDenom(denom)).Amount.BigInt()

	if lendingTokenSupply.Cmp(big.NewInt(0)) == 0 {
		return types.TWENTY_SEVEN_DECIMALS, nil
	}

	totalLiquidity, err := k.CalculateTotalLiquidityIncludingAccruedInterest(ctx, denom)
	if err != nil {
		return nil, err
	}

	exchangeRate = new(big.Int).Mul(totalLiquidity, types.TWENTY_SEVEN_DECIMALS)
	exchangeRate = exchangeRate.Div(exchangeRate, lendingTokenSupply)
	return exchangeRate, nil
}

func (k Keeper) UpdateTotalLiquidity(ctx sdk.Context, tokenDenom string, amount *big.Int) error {

	totalLiquidityLU, found := k.GetLastUpdatedTotalLiquidity(ctx, tokenDenom)
	if !found {
		return errorsmod.Wrap(types.ErrInvalidTokenDenom, "total liquidity last updated not found")
	}

	k.SetLastUpdatedTotalLiquidity(ctx, tokenDenom, new(big.Int).Add(totalLiquidityLU, amount))
	return nil
}

func (k Keeper) UpdateTotalBorrowed(ctx sdk.Context, tokenDenom string, amount *big.Int) error {

	totalBorrowed, found := k.GetTotalBorrowed(ctx, tokenDenom)
	if !found {
		return errorsmod.Wrap(types.ErrInvalidTokenDenom, "total borrowed not found")
	}

	k.SetTotalBorrowed(ctx, tokenDenom, new(big.Int).Add(totalBorrowed, amount))
	return nil
}

func (k Keeper) UpdatePool(ctx sdk.Context, loss *big.Int, tokenDenom string) error {

	totalLiquidity, err := k.CalculateTotalLiquidityIncludingAccruedInterest(ctx, tokenDenom)
	if err != nil {
		return err
	}
	totalLiquidity = new(big.Int).Sub(totalLiquidity, loss)
	availableLiquidity := k.GetLendingPoolBalance(ctx, tokenDenom)

	borrowAPY, err := k.CalculateBorrowRate(ctx, tokenDenom, totalLiquidity, availableLiquidity)
	if err != nil {
		return err
	}

	k.SetLastUpdatedTotalLiquidity(ctx, tokenDenom, totalLiquidity)
	k.SetCurrentBorrowAPY(ctx, tokenDenom, borrowAPY)
	k.SetLastUpdatedTime(ctx, tokenDenom, big.NewInt(ctx.BlockTime().Unix()))
	return nil
}

func (k Keeper) UpdateCumulativeBorrowIndex(ctx sdk.Context, tokenDenom string) error {
	newBorrowIndex, err := k.CalculateCumulativeBorrowIndex(ctx, tokenDenom)
	if err != nil {
		return err
	}
	k.SetCumulativeInterestRate(ctx, tokenDenom, newBorrowIndex)
	return nil
}

func (k Keeper) CalculateCumulativeBorrowIndex(ctx sdk.Context, tokenDenom string) (newBorrowIndex *big.Int, err error) {

	timeSinceYieldLastClaimed, err := k.GetTimeSinceYieldLastClaimed(ctx, tokenDenom)
	if err != nil {
		return nil, err
	}

	cumulativeIndex, found := k.GetCumulativeInterestRate(ctx, tokenDenom)
	if !found {
		return nil, errorsmod.Wrap(types.ErrInvalidTokenDenom, "cumulative interest rate not found")
	}

	borrowAPY, found := k.GetCurrentBorrowAPY(ctx, tokenDenom)
	if !found {
		return nil, errorsmod.Wrap(types.ErrInvalidTokenDenom, "current borrow APY not found")
	}
	return CalculateCumulativeBorrowIndex(cumulativeIndex, borrowAPY, timeSinceYieldLastClaimed), nil
}

func CalculateCumulativeBorrowIndex(currentBorrowIndex, currentBorrowRate, timeSinceYieldLastClaimed *big.Int) (newBorrowIndex *big.Int) {

	//                                          /     currentBorrowRate * timeSinceYieldLastClaimed \
	//  newBorrowIndex  = currentBorrowIndex * | 1 + ----------------------------------------------- |
	//                                          \                    SECONDS_PER_YEAR               /
	newBorrowIndex = new(big.Int).Mul(timeSinceYieldLastClaimed, currentBorrowRate)
	newBorrowIndex = newBorrowIndex.Div(newBorrowIndex, types.SECONDS_PER_YEAR)
	newBorrowIndex = newBorrowIndex.Add(newBorrowIndex, types.TWENTY_SEVEN_DECIMALS)
	newBorrowIndex = newBorrowIndex.Mul(newBorrowIndex, currentBorrowIndex)
	return newBorrowIndex.Div(newBorrowIndex, types.TWENTY_SEVEN_DECIMALS)

}

func (k Keeper) IsAccountPermissioned(ctx sdk.Context, tokenDenom string, account string) error {
	poolParams, found := k.GetPoolParams(ctx, tokenDenom)
	if !found {
		return errorsmod.Wrap(types.ErrInvalidTokenDenom, "pool parameters not found")
	}

	if slices.Contains(poolParams.PermissionedCreditAccounts, account) {
		return nil
	}

	return types.ErrCreditAccountNotPermissioned
}

func (k Keeper) GetLendingPoolBalance(ctx sdk.Context, denom string) *big.Int {

	moduleAddress := authtypes.NewModuleAddress(types.ModuleName)
	return k.bankKeeper.GetBalance(ctx, moduleAddress, denom).Amount.BigInt()
}

func (k Keeper) GetInsuranceFundBalance(ctx sdk.Context, denom string, insuranceFund string) *big.Int {

	moduleAddress := authtypes.NewModuleAddress(insuranceFund)
	return k.bankKeeper.GetBalance(ctx, moduleAddress, denom).Amount.BigInt()
}

func (k Keeper) GetInsuranceFundForPool(ctx sdk.Context, baseDenom string) (insuranceFund string, err error) {

	poolParams, found := k.GetPoolParams(ctx, baseDenom)
	if !found {
		return "", types.ErrInvalidTokenDenom
	}

	if poolParams.IsIsolated {
		insuranceFund = types.LENDING_POOL_INSURANCE_FUND + baseDenom
	} else {
		insuranceFund = types.LENDING_POOL_INSURANCE_FUND
	}
	return insuranceFund, nil
}
