package keeper

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) DepositLiquidity(ctx sdk.Context, amount *big.Int, liquidityProvider sdk.AccAddress, tokenDenom string) error {

	err := k.validateCanAddLiquidity(ctx, amount, tokenDenom)
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, liquidityProvider, types.ModuleName, sdk.NewCoins(sdk.NewCoin(tokenDenom, sdkmath.NewIntFromBigInt(amount))))
	if err != nil {
		return errorsmod.Wrap(err, "failed to deposit liquidity into the lending pool")
	}

	lendingTokenAmountToMint, err := k.ConvertBaseToLendingToken(ctx, amount, tokenDenom)
	if err != nil {
		return err
	}
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(types.GetLendingTokenDenom(tokenDenom), sdkmath.NewIntFromBigInt(lendingTokenAmountToMint))))
	if err != nil {
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, liquidityProvider, sdk.NewCoins(sdk.NewCoin(types.GetLendingTokenDenom(tokenDenom), sdkmath.NewIntFromBigInt(lendingTokenAmountToMint))))
	if err != nil {
		return err
	}

	err = k.UpdateTotalLiquidity(ctx, tokenDenom, amount)
	if err != nil {
		return err
	}

	err = k.UpdateCumulativeBorrowIndex(ctx, tokenDenom)
	if err != nil {
		return err
	}

	err = k.UpdatePool(ctx, big.NewInt(0), tokenDenom)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) RemoveLiquidity(ctx sdk.Context, lendingTokenAmount *big.Int, liquidityProvider sdk.AccAddress, tokenDenom string) (amountSent *big.Int, err error) {

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, liquidityProvider, types.ModuleName, sdk.NewCoins(sdk.NewCoin(types.GetLendingTokenDenom(tokenDenom), sdkmath.NewIntFromBigInt(lendingTokenAmount))))
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to withdraw liquidity from the lending pool")
	}

	amountToWithdraw, err := k.ConvertLendingTokenToBase(ctx, lendingTokenAmount, tokenDenom)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, liquidityProvider, sdk.NewCoins(sdk.NewCoin(tokenDenom, sdkmath.NewIntFromBigInt(amountToWithdraw))))
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(types.GetLendingTokenDenom(tokenDenom), sdkmath.NewIntFromBigInt(lendingTokenAmount))))
	if err != nil {
		return nil, err
	}

	err = k.UpdateTotalLiquidity(ctx, tokenDenom, new(big.Int).Neg(amountToWithdraw))
	if err != nil {
		return nil, err
	}

	err = k.UpdateCumulativeBorrowIndex(ctx, tokenDenom)
	if err != nil {
		return nil, err
	}

	err = k.UpdatePool(ctx, big.NewInt(0), tokenDenom)
	if err != nil {
		return nil, err
	}
	return amountToWithdraw, nil
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

	err = k.UpdateCumulativeBorrowIndex(ctx, tokenDenom)
	if err != nil {
		return err
	}

	err = k.UpdatePool(ctx, big.NewInt(0), tokenDenom)
	if err != nil {
		return err
	}

	err = k.UpdateTotalBorrowed(ctx, tokenDenom, amount)
	if err != nil {
		return err
	}

	return nil
}

// Assumes that the underlying (including principal + interest + fees) was already transferred
// Profit will be included as a fee surplus
func (k Keeper) RepayFromCreditAccount(ctx sdk.Context, borrowedAmount *big.Int, profit *big.Int, loss *big.Int, creditAccount string, tokenDenom string) error {

	err := k.IsAccountPermissioned(ctx, tokenDenom, creditAccount)
	if err != nil {
		return err
	}

	if profit.Cmp(big.NewInt(0)) > 0 {

		err = k.HandleProfit(ctx, profit, tokenDenom)
		if err != nil {
			return err
		}
	} else if loss.Cmp(big.NewInt(0)) > 0 {
		err = k.HandleLoss(ctx, loss, tokenDenom)
		if err != nil {
			return err
		}
	}

	err = k.UpdateCumulativeBorrowIndex(ctx, tokenDenom)
	if err != nil {
		return err
	}

	err = k.UpdatePool(ctx, loss, tokenDenom)
	if err != nil {
		return err
	}

	err = k.UpdateTotalBorrowed(ctx, tokenDenom, new(big.Int).Neg(borrowedAmount))
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) HandleProfit(ctx sdk.Context, profit *big.Int, tokenDenom string) error {

	lendingTokenAmountToMint, err := k.ConvertBaseToLendingToken(ctx, profit, tokenDenom)
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

	err = k.UpdateTotalLiquidity(ctx, tokenDenom, profit)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) HandleLoss(ctx sdk.Context, loss *big.Int, tokenDenom string) error {

	lendingTokenAmountToBurn, err := k.ConvertBaseToLendingToken(ctx, loss, tokenDenom)
	if err != nil {
		return err
	}

	insuranceFundBalance := k.GetInsuranceFundBalance(ctx, types.GetLendingTokenDenom(tokenDenom))
	if insuranceFundBalance.Cmp(lendingTokenAmountToBurn) < 0 {
		lendingTokenAmountToBurn = insuranceFundBalance
	}

	k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.LENDING_POOL_INSURANCE_FUND, types.ModuleName, sdk.NewCoins(sdk.NewCoin(types.GetLendingTokenDenom(tokenDenom), sdkmath.NewIntFromBigInt(lendingTokenAmountToBurn))))

	k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(types.GetLendingTokenDenom(tokenDenom), sdkmath.NewIntFromBigInt(lendingTokenAmountToBurn))))
	return nil
}
