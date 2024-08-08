package keeper

import (
	"fmt"
	"math/big"

	types "github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingaccount/types"
	lendingPoolTypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/types"
	subaccounttypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/subaccounts/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (k Keeper) openCreditAccountOnManager(ctx sdk.Context, managerName string, lendingAccountId uint32, address string, borrowedAmount *big.Int, collateralAssetIds []uint32, amounts []*big.Int) error {

	err := k.CheckLendingAccountExistsAndAddToAddressMapping(ctx, managerName, lendingAccountId, address)
	if err != nil {
		return err
	}

	err = k.CheckAddressNotInUseAndAddToSybilResistanceMapping(ctx, managerName, address)
	if err != nil {
		return err
	}

	lendingManager, found := k.GetLendingManager(ctx, managerName)
	if !found {
		return fmt.Errorf("lending manager with name %s not found", managerName)
	}

	if lendingManager.IsFrozen {
		return fmt.Errorf("lending manager is frozen")
	}

	baseAsset, found := k.assetsKeeper.GetAsset(ctx, lendingManager.AssetId)
	if !found {
		return fmt.Errorf("base asset with id %d not found", lendingManager.AssetId)
	}

	borrowIndex, err := k.lendingPoolKeeper.CalculateCumulativeBorrowIndex(ctx, baseAsset.Denom)
	if err != nil {
		return err
	}

	subaccountId := subaccounttypes.SubaccountId{
		Owner:  authtypes.NewModuleAddress(managerName).String(),
		Number: lendingAccountId,
	}

	lendingAccount := types.LendingAccount{
		BorrowedAmount:     borrowedAmount.String(),
		InitialBorrowIndex: borrowIndex.String(),
		SubaccountId:       subaccountId,
	}

	k.SetLendingAccount(ctx, managerName, lendingAccountId, lendingAccount)

	err = k.lendingPoolKeeper.LendToCreditAccount(ctx, borrowedAmount, address, baseAsset.Denom)
	if err != nil {
		return err
	}

	err = k.subaccountsKeeper.DepositFundsFromAccountToSubaccount(ctx, authtypes.NewModuleAddress(managerName), subaccountId, lendingManager.AssetId, borrowedAmount)
	if err != nil {
		return err
	}

	for i, collateralAssetId := range collateralAssetIds {
		err = k.subaccountsKeeper.DepositFundsFromAccountToSubaccount(ctx, sdk.AccAddress(address), subaccountId, collateralAssetId, amounts[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) closeCreditAccountOnManager(ctx sdk.Context, managerName string, lendingAccount types.LendingAccount, isLiquidation bool, portfolioValueForLiquidation *big.Int, owner sdk.Address, payer sdk.Address) error {

	// check if perpetual positions are open
	// if they are we sell them through the liquidation flow

	// next we figure out the total value of what needs to be payed (depends on if is liquidation or not)
}

func (k Keeper) checkCollateralOfPosition(ctx sdk.Context, managerName string, collateralAssetIds []uint32, amounts []*big.Int, borrowedAmountWithInterestAndFees *big.Int) error {

	lendingManager, found := k.GetLendingManager(ctx, managerName)
	if !found {
		return fmt.Errorf("lending manager with name %s not found", managerName)
	}

	usdValueOfDebt, err := k.pricesKeeper.GetUsdValue(ctx, lendingManager.AssetId, borrowedAmountWithInterestAndFees.Mul(borrowedAmountWithInterestAndFees, types.PERCENTAGE_FACTOR))
	if err != nil {
		return err
	}

	usdValueOfCollateral := big.NewInt(0)
	for i, collateralAssetId := range collateralAssetIds {

		value, err := k.pricesKeeper.GetUsdValue(ctx, collateralAssetId, amounts[i])
		if err != nil {
			return err
		}

		collateralRatio, found := lendingManager.EnabledCollateralAssets[collateralAssetId]
		if !found {
			return fmt.Errorf("collateral asset with id %d is not allowed", collateralAssetId)
		}

		usdValueOfCollateral.Add(usdValueOfCollateral, value.Mul(value, big.NewInt(int64(collateralRatio.CollateralRatio))))

		if usdValueOfCollateral.Cmp(usdValueOfDebt) >= 0 {
			return nil
		}
	}

	return fmt.Errorf("collateral is not enough")

}

func (k Keeper) calculateBorrowedAmountWithInterestAndFees(ctx sdk.Context, managerName string, lendingAccount types.LendingAccount) (borrowedAmountWithInterest *big.Int, borrowedAmountWithInterestAndFees *big.Int, err error) {

	lendingManager, found := k.GetLendingManager(ctx, managerName)
	if !found {
		return nil, nil, fmt.Errorf("lending manager with name %s not found", managerName)
	}

	asset, found := k.assetsKeeper.GetAsset(ctx, lendingManager.AssetId)
	if !found {
		return nil, nil, fmt.Errorf("asset with id %d not found", lendingManager.AssetId)
	}

	currentBorrowAPY, found := k.lendingPoolKeeper.GetCurrentBorrowAPY(ctx, asset.Denom)
	if !found {
		return nil, nil, fmt.Errorf("current borrow APY for asset %s not found", asset.Denom)
	}

	bigBorrowedAmount, err := lendingPoolTypes.ConvertStringToBigInt(lendingAccount.BorrowedAmount)
	if err != nil {
		return nil, nil, err
	}

	bigInitialBorrowIndex, err := lendingPoolTypes.ConvertStringToBigInt(lendingAccount.InitialBorrowIndex)
	if err != nil {
		return nil, nil, err
	}

	borrowedAmountWithInterest = new(big.Int).Div(new(big.Int).Mul(bigBorrowedAmount, currentBorrowAPY), bigInitialBorrowIndex)

	feeInterest := big.NewInt(int64(lendingManager.InterestFee))
	borrowedAmountWithInterestAndFees = new(big.Int).Sub(borrowedAmountWithInterest, bigBorrowedAmount)
	borrowedAmountWithInterestAndFees = new(big.Int).Div(new(big.Int).Mul(borrowedAmountWithInterestAndFees, feeInterest), types.PERCENTAGE_FACTOR)
	borrowedAmountWithInterestAndFees = new(big.Int).Add(borrowedAmountWithInterest, borrowedAmountWithInterest)

	return borrowedAmountWithInterest, borrowedAmountWithInterestAndFees, nil
}
