package keeper

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
	perptypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/perpetuals/types"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/subaccounts/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) checkCollateralPoolConstraints(
	settledUpdates []SettledUpdate,
	perpetuals []perptypes.Perpetual,
) (
	success bool,
	successPerUpdate []types.UpdateResult,
	err error,
) {
	success = true
	successPerUpdate = make([]types.UpdateResult, len(settledUpdates))
	perpIdToCollateralPoolId := getPerpIdToCollateralPoolIdMap(perpetuals)

	for i, u := range settledUpdates {
		result, err := isValidCollateralPoolUpdates(u, perpIdToCollateralPoolId)
		if err != nil {
			return false, nil, err
		}
		if result != types.Success {
			success = false
		}

		successPerUpdate[i] = result
	}

	return success, successPerUpdate, nil
}

// Checks whether the perpetual updates to a settled subaccount violates constraints for the
// collateral pool. This function assumes the settled subaccount is valid and does not violate the
// the constraints.
// The constraint being checked is:
//   - there cannot be updates for multiple collateral pools
//   - a subaccount with a position in a collateral pool cannot have updates for perpetuals in other
//     collateral pools
//   - a subaccount with no positions cannot be updated to have positions in multiple collateral pools
func isValidCollateralPoolUpdates(
	settledUpdate SettledUpdate,
	perpIdToCollateralPoolId map[uint32]uint32,
) (types.UpdateResult, error) {

	if len(settledUpdate.PerpetualUpdates) == 0 {
		return types.Success, nil
	}

	collateralPoolIdBeingUpdated := perpIdToCollateralPoolId[settledUpdate.PerpetualUpdates[0].PerpetualId]

	for _, perpetualUpdate := range settledUpdate.PerpetualUpdates[1:] {
		collateralPoolId, exists := perpIdToCollateralPoolId[perpetualUpdate.PerpetualId]
		if !exists {
			return types.UpdateCausedError, errorsmod.Wrap(
				perptypes.ErrPerpetualDoesNotExist, lib.UintToString(perpetualUpdate.PerpetualId),
			)
		}
		if collateralPoolId != collateralPoolIdBeingUpdated {
			return types.ViolatesCollateralPoolConstraints, nil
		}
	}

	if len(settledUpdate.SettledSubaccount.PerpetualPositions) == 0 {
		return types.Success, nil
	}

	collateralPoolId := perpIdToCollateralPoolId[settledUpdate.SettledSubaccount.PerpetualPositions[0].PerpetualId]

	if collateralPoolId != collateralPoolIdBeingUpdated {
		return types.ViolatesCollateralPoolConstraints, nil
	}

	return types.Success, nil
}

// GetIsolatedPerpetualStateTransition computes whether an isolated perpetual position will be
// opened or closed for a subaccount.
// This function assumes that the subaccount is valid under isolated perpetual constraints.
// The input `settledUpdate` must have an updated subaccount (`settledUpdate.SettledSubaccount`),
// so all the updates must have been applied already to the subaccount.
func GetIsolatedPerpetualStateTransition(
	settledUpdateWithUpdatedSubaccount SettledUpdate,
	perpetuals []perptypes.Perpetual,
) (*types.IsolatedPerpetualPositionStateTransition, error) {
	perpIdToMarketType := getPerpIdToMarketTypeMap(perpetuals)
	// This subaccount needs to have had the updates in the `settledUpdate` already applied to it.
	updatedSubaccount := settledUpdateWithUpdatedSubaccount.SettledSubaccount
	// If there are no perpetual updates, then no perpetual position could have been opened or closed
	// on the subaccount.
	if len(settledUpdateWithUpdatedSubaccount.PerpetualUpdates) == 0 {
		return nil, nil
	}

	// If there are more than 1 valid perpetual update, or more than 1 valid perpetual position on the
	// subaccount, it is not isolated to an isolated perpetual, and so no isolated perpetual position
	// could have been opened or closed.
	if len(settledUpdateWithUpdatedSubaccount.PerpetualUpdates) > 1 ||
		len(updatedSubaccount.PerpetualPositions) > 1 {
		return nil, nil
	}

	// Now, from the above checks, we know there is only a single perpetual update and 0 or 1 perpetual
	// positions.
	perpetualUpdate := settledUpdateWithUpdatedSubaccount.PerpetualUpdates[0]
	marketType, exists := perpIdToMarketType[perpetualUpdate.PerpetualId]
	if !exists {
		return nil, errorsmod.Wrap(
			perptypes.ErrPerpetualDoesNotExist, lib.UintToString(perpetualUpdate.PerpetualId),
		)
	}
	// If the perpetual update is not for an isolated perpetual, no isolated perpetual position is
	// being opened or closed.
	if marketType != perptypes.PerpetualMarketType_PERPETUAL_MARKET_TYPE_ISOLATED {
		return nil, nil
	}

	// If the updated subaccount does not have any perpetual positions, then an isolated perpetual
	// position must have been closed due to the perpetual update.
	if len(updatedSubaccount.PerpetualPositions) == 0 {

		assetIds := make([]uint32, 0, len(updatedSubaccount.AssetPositions))
		assetSizes := make([]*big.Int, 0, len(updatedSubaccount.AssetPositions))

		for _, assetPosition := range updatedSubaccount.AssetPositions {
			assetIds = append(assetIds, assetPosition.AssetId)
			assetSizes = append(assetSizes, assetPosition.GetBigQuantums())
		}

		return &types.IsolatedPerpetualPositionStateTransition{
			SubaccountId: updatedSubaccount.Id,
			PerpetualId:  perpetualUpdate.PerpetualId,
			AssetIds:     assetIds,
			BigQuantums:  assetSizes,
			Transition:   types.Closed,
		}, nil
	}

	// After the above checks, the subaccount must have only a single perpetual position, which is for
	// the same isolated perpetual as the perpetual update.
	perpetualPosition := updatedSubaccount.PerpetualPositions[0]
	// If the size of the update and the position are the same, the perpetual update must have opened
	// the position.
	if perpetualUpdate.GetBigQuantums().Cmp(perpetualPosition.GetBigQuantums()) == 0 {
		// Collateral equal to the quote currency asset position before the update was applied needs to be transferred.
		// Subtract the delta from the updated subaccount's quote currency asset position size to get the size
		// of the quote currency asset position.
		assetIds := make([]uint32, 0, len(updatedSubaccount.AssetPositions))
		assetSizes := make([]*big.Int, 0, len(updatedSubaccount.AssetPositions))

		assetUpdateMap := make(map[uint32]*big.Int)
		for _, assetUpdate := range settledUpdateWithUpdatedSubaccount.AssetUpdates {
			assetUpdateMap[assetUpdate.AssetId] = assetUpdate.GetBigQuantums()
		}

		for _, assetPosition := range updatedSubaccount.AssetPositions {
			assetIds = append(assetIds, assetPosition.AssetId)

			updateQuantums, exists := assetUpdateMap[assetPosition.AssetId]
			if !exists {
				assetSizes = append(assetSizes, assetPosition.GetBigQuantums())
			} else {
				assetSizes = append(assetSizes, new(big.Int).Sub(assetPosition.GetBigQuantums(), updateQuantums))
			}
		}

		return &types.IsolatedPerpetualPositionStateTransition{
			SubaccountId: updatedSubaccount.Id,
			PerpetualId:  perpetualUpdate.PerpetualId,
			AssetIds:     assetIds,
			BigQuantums:  assetSizes,
			Transition:   types.Opened,
		}, nil
	}

	// The isolated perpetual position changed size but was not opened or closed.
	return nil, nil
}

// transferCollateralForIsolatedPerpetual transfers collateral between an isolated collateral pool
// and the cross-perpetual collateral pool based on whether an isolated perpetual position was
// opened or closed in a subaccount.
// Note: This uses the `x/bank` keeper and modifies `x/bank` state.
func (k *Keeper) transferCollateralForIsolatedPerpetual(
	ctx sdk.Context,
	stateTransition *types.IsolatedPerpetualPositionStateTransition,
) error {
	// No collateral to transfer if no state transition.
	if stateTransition == nil {
		return nil
	}

	if len(stateTransition.AssetIds) != len(stateTransition.BigQuantums) {
		return errorsmod.Wrap(
			types.ErrFailedToUpdateSubaccounts,
			"Asset IDs and big quantums arrays must be the same length",
		)
	}

	isolatedCollateralPoolAddr, err := k.GetCollateralPoolFromPerpetualId(ctx, stateTransition.PerpetualId)
	if err != nil {
		return err
	}
	var toModuleAddr sdk.AccAddress
	var fromModuleAddr sdk.AccAddress

	// If an isolated perpetual position was opened in the subaccount, then move collateral from the
	// cross-perpetual collateral pool to the isolated perpetual collateral pool.
	if stateTransition.Transition == types.Opened {
		toModuleAddr = isolatedCollateralPoolAddr
		fromModuleAddr = types.ModuleAddress
		// If the isolated perpetual position was closed, then move collateral from the isolated
		// perpetual collateral pool to the cross-perpetual collateral pool.
	} else if stateTransition.Transition == types.Closed {
		toModuleAddr = types.ModuleAddress
		fromModuleAddr = isolatedCollateralPoolAddr
	} else {
		// Should never hit this.
		return errorsmod.Wrapf(
			types.ErrFailedToUpdateSubaccounts,
			"Invalid state transition %v for isolated perpetual with id %d in subaccount with id %v",
			stateTransition,
			stateTransition.PerpetualId,
			stateTransition.SubaccountId,
		)
	}

	for i, _ := range stateTransition.AssetIds {

		// If there are zero quantums to transfer, don't transfer collateral.
		if stateTransition.BigQuantums[i].Sign() == 0 {
			continue
		}

		// Invalid to transfer negative quantums. This should already be caught by collateralization
		// checks as well.
		if stateTransition.BigQuantums[i].Sign() == -1 {
			return errorsmod.Wrapf(
				types.ErrFailedToUpdateSubaccounts,
				"Subaccount with id %v %s perpteual position with perpetual id %d with negative collateral asset id %d with size %s to transfer",
				stateTransition.SubaccountId,
				stateTransition.Transition.String(),
				stateTransition.PerpetualId,
				stateTransition.AssetIds[i],
				stateTransition.BigQuantums[i].String(),
			)
		}

		// Transfer collateral between collateral pools.
		_, coinToTransfer, err := k.assetsKeeper.ConvertAssetToCoin(
			ctx,
			stateTransition.AssetIds[i],
			stateTransition.BigQuantums[i],
		)
		if err != nil {
			return err
		}

		if err = k.bankKeeper.SendCoins(
			ctx,
			fromModuleAddr,
			toModuleAddr,
			[]sdk.Coin{coinToTransfer},
		); err != nil {
			return err
		}

	}
	return nil
}

// computeAndExecuteCollateralTransfer computes collateral transfers resulting from updates to
// a subaccount and executes the collateral transfer using `x/bank`.`
// The input `settledUpdate` must have an updated subaccount (`settledUpdate.SettledSubaccount`),
// so all the updates must have been applied already to the subaccount.
// Note: This uses the `x/bank` keeper and modifies `x/bank` state.
func (k *Keeper) computeAndExecuteCollateralTransfer(
	ctx sdk.Context,
	settledUpdateWithUpdatedSubaccount SettledUpdate,
	perpetuals []perptypes.Perpetual,
) error {
	// The subaccount in `settledUpdateWithUpdatedSubaccount` already has the perpetual updates
	// and asset updates applied to it.
	stateTransition, err := GetIsolatedPerpetualStateTransition(
		settledUpdateWithUpdatedSubaccount,
		perpetuals,
	)
	if err != nil {
		return err
	}
	if err := k.transferCollateralForIsolatedPerpetual(
		ctx,
		stateTransition,
	); err != nil {
		return err
	}

	return nil
}

func getPerpIdToCollateralPoolIdMap(
	perpetuals []perptypes.Perpetual,
) map[uint32]uint32 {
	var perpIdToCollateralPoolId = make(map[uint32]uint32)

	for _, perpetual := range perpetuals {
		perpIdToCollateralPoolId[perpetual.GetId()] = perpetual.Params.CollateralPoolId
	}

	return perpIdToCollateralPoolId
}
