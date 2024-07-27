package keeper

import (
	"fmt"
	"math/big"

	types "github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingaccount/types"
	lendingpooltypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// collateralAssetIds is expected to be sorted by asset ID
func (k Keeper) OpenCreditAccount(ctx sdk.Context, managerName string, lendingAccountId uint32, address string, borrowedAmount *big.Int, collateralAssetIds []uint32, amounts []*big.Int) error {

	lendingInterface, found := k.GetLendingInterface(ctx, managerName)
	if !found {
		return fmt.Errorf("lending interface with manager name %s not found", managerName)
	}

	err := checkBorrowAccountLimit(lendingInterface, borrowedAmount)
	if err != nil {
		return err
	}

	err = k.checkAndUpdateBorrowBlockLimit(ctx, lendingInterface, borrowedAmount)
	if err != nil {
		return err
	}

	err = k.checkAndUpdateTotalDebt(ctx, managerName, borrowedAmount, true, lendingInterface.TotalDebtLimit)
	if err != nil {
		return err
	}

	lendingManager, found := k.GetLendingManager(ctx, managerName)
	if !found {
		return fmt.Errorf("lending manager with manager name %s not found", managerName)
	}

	err = checkCollateralAssetsAreAllowed(lendingManager.EnabledCollateralAssets, collateralAssetIds)
	if err != nil {
		return err
	}

	err = k.checkCollateralOfPosition(ctx, managerName, collateralAssetIds, amounts, borrowedAmount)
	if err != nil {
		return err
	}

	err = k.openCreditAccountOnManager(ctx, managerName, lendingAccountId, address, borrowedAmount, collateralAssetIds, amounts)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) checkAndUpdateBorrowBlockLimit(ctx sdk.Context, lendingInterface types.LendingInterface, borrowedAmount *big.Int) error {

	bigMaxBorrowPerBlock, err := lendingpooltypes.ConvertStringToBigInt(lendingInterface.MaxBorrowPerBlock)
	if err != nil {
		return err
	}

	totalBorrowedLastBlock, err := k.GetTotalBorrowedLastBlock(ctx, lendingInterface.ManagerName)
	if err != nil {
		return err
	}

	lastBorrowedInBlock, err := k.GetLastBorrowedInBlock(ctx, lendingInterface.ManagerName)
	if err != nil {
		return err
	}

	currentBlock := ctx.BlockHeight()
	newLimit := new(big.Int)
	if lastBorrowedInBlock.Cmp(big.NewInt(currentBlock)) == 0 {
		newLimit.Add(borrowedAmount, totalBorrowedLastBlock)
	} else {
		newLimit.Set(borrowedAmount)
		k.SetLastBorrowedInBlock(ctx, lendingInterface.ManagerName, big.NewInt(currentBlock))
	}

	if newLimit.Cmp(bigMaxBorrowPerBlock) > 0 {
		return fmt.Errorf("borrowed amount is greater than the block limit")
	}

	k.SetTotalBorrowedLastBlock(ctx, lendingInterface.ManagerName, newLimit)

	return nil

}

func checkBorrowAccountLimit(lendingInterface types.LendingInterface, borrowedAmount *big.Int) error {

	bigMaxBorrowPerAccount, err := lendingpooltypes.ConvertStringToBigInt(lendingInterface.MaxBorrowPerAccount)
	if err != nil {
		return err
	}

	bigMinBorrowPerAccount, err := lendingpooltypes.ConvertStringToBigInt(lendingInterface.MinBorrowPerAccount)
	if err != nil {
		return err
	}

	if borrowedAmount.Cmp(bigMinBorrowPerAccount) < 0 {
		return fmt.Errorf("borrowed amount is less than the minimum borrow per account")
	}

	if borrowedAmount.Cmp(bigMaxBorrowPerAccount) > 0 {
		return fmt.Errorf("borrowed amount is greater than the maximum borrow per account")
	}

	return nil
}

func (k Keeper) checkAndUpdateTotalDebt(ctx sdk.Context, managerName string, delta *big.Int, isIncrease bool, debtLimit string) error {
	if delta.Cmp(big.NewInt(0)) > 0 {
		totalDebt, err := k.GetTotalDebt(ctx, managerName)
		if err != nil {
			return err
		}

		totalDebtLimit, err := lendingpooltypes.ConvertStringToBigInt(debtLimit)
		if err != nil {
			return err
		}

		if isIncrease {
			totalDebt.Add(totalDebt, delta)
			if totalDebt.Cmp(totalDebtLimit) > 0 {
				return fmt.Errorf("borrow amount out of limits")
			}
		} else {
			if totalDebt.Cmp(delta) < 0 {
				return fmt.Errorf("borrow amount out of limits")
			}
			totalDebt.Sub(totalDebt, delta)
		}

		k.SetTotalDebt(ctx, managerName, totalDebt)
	}
	return nil
}

// Both arrays are sorted by assetId
func checkCollateralAssetsAreAllowed(enabledCollateralAssets map[uint32]*types.EnabledCollateralAssets, collateralAssetIds []uint32) error {

	for _, collateralAssetId := range collateralAssetIds {
		_, found := enabledCollateralAssets[collateralAssetId]
		if !found {
			return fmt.Errorf("collateral asset with id %d is not allowed", collateralAssetId)
		}
	}

	return nil
}
