package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
	perptypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/perpetuals/types"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/subaccounts/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) checkMultiCollateralAssetConstraints(
	ctx sdk.Context,
	settledUpdates []SettledUpdate,
	perpetuals []perptypes.Perpetual,
) (
	success bool,
	successPerUpdate []types.UpdateResult,
	err error,
) {
	success = true
	successPerUpdate = make([]types.UpdateResult, len(settledUpdates))
	perpIdToMarketType := getPerpIdToParams(perpetuals)

	for i, u := range settledUpdates {
		result, err := k.isValidMultiCollateralUpdate(ctx, u, perpIdToMarketType)
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

func (k Keeper) isValidMultiCollateralUpdate(
	ctx sdk.Context,
	settledUpdate SettledUpdate,
	perpIdToParams map[uint32]perptypes.PerpetualParams,
) (types.UpdateResult, error) {

	if len(settledUpdate.AssetUpdates) == 0 {
		return types.Success, nil
	}

	collateralPoolId := uint32(0)

	if len(settledUpdate.SettledSubaccount.PerpetualPositions) > 0 {
		params, exists := perpIdToParams[settledUpdate.SettledSubaccount.PerpetualPositions[0].PerpetualId]
		if !exists {
			return types.UpdateCausedError, errorsmod.Wrap(
				perptypes.ErrPerpetualDoesNotExist, lib.UintToString(settledUpdate.SettledSubaccount.PerpetualPositions[0].PerpetualId),
			)
		}

		collateralPoolId = params.CollateralPoolId
	} else if len(settledUpdate.PerpetualUpdates) > 0 {
		params, exists := perpIdToParams[settledUpdate.PerpetualUpdates[0].PerpetualId]
		if !exists {
			return types.UpdateCausedError, errorsmod.Wrap(
				perptypes.ErrPerpetualDoesNotExist, lib.UintToString(settledUpdate.PerpetualUpdates[0].PerpetualId),
			)
		}

		collateralPoolId = params.CollateralPoolId
	}

	collateralPool, err := k.perpetualsKeeper.GetCollateralPool(ctx, collateralPoolId)
	if err != nil {
		return types.UpdateCausedError, err
	}
	supportedAssetIds := getValidAssetIdMap(collateralPool.MultiCollateralAssets.MultiCollateralAssets)
	for _, assetUpdate := range settledUpdate.AssetUpdates {
		_, ok := supportedAssetIds[assetUpdate.AssetId]
		if !ok {
			return types.ViolatesMultiCollateralConstraints, nil
		}
	}

	return types.Success, nil
}

func getPerpIdToParams(
	perpetuals []perptypes.Perpetual,
) map[uint32]perptypes.PerpetualParams {
	var perpIdToMarketType = make(map[uint32]perptypes.PerpetualParams)

	for _, perpetual := range perpetuals {
		perpIdToMarketType[perpetual.GetId()] = perpetual.Params
	}

	return perpIdToMarketType
}

func getValidAssetIdMap(
	assetIds []uint32,
) (assetIdsMap map[uint32]struct{}) {

	assetIdsMap = make(map[uint32]struct{})
	for _, asset := range assetIds {
		assetIdsMap[asset] = struct{}{}
	}

	return assetIdsMap
}
