package keeper

import (
	"sort"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	indexerevents "github.com/StreamFinance-Protocol/stream-chain/protocol/indexer/events"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/indexer/indexer_manager"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/names/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateName(
	ctx sdk.Context,
	nameId uint32,
	nameName string,
) (types.Name, error) {
	if prevName, exists := k.GetName(ctx, nameId); exists {
		return types.Name{}, errorsmod.Wrapf(
			types.ErrNameIdAlreadyExists,
			"previous name = %v",
			prevName,
		)
	}

	if nameId == types.NameJota.Id {
		// Ensure nameId zero is always Jota. This is a protocol-wide invariant.
		if nameName != types.NameJota.Name {
			return types.Name{}, types.ErrJotaMustBeNameZero
		}
	}

	// Ensure Jota is not created with a non-zero nameId. This is a protocol-wide invariant.
	if nameId != types.NameJota.Id && nameName == types.NameJota.Name {
		return types.Name{}, types.ErrJotaMustBeNameZero
	}

	// Ensure the name is unique versus existing names.
	allNames := k.GetAllNames(ctx)
	for _, name := range allNames {
		if name.Name == nameName {
			return types.Name{}, errorsmod.Wrap(types.ErrNameNameAlreadyExists, nameName)
		}
	}

	// Create the name
	name := types.Name{
		Id:   nameId,
		Name: nameName,
	}

	// Store the new name
	k.setName(ctx, name)

	k.GetIndexerEventManager().AddTxnEvent(
		ctx,
		indexerevents.SubtypeName,
		indexerevents.NameEventVersion,
		indexer_manager.GetBytes(
			indexerevents.NewNameCreateEvent(
				nameId,
				nameName,
			),
		),
	)

	return name, nil
}

// func (k Keeper) ModifyAsset(
// 	ctx sdk.Context,
// 	id uint32,
// 	hasMarket bool,
// 	marketId uint32,
// ) (types.Asset, error) {
// 	// Get asset
// 	asset, exists := k.GetAsset(ctx, id)
// 	if !exists {
// 		return asset, errorsmod.Wrap(types.ErrAssetDoesNotExist, lib.UintToString(id))
// 	}

// 	// Validate market
// 	if _, err := k.pricesKeeper.GetMarketPrice(ctx, marketId); err != nil {
// 		return asset, err
// 	}

// 	// Modify asset
// 	asset.HasMarket = hasMarket
// 	asset.MarketId = marketId

// 	// Store the modified asset
// 	k.setAsset(ctx, asset)

// 	return asset, nil
// }

func (k Keeper) setName(
	ctx sdk.Context,
	name types.Name,
) {
	b := k.cdc.MustMarshal(&name)
	nameStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NameKeyPrefix))
	nameStore.Set(lib.Uint32ToKey(name.Id), b)
}

func (k Keeper) GetName(
	ctx sdk.Context,
	id uint32,
) (val types.Name, exists bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NameKeyPrefix))

	b := store.Get(lib.Uint32ToKey(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllNames(
	ctx sdk.Context,
) (list []types.Name) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NameKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Name
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Id < list[j].Id
	})

	return list
}

// GetNetCollateral returns the net collateral that a given position (quantums)
// for a given assetId contributes to an account.
// func (k Keeper) GetNetCollateral(
// 	ctx sdk.Context,
// 	id uint32,
// 	bigQuantums *big.Int,
// ) (
// 	bigNetCollateralQuoteQuantums *big.Int,
// 	err error,
// ) {
// 	if id == types.AssetTDai.Id {
// 		return new(big.Int).Set(bigQuantums), nil
// 	}

// 	// Get asset
// 	_, exists := k.GetAsset(ctx, id)
// 	if !exists {
// 		return big.NewInt(0), errorsmod.Wrap(types.ErrAssetDoesNotExist, lib.UintToString(id))
// 	}

// 	// Balance is zero.
// 	if bigQuantums.BitLen() == 0 {
// 		return big.NewInt(0), nil
// 	}

// 	// Balance is positive.
// 	// TODO(DEC-581): add multi-collateral support.
// 	if bigQuantums.Sign() == 1 {
// 		return big.NewInt(0), types.ErrNotImplementedMulticollateral
// 	}

// 	// Balance is negative.
// 	// TODO(DEC-582): add margin-trading support.
// 	return big.NewInt(0), types.ErrNotImplementedMargin
// }

// GetMarginRequirements returns the initial and maintenance margin-
// requirements for a given position size for a given assetId.
// func (k Keeper) GetMarginRequirements(
// 	ctx sdk.Context,
// 	id uint32,
// 	bigQuantums *big.Int,
// ) (
// 	bigInitialMarginQuoteQuantums *big.Int,
// 	bigMaintenanceMarginQuoteQuantums *big.Int,
// 	err error,
// ) {
// 	// QuoteBalance does not contribute to any margin requirements.
// 	if id == types.AssetTDai.Id {
// 		return big.NewInt(0), big.NewInt(0), nil
// 	}

// 	// Get asset
// 	_, exists := k.GetAsset(ctx, id)
// 	if !exists {
// 		return big.NewInt(0), big.NewInt(0), errorsmod.Wrap(
// 			types.ErrAssetDoesNotExist, lib.UintToString(id))
// 	}

// 	// Balance is zero or positive.
// 	if bigQuantums.Sign() >= 0 {
// 		return big.NewInt(0), big.NewInt(0), nil
// 	}

// 	// Balance is negative.
// 	// TODO(DEC-582): margin-trading
// 	return big.NewInt(0), big.NewInt(0), types.ErrNotImplementedMargin
// }

// ConvertAssetToCoin converts the given `assetId` and `quantums` used in `x/asset`,
// to an `sdk.Coin` in correspoding `denom` and `amount` used in `x/bank`.
// Also outputs `convertedQuantums` which has the equal value as converted `sdk.Coin`.
// The conversion is done with the formula:
//
//	denom_amount = quantums * 10^(atomic_resolution - denom_exponent)
//
// If the resulting `denom_amount` is not an integer, it is rounded down,
// and `convertedQuantums` of the equal value is returned. The upstream
// transfer function should adjust asset balance with `convertedQuantums`
// to ensure that that no fund is ever lost in the conversion due to rounding error.
//
// Example:
// Assume `denom_exponent` = -7, `atomic_resolution` = -8,
// ConvertAssetToCoin(`101 quantums`) should output:
// - `convertedQuantums` = 100 quantums
// -  converted coin amount = 10 coin
// func (k Keeper) ConvertAssetToCoin(
// 	ctx sdk.Context,
// 	assetId uint32,
// 	quantums *big.Int,
// ) (
// 	convertedQuantums *big.Int,
// 	coin sdk.Coin,
// 	err error,
// ) {
// 	asset, exists := k.GetAsset(ctx, assetId)
// 	if !exists {
// 		return nil, sdk.Coin{}, errorsmod.Wrap(
// 			types.ErrAssetDoesNotExist, lib.UintToString(assetId))
// 	}

// 	if lib.AbsInt32(asset.AtomicResolution) > types.MaxAssetUnitExponentAbs {
// 		return nil, sdk.Coin{}, errorsmod.Wrapf(
// 			types.ErrInvalidAssetAtomicResolution,
// 			"asset: %+v",
// 			asset,
// 		)
// 	}

// 	if lib.AbsInt32(asset.DenomExponent) > types.MaxAssetUnitExponentAbs {
// 		return nil, sdk.Coin{}, errorsmod.Wrapf(
// 			types.ErrInvalidDenomExponent,
// 			"asset: %+v",
// 			asset,
// 		)
// 	}

// 	bigRatDenomAmount := lib.BigMulPow10(
// 		quantums,
// 		asset.AtomicResolution-asset.DenomExponent,
// 	)

// 	// round down to get denom amount that was converted.
// 	bigConvertedDenomAmount := lib.BigRatRound(bigRatDenomAmount, false)

// 	bigRatConvertedQuantums := lib.BigMulPow10(
// 		bigConvertedDenomAmount,
// 		asset.DenomExponent-asset.AtomicResolution,
// 	)

// 	bigConvertedQuantums := bigRatConvertedQuantums.Num()

// 	return bigConvertedQuantums, sdk.NewCoin(asset.Denom, sdkmath.NewIntFromBigInt(bigConvertedDenomAmount)), nil
// }

// ConvertCoinToAsset converts the given `sdk.Coin` used in `x/bank` to
// the corresponding `quantums` used in `x/asset` for the given `assetId`.
// The conversion is done with the inverse formula of ConvertAssetToCoin:
//
//	quantums = coin_amount * 10^(denom_exponent - atomic_resolution)
//
// If the resulting `quantums` is not an integer, it is rounded down.
// This ensures consistency with ConvertAssetToCoin and prevents
// creation of assets from rounding up.
// func (k Keeper) ConvertCoinToAsset(
// 	ctx sdk.Context,
// 	assetId uint32,
// 	coin sdk.Coin,
// ) (
// 	quantums *big.Int,
// 	convertedDenom *big.Int,
// 	err error,
// ) {
// 	asset, exists := k.GetAsset(ctx, assetId)
// 	if !exists {
// 		return nil, nil, errorsmod.Wrap(
// 			types.ErrAssetDoesNotExist, lib.UintToString(assetId))
// 	}

// 	if lib.AbsInt32(asset.AtomicResolution) > types.MaxAssetUnitExponentAbs {
// 		return nil, nil, errorsmod.Wrapf(
// 			types.ErrInvalidAssetAtomicResolution,
// 			"asset: %+v",
// 			asset,
// 		)
// 	}

// 	if lib.AbsInt32(asset.DenomExponent) > types.MaxAssetUnitExponentAbs {
// 		return nil, nil, errorsmod.Wrapf(
// 			types.ErrInvalidDenomExponent,
// 			"asset: %+v",
// 			asset,
// 		)
// 	}

// 	bigRatQuantums := lib.BigMulPow10(
// 		coin.Amount.BigInt(),
// 		asset.DenomExponent-asset.AtomicResolution,
// 	)

// 	quantums = lib.BigRatRound(bigRatQuantums, false)

// 	// If the result is zero, return a true zero for backwards compatibility
// 	if quantums.Sign() == 0 {
// 		return big.NewInt(0), big.NewInt(0), nil
// 	}

// 	bigRatConvertedDenomAmount := lib.BigMulPow10(
// 		quantums,
// 		asset.AtomicResolution-asset.DenomExponent,
// 	)

// 	convertedDenom = bigRatConvertedDenomAmount.Num()

// 	return quantums, convertedDenom, nil
// }

// ConvertAssetToFullCoin converts the given `assetId` and `quantums`
// to the amount of full coins given by the atomic resolution.
// fullCointAmount = quantums * 10^(atomic_resolution)
//
// If the resulting full coin amount is not an integer, it is rounded
// down and `convertedQuantums` of the equal value is returned. If
// quantums amount is negative or 0, returns 0 as a result.
// func (k Keeper) ConvertAssetToFullCoin(
// 	ctx sdk.Context,
// 	assetId uint32,
// 	quantums *big.Int,
// ) (
// 	convertedQuantums *big.Int,
// 	fullCoinAmount *big.Int,
// 	err error,
// ) {
// 	asset, exists := k.GetAsset(ctx, assetId)
// 	if !exists {
// 		return nil, nil, errorsmod.Wrap(
// 			types.ErrAssetDoesNotExist, lib.UintToString(assetId))
// 	}

// 	if lib.AbsInt32(asset.AtomicResolution) > types.MaxAssetUnitExponentAbs {
// 		return nil, nil, errorsmod.Wrapf(
// 			types.ErrInvalidAssetAtomicResolution,
// 			"asset: %+v",
// 			asset,
// 		)
// 	}

// 	if quantums.Sign() <= 0 {
// 		return big.NewInt(0), big.NewInt(0), nil
// 	}

// 	fullCoinAmount = lib.QuoteQuantumsToFullCoinAmount(quantums, asset.AtomicResolution)

// 	convertedQuantums = lib.BigMulPow10(
// 		fullCoinAmount,
// 		asset.AtomicResolution,
// 	).Num()

// 	return convertedQuantums, fullCoinAmount, nil
// }

// IsPositionUpdatable returns whether position of an asset is updatable.
// func (k Keeper) IsPositionUpdatable(
// 	ctx sdk.Context,
// 	id uint32,
// ) (
// 	updatable bool,
// 	err error,
// ) {
// 	_, exists := k.GetAsset(ctx, id)
// 	if !exists {
// 		return false, errorsmod.Wrap(types.ErrAssetDoesNotExist, lib.UintToString(id))
// 	}
// 	// All existing assets are by default updatable.
// 	return true, nil
// }
