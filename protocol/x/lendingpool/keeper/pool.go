package keeper

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (k Keeper) DepositLiquidity(ctx sdk.Context, amount *big.Int, onBehalfOf sdk.AccAddress, tokenDenom string) error {

	err := k.validateCanAddLiquidity(ctx, amount, onBehalfOf, tokenDenom)
	if err != nil {
		return err
	}

	initialLendingPoolBalance := k.GetLendingPoolBalance(ctx, tokenDenom)

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, onBehalfOf, types.ModuleName, sdk.NewCoins(sdk.NewCoin(tokenDenom, sdkmath.NewIntFromBigInt(amount))))
	if err != nil {
		return errorsmod.Wrap(err, "failed to deposit liquidity into the lending pool")
	}

	newLendingPoolBalance := k.GetLendingPoolBalance(ctx, tokenDenom)
	amount = new(big.Int).Sub(newLendingPoolBalance, initialLendingPoolBalance)

	lendingTokenAmountToMint, err := k.ToLendingToken(ctx, amount, tokenDenom)
	if err != nil {
		return err
	}
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(types.GetLendingTokenDenom(tokenDenom), sdkmath.NewIntFromBigInt(lendingTokenAmountToMint))))
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, onBehalfOf, sdk.NewCoins(sdk.NewCoin(types.GetLendingTokenDenom(tokenDenom), sdkmath.NewIntFromBigInt(lendingTokenAmountToMint))))
	if err != nil {
		return err
	}

	err = k.UpdateExpectedLiquidity(ctx, tokenDenom, amount, false)
	if err != nil {
		return err
	}

	err = k.UpdateBorrowRate(ctx, big.NewInt(0), tokenDenom)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) RemoveLiquidity(ctx sdk.Context, amount *big.Int, onBehalfOf sdk.AccAddress, tokenDenom string) (amountSent *big.Int, err error) {
	// Check if the onBehalfOf address is empty
	if onBehalfOf.Empty() {
		return nil, errorsmod.Wrap(types.ErrInvalidAddress, "onBehalfOf address is empty")
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, onBehalfOf, types.ModuleName, sdk.NewCoins(sdk.NewCoin(types.GetLendingTokenDenom(tokenDenom), sdkmath.NewIntFromBigInt(amount))))
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to withdraw liquidity from the lending pool")
	}

	amountToWithdraw, err := k.FromLendingToken(ctx, amount, tokenDenom)
	if err != nil {
		return nil, err
	}

	amountForInsuranceFund := k.CalculateInsuranceFundFee(ctx, tokenDenom, amountToWithdraw)

	amountToSend := new(big.Int).Sub(amountToWithdraw, amountForInsuranceFund)

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, onBehalfOf, sdk.NewCoins(sdk.NewCoin(tokenDenom, sdkmath.NewIntFromBigInt(amountToSend))))
	if err != nil {
		return nil, err
	}

	if amountForInsuranceFund.Cmp(big.NewInt(0)) > 0 {
		err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.LENDING_POOL_INSURANCE_FUND, sdk.NewCoins(sdk.NewCoin(tokenDenom, sdkmath.NewIntFromBigInt(amountForInsuranceFund))))
		if err != nil {
			return nil, err
		}
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(types.GetLendingTokenDenom(tokenDenom), sdkmath.NewIntFromBigInt(amount))))
	if err != nil {
		return nil, err
	}

	err = k.UpdateExpectedLiquidity(ctx, tokenDenom, amountToWithdraw, true)
	if err != nil {
		return nil, err
	}

	err = k.UpdateBorrowRate(ctx, big.NewInt(0), tokenDenom)
	if err != nil {
		return nil, err
	}
	return amountToSend, nil
}

// Assumes all collateral checks are done by the credit account
func (k Keeper) LendToCreditAccount(ctx sdk.Context, amount *big.Int, creditAccount string, tokenDenom string) error {
	err := k.IsAccountPermissioned(ctx, tokenDenom, creditAccount)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, creditAccount, sdk.NewCoins(sdk.NewCoin(tokenDenom, sdkmath.NewIntFromBigInt(amount))))
	if err != nil {
		return errorsmod.Wrap(err, "failed to send funds to credit account")
	}

	err = k.UpdateBorrowRate(ctx, big.NewInt(0), tokenDenom)
	if err != nil {
		return err
	}

	err = k.UpdateTotalBorrowed(ctx, tokenDenom, amount, false)
	if err != nil {
		return err
	}

	return nil
}

// Assumes that the underlying (including principal + interest + fees) was already transferred
func (k Keeper) RepayFromCreditAccount(ctx sdk.Context, borrowedAmount *big.Int, profit *big.Int, loss *big.Int, creditAccount string, tokenDenom string) error {

	if profit.Cmp(big.NewInt(0)) > 0 && loss.Cmp(big.NewInt(0)) > 0 {
		return errorsmod.Wrap(types.ErrInvalidRepayFromCreditAccount, "both profit and loss cannot be greater than zero")
	}

	err := k.IsAccountPermissioned(ctx, tokenDenom, creditAccount)
	if err != nil {
		return err
	}

	err = k.HandleProfit(ctx, profit, tokenDenom)
	if err != nil {
		return err
	}

	err = k.HandleLoss(ctx, loss, tokenDenom)
	if err != nil {
		return err
	}

	err = k.UpdateBorrowRate(ctx, loss, tokenDenom)
	if err != nil {
		return err
	}

	err = k.UpdateTotalBorrowed(ctx, tokenDenom, borrowedAmount, true)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) validateCanAddLiquidity(ctx sdk.Context, amount *big.Int, onBehalfOf sdk.AccAddress, tokenDenom string) error {
	// Check if the onBehalfOf address is empty
	if onBehalfOf.Empty() {
		return errorsmod.Wrap(types.ErrInvalidAddress, "onBehalfOf address is empty")
	}

	poolParams, ok := k.GetPoolParams(ctx, tokenDenom)
	if !ok {
		return errorsmod.Wrap(types.ErrInvalidTokenDenom, "pool params not found")
	}

	PoolLiquidityLimit, err := ConvertStringToBigInt(poolParams.MaxPoolLiquidity)
	if err != nil {
		return err
	}

	err = k.ValidateExpectedLiquidity(ctx, tokenDenom, amount, PoolLiquidityLimit)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) CalculateExpectedLiquidity(ctx sdk.Context, tokenDenom string) (*big.Int, error) {

	timeDifference, err := k.GetTimeDifference(ctx, tokenDenom)
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

	interestAccrued := k.CalculateInterestAccrued(totalBorrowed, borrowAPY, timeDifference)

	expectedLiquidityLU, found := k.GetLastUpdatedTotalLiquidity(ctx, tokenDenom)
	if !found {
		return nil, errorsmod.Wrap(types.ErrInvalidTokenDenom, "expected liquidity last updated not found")
	}

	// Calculate the expected liquidity
	expectedLiquidity := new(big.Int).Add(expectedLiquidityLU, interestAccrued)

	return expectedLiquidity, nil
}

//	currentBorrowRate * timeDifference
//
// interestAccrued = totalBorrow *  currentBorrowRate * timeDifference / SECONDS_PER_YEAR
func (k Keeper) CalculateInterestAccrued(totalBorrowed, borrowAPY *big.Int, timeDifference uint64) *big.Int {
	interestAccrued := new(big.Int).Mul(totalBorrowed, borrowAPY)
	interestAccrued = interestAccrued.Mul(interestAccrued, big.NewInt(int64(timeDifference)))
	interestAccrued = interestAccrued.Div(interestAccrued, big.NewInt(types.TWENTY_SEVEN_DECIMALS))
	interestAccrued = interestAccrued.Div(interestAccrued, big.NewInt(types.SECONDS_PER_YEAR))
	return interestAccrued
}

func (k Keeper) GetTimeDifference(ctx sdk.Context, tokenDenom string) (uint64, error) {
	// Get the last updated time
	lastUpdatedTime, found := k.GetLastUpdatedTime(ctx, tokenDenom)
	if !found {
		return 0, errorsmod.Wrap(types.ErrInvalidTokenDenom, "last updated time not found")
	}

	// Calculate the time difference
	timeDifference := uint64(ctx.BlockTime().Unix()) - lastUpdatedTime
	return timeDifference, nil
}

func (k Keeper) ValidateExpectedLiquidity(ctx sdk.Context, tokenDenom string, amount *big.Int, expectedLiquidityLimit *big.Int) error {
	// Get the expected liquidity
	expectedLiquidity, err := k.CalculateExpectedLiquidity(ctx, tokenDenom)
	if err != nil {
		return err
	}

	// Check if the expected liquidity plus the amount is within the expected liquidity limit
	if new(big.Int).Add(expectedLiquidity, amount).Cmp(expectedLiquidityLimit) > 0 {
		return errorsmod.Wrap(types.ErrPoolMoreThanMaxLiquidityLimit, "expected liquidity would exceed the limit if deposit")
	}

	return nil
}

func (k Keeper) GetLendingPoolBalance(ctx sdk.Context, denom string) *big.Int {
	moduleAddress := authtypes.NewModuleAddress(types.ModuleName)
	balance := k.bankKeeper.GetBalance(ctx, moduleAddress, denom)
	return balance.Amount.BigInt()
}

func (k Keeper) GetInsuranceFundBalance(ctx sdk.Context, denom string) *big.Int {
	moduleAddress := authtypes.NewModuleAddress(types.LENDING_POOL_INSURANCE_FUND)
	balance := k.bankKeeper.GetBalance(ctx, moduleAddress, denom)
	return balance.Amount.BigInt()
}

func (k Keeper) GetTotalSupply(ctx sdk.Context, denom string) (*big.Int, error) {
	// Get the total supply of the lending pool for the given denom
	totalSupply := k.bankKeeper.GetSupply(ctx, denom)
	return totalSupply.Amount.BigInt(), nil
}

func (k Keeper) ToLendingToken(ctx sdk.Context, amount *big.Int, denom string) (*big.Int, error) {
	exchangeRate, err := k.GetLendingTokenExchangeRate(ctx, types.GetLendingTokenDenom(denom))
	if err != nil {
		return nil, err
	}
	lendingToken := new(big.Int).Mul(amount, big.NewInt(types.TWENTY_SEVEN_DECIMALS))
	lendingToken = lendingToken.Div(lendingToken, exchangeRate)
	return lendingToken, nil
}

func (k Keeper) FromLendingToken(ctx sdk.Context, amount *big.Int, denom string) (*big.Int, error) {
	exchangeRate, err := k.GetLendingTokenExchangeRate(ctx, types.GetLendingTokenDenom(denom))
	if err != nil {
		return nil, err
	}
	token := new(big.Int).Mul(amount, exchangeRate)
	token = token.Div(token, big.NewInt(types.TWENTY_SEVEN_DECIMALS))
	return token, nil
}

func (k Keeper) GetLendingTokenExchangeRate(ctx sdk.Context, denom string) (*big.Int, error) {

	lendingTokenSupply, err := k.GetTotalSupply(ctx, denom)
	if err != nil {
		return nil, err
	}
	if lendingTokenSupply.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(types.TWENTY_SEVEN_DECIMALS), nil
	}

	expectedLiquidity, err := k.CalculateExpectedLiquidity(ctx, denom)
	if err != nil {
		return nil, err
	}

	exchangeRate := new(big.Int).Mul(expectedLiquidity, big.NewInt(types.TWENTY_SEVEN_DECIMALS))
	exchangeRate = exchangeRate.Div(exchangeRate, lendingTokenSupply)
	return exchangeRate, nil
}

func (k Keeper) UpdateExpectedLiquidity(ctx sdk.Context, tokenDenom string, amount *big.Int, isNegative bool) error {

	// Get the expected liquidity last updated
	expectedLiquidityLU, found := k.GetLastUpdatedTotalLiquidity(ctx, tokenDenom)
	if !found {
		return errorsmod.Wrap(types.ErrInvalidTokenDenom, "expected liquidity last updated not found")
	}

	newExpectedLiquidity := new(big.Int)

	if isNegative {
		newExpectedLiquidity = new(big.Int).Sub(expectedLiquidityLU, amount)
	} else {
		newExpectedLiquidity = new(big.Int).Add(expectedLiquidityLU, amount)
	}

	// Set the expected liquidity last updated
	k.SetLastUpdatedTotalLiquidity(ctx, tokenDenom, newExpectedLiquidity)

	return nil
}

func (k Keeper) UpdateTotalBorrowed(ctx sdk.Context, tokenDenom string, amount *big.Int, isNegative bool) error {

	// Get the expected liquidity last updated
	totalBorrowed, found := k.GetTotalBorrowed(ctx, tokenDenom)
	if !found {
		return errorsmod.Wrap(types.ErrInvalidTokenDenom, "total borrowed not found")
	}

	newTotalBorrowed := new(big.Int)

	if isNegative {
		newTotalBorrowed = new(big.Int).Sub(totalBorrowed, amount)
	} else {
		newTotalBorrowed = new(big.Int).Add(totalBorrowed, amount)
	}

	// Set the expected liquidity last updated
	k.SetTotalBorrowed(ctx, tokenDenom, newTotalBorrowed)

	return nil
}

func (k Keeper) UpdateBorrowRate(ctx sdk.Context, loss *big.Int, tokenDenom string) error {
	// Get the expected liquidity last updated
	expectedLiquidity, err := k.CalculateExpectedLiquidity(ctx, tokenDenom)
	if err != nil {
		return err
	}

	// Update total expected liquidity
	expectedLiquidity = new(big.Int).Sub(expectedLiquidity, loss)

	// Set the expected liquidity last updated
	k.SetLastUpdatedTotalLiquidity(ctx, tokenDenom, expectedLiquidity)

	// Update cumulative index
	cumulativeIndex, err := k.CalculateLinearCumulative(ctx, tokenDenom)
	if err != nil {
		return err
	}
	k.SetCumulativeInterestRate(ctx, tokenDenom, cumulativeIndex)

	// Update borrow APY
	availableLiquidity := k.GetLendingPoolBalance(ctx, tokenDenom)
	borrowAPY, err := k.CalcBorrowRate(ctx, tokenDenom, expectedLiquidity, availableLiquidity)
	if err != nil {
		return err
	}
	k.SetCurrentBorrowAPY(ctx, tokenDenom, borrowAPY)

	// Update the last updated timestamp
	k.SetLastUpdatedTime(ctx, tokenDenom, uint64(ctx.BlockTime().Unix()))

	return nil
}

func (k Keeper) CalculateLinearCumulative(ctx sdk.Context, tokenDenom string) (*big.Int, error) {

	// Calculate the time difference
	timeDifference, err := k.GetTimeDifference(ctx, tokenDenom)
	if err != nil {
		return nil, err
	}

	// Get the cumulative interest rate
	cumulativeIndex, found := k.GetCumulativeInterestRate(ctx, tokenDenom)
	if !found {
		return nil, errorsmod.Wrap(types.ErrInvalidTokenDenom, "cumulative interest rate not found")
	}

	// Get the current borrow APY
	borrowAPY, found := k.GetCurrentBorrowAPY(ctx, tokenDenom)
	if !found {
		return nil, errorsmod.Wrap(types.ErrInvalidTokenDenom, "current borrow APY not found")
	}

	// Calculate the linear cumulative index
	linearCumulativeIndex := k.CalculateNewCumulativeIndex(cumulativeIndex, borrowAPY, timeDifference)

	return linearCumulativeIndex, nil
}

// newIndex  = currentIndex * (1 + borrowAPY * timeDifference / SECONDS_PER_YEAR)
func (k Keeper) CalculateNewCumulativeIndex(cumulativeIndex, borrowAPY *big.Int, timeDifference uint64) *big.Int {
	// Calculate the linear index
	linearIndex := new(big.Int).Mul(big.NewInt(int64(timeDifference)), borrowAPY)
	linearIndex = linearIndex.Div(linearIndex, big.NewInt(types.SECONDS_PER_YEAR))
	linearIndex = linearIndex.Add(linearIndex, big.NewInt(types.TWENTY_SEVEN_DECIMALS))
	linearIndex = linearIndex.Mul(linearIndex, cumulativeIndex)
	linearIndex = linearIndex.Div(linearIndex, big.NewInt(types.TWENTY_SEVEN_DECIMALS))
	return linearIndex
}

func (k Keeper) CalculateInsuranceFundFee(ctx sdk.Context, tokenDenom string, underlyingTokensAmount *big.Int) *big.Int {
	// Get the pool parameters
	poolParams, found := k.GetPoolParams(ctx, tokenDenom)
	if !found {
		return big.NewInt(0) // or handle the error appropriately
	}

	// Get the withdraw fee from the pool parameters
	withdrawFee, err := ConvertStringToBigInt(poolParams.WithdrawFee)
	if err != nil {
		return big.NewInt(0) // or handle the error appropriately
	}

	// Calculate the insurance fund fee
	insuranceFundFee := types.PercentMul(underlyingTokensAmount, withdrawFee)

	return insuranceFundFee
}

func (k Keeper) IsAccountPermissioned(ctx sdk.Context, tokenDenom string, account string) error {
	// Get the pool parameters
	poolParams, found := k.GetPoolParams(ctx, tokenDenom)
	if !found {
		return errorsmod.Wrap(types.ErrInvalidTokenDenom, "pool parameters not found")
	}

	// Check if the account is in the permissioned credit accounts
	for _, permissionedAccount := range poolParams.PermissionedCreditAccounts {
		if permissionedAccount == account {
			return nil
		}
	}

	return types.ErrCreditAccountNotPermissioned
}

func (k Keeper) HandleProfit(ctx sdk.Context, profit *big.Int, tokenDenom string) error {

	if profit.Cmp(big.NewInt(0)) > 0 {
		lendingTokenAmountToMint, err := k.ToLendingToken(ctx, profit, tokenDenom)
		if err != nil {
			return err
		}
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(types.GetLendingTokenDenom(tokenDenom), sdkmath.NewIntFromBigInt(lendingTokenAmountToMint))))
		if err != nil {
			return err
		}

		err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, types.LENDING_POOL_INSURANCE_FUND, sdk.NewCoins(sdk.NewCoin(types.GetLendingTokenDenom(tokenDenom), sdkmath.NewIntFromBigInt(lendingTokenAmountToMint))))
		if err != nil {
			return err
		}

		err = k.UpdateExpectedLiquidity(ctx, tokenDenom, profit, false)
		if err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) HandleLoss(ctx sdk.Context, loss *big.Int, tokenDenom string) error {

	if loss.Cmp(big.NewInt(0)) > 0 {
		lendingTokenAmountToBurn, err := k.ToLendingToken(ctx, loss, tokenDenom)
		if err != nil {
			return err
		}

		insuranceFundBalance := k.GetInsuranceFundBalance(ctx, tokenDenom)
		if insuranceFundBalance.Cmp(lendingTokenAmountToBurn) < 0 {
			lendingTokenAmountToBurn = insuranceFundBalance
		}

		k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.LENDING_POOL_INSURANCE_FUND, types.ModuleName, sdk.NewCoins(sdk.NewCoin(types.GetLendingTokenDenom(tokenDenom), sdkmath.NewIntFromBigInt(lendingTokenAmountToBurn))))

		k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(types.GetLendingTokenDenom(tokenDenom), sdkmath.NewIntFromBigInt(lendingTokenAmountToBurn))))
	}

	return nil
}
