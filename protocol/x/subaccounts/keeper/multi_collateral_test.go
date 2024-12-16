package keeper_test

import (
	"testing"

	"math/big"

	errorsmod "cosmossdk.io/errors"
	constants "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/constants"
	testutil "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/keeper"
	assetskeeper "github.com/StreamFinance-Protocol/stream-chain/protocol/x/assets/keeper"
	asstypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/assets/types"
	perpetualskeeper "github.com/StreamFinance-Protocol/stream-chain/protocol/x/perpetuals/keeper"
	perptypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/perpetuals/types"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/subaccounts/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/subaccounts/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

type MultiCollatUpdateTest struct {
	settledUpdate  keeper.SettledUpdate
	perpIdToParams map[uint32]perptypes.PerpetualParams
	perpetuals     []perptypes.Perpetual
	assets         []asstypes.Asset
	expectedResult types.UpdateResult
	expectedErr    error
}

func TestIsValidMultiCollateralUpdate(t *testing.T) {
	tests := map[string]MultiCollatUpdateTest{
		"Success: no asset updates": {
			settledUpdate: keeper.SettledUpdate{
				SettledSubaccount: types.Subaccount{},
				AssetUpdates:      []types.AssetUpdate{},
			},
			perpetuals: []perptypes.Perpetual{},
			assets: []asstypes.Asset{
				*constants.TDai,
			},
			perpIdToParams: map[uint32]perptypes.PerpetualParams{},
			expectedResult: types.Success,
		},
		"Success: asset updates, no perps or perp updates": {
			settledUpdate: keeper.SettledUpdate{
				SettledSubaccount: constants.Alice_Num0_1USD,
				AssetUpdates: []types.AssetUpdate{
					{
						AssetId:          0,
						BigQuantumsDelta: big.NewInt(10_000_000),
					},
				},
			},
			perpetuals: []perptypes.Perpetual{
				constants.BtcUsd_0DefaultFunding_10AtomicResolution,
			},
			assets: []asstypes.Asset{
				*constants.TDai,
			},
			perpIdToParams: map[uint32]perptypes.PerpetualParams{
				0: constants.BtcUsd_0DefaultFunding_10AtomicResolution_Params,
			},
			expectedResult: types.Success,
		},
		"Success: New Position with valid asset updates": {
			settledUpdate: keeper.SettledUpdate{
				SettledSubaccount: constants.Alice_Num0_1USD,
				AssetUpdates: []types.AssetUpdate{
					{
						AssetId:          0,
						BigQuantumsDelta: big.NewInt(10_000_000),
					},
				},
				PerpetualUpdates: []types.PerpetualUpdate{
					{
						PerpetualId:      0,
						BigQuantumsDelta: big.NewInt(10_000_000_000),
					},
				},
			},
			perpetuals: []perptypes.Perpetual{
				constants.BtcUsd_0DefaultFunding_10AtomicResolution,
			},
			assets: []asstypes.Asset{
				*constants.TDai,
			},
			perpIdToParams: map[uint32]perptypes.PerpetualParams{
				0: constants.BtcUsd_0DefaultFunding_10AtomicResolution_Params,
			},
			expectedResult: types.Success,
		},
		"Success: Closing Position with valid asset updates": {
			settledUpdate: keeper.SettledUpdate{
				SettledSubaccount: constants.Alice_Num1_1BTC_Short_100_000USD,
				AssetUpdates: []types.AssetUpdate{
					{
						AssetId:          0,
						BigQuantumsDelta: big.NewInt(100_000_000),
					},
				},
				PerpetualUpdates: []types.PerpetualUpdate{
					{
						PerpetualId:      0,
						BigQuantumsDelta: big.NewInt(100_000_000),
					},
				},
			},
			perpetuals: []perptypes.Perpetual{
				constants.BtcUsd_0DefaultFunding_10AtomicResolution,
			},
			assets: []asstypes.Asset{
				*constants.TDai,
			},
			perpIdToParams: map[uint32]perptypes.PerpetualParams{
				0: constants.BtcUsd_0DefaultFunding_10AtomicResolution_Params,
			},
			expectedResult: types.Success,
		},
		"Success: Updating Position with valid asset updates": {
			settledUpdate: keeper.SettledUpdate{
				SettledSubaccount: constants.Alice_Num1_1BTC_Short_100_000USD,
				AssetUpdates: []types.AssetUpdate{
					{
						AssetId:          0,
						BigQuantumsDelta: big.NewInt(100_000_000),
					},
				},
				PerpetualUpdates: []types.PerpetualUpdate{
					{
						PerpetualId:      0,
						BigQuantumsDelta: big.NewInt(-100_000_000),
					},
				},
			},
			perpetuals: []perptypes.Perpetual{
				constants.BtcUsd_0DefaultFunding_10AtomicResolution,
			},
			assets: []asstypes.Asset{
				*constants.TDai,
			},
			perpIdToParams: map[uint32]perptypes.PerpetualParams{
				0: constants.BtcUsd_0DefaultFunding_10AtomicResolution_Params,
			},
			expectedResult: types.Success,
		},
		"Failure: perp update for perp that doesn't exist": {
			settledUpdate: keeper.SettledUpdate{
				SettledSubaccount: constants.Alice_Num0_1USD,
				AssetUpdates: []types.AssetUpdate{
					{
						AssetId:          0,
						BigQuantumsDelta: big.NewInt(10_000_000),
					},
				},
				PerpetualUpdates: []types.PerpetualUpdate{
					{
						PerpetualId:      1,
						BigQuantumsDelta: big.NewInt(10_000_000),
					},
				},
			},

			perpetuals: []perptypes.Perpetual{
				constants.BtcUsd_0DefaultFunding_10AtomicResolution,
			},
			assets: []asstypes.Asset{
				*constants.TDai,
			},
			perpIdToParams: map[uint32]perptypes.PerpetualParams{
				999: constants.BtcUsd_0DefaultFunding_10AtomicResolution_Params,
			},
			expectedResult: types.UpdateCausedError,
			expectedErr:    errorsmod.Wrap(perptypes.ErrPerpetualDoesNotExist, "999"),
		},
		"Failure: asset doesn't exist in collat pool - multi asset updates": {
			settledUpdate: keeper.SettledUpdate{
				SettledSubaccount: constants.Alice_Num1_1BTC_Short_100_000USD,
				AssetUpdates: []types.AssetUpdate{
					{
						AssetId:          0,
						BigQuantumsDelta: big.NewInt(100_000_000),
					},
					{
						AssetId:          1,
						BigQuantumsDelta: big.NewInt(100_000_000),
					},
				},
				PerpetualUpdates: []types.PerpetualUpdate{
					{
						PerpetualId:      0,
						BigQuantumsDelta: big.NewInt(100_000_000),
					},
				},
			},
			perpetuals: []perptypes.Perpetual{
				constants.BtcUsd_0DefaultFunding_10AtomicResolution,
			},
			assets: []asstypes.Asset{
				*constants.TDai,
				*constants.BtcUsd,
			},
			perpIdToParams: map[uint32]perptypes.PerpetualParams{
				0: constants.BtcUsd_0DefaultFunding_10AtomicResolution_Params,
			},
			expectedResult: types.ViolatesMultiCollateralConstraints,
		},
		"Failure: asset doesn't exist in collat pool - single asset updates": {
			settledUpdate: keeper.SettledUpdate{
				SettledSubaccount: constants.Alice_Num1_1BTC_Short_100_000USD,
				AssetUpdates: []types.AssetUpdate{
					{
						AssetId:          1,
						BigQuantumsDelta: big.NewInt(100_000_000),
					},
				},
				PerpetualUpdates: []types.PerpetualUpdate{
					{
						PerpetualId:      0,
						BigQuantumsDelta: big.NewInt(100_000_000),
					},
				},
			},
			perpetuals: []perptypes.Perpetual{
				constants.BtcUsd_0DefaultFunding_10AtomicResolution,
			},
			assets: []asstypes.Asset{
				*constants.TDai,
				*constants.BtcUsd,
			},
			perpIdToParams: map[uint32]perptypes.PerpetualParams{
				0: constants.BtcUsd_0DefaultFunding_10AtomicResolution_Params,
			},
			expectedResult: types.ViolatesMultiCollateralConstraints,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ctx, keeper, _, perpKeeper, _, _, assetsKeeper, _, _, _ := testutil.SubaccountsKeepers(
				t,
				true,
			)

			createAssets(ctx, assetsKeeper, tc.assets)

			createPerpetuals(ctx, perpKeeper, tc.perpetuals)

			result, err := keeper.IsValidMultiCollateralUpdate(
				ctx,
				tc.settledUpdate,
				tc.perpIdToParams,
			)

			if tc.expectedErr != nil {
				require.ErrorIs(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.expectedResult, result)

		})
	}
}

func createPerpetuals(
	ctx sdk.Context,
	perpKeeper *perpetualskeeper.Keeper,
	perpetuals []perptypes.Perpetual,
) {

	perpKeeper.UpsertCollateralPool(
		ctx,
		0,
		1000,
		&perptypes.MultiCollateralAssetsArray{
			MultiCollateralAssets: []uint32{0},
		},
		0,
	)

	for _, perp := range perpetuals {
		perpKeeper.CreatePerpetual(
			ctx,
			perp.Params.Id,
			perp.Params.Ticker,
			perp.Params.MarketId,
			perp.Params.AtomicResolution,
			perp.Params.DefaultFundingPpm,
			perp.Params.LiquidityTier,
			perp.Params.DangerIndexPpm,
			perp.Params.CollateralPoolId,
			perp.YieldIndex,
		)
	}
}

func createAssets(
	ctx sdk.Context,
	assetsKeeper *assetskeeper.Keeper,
	assets []asstypes.Asset,
) {
	for _, asset := range assets {
		assetsKeeper.CreateAsset(
			ctx,
			asset.Id,
			asset.Symbol,
			asset.Denom,
			asset.DenomExponent,
			asset.HasMarket,
			asset.MarketId,
			asset.AtomicResolution,
			asset.AssetYieldIndex,
			asset.MaxSlippagePpm,
		)
	}
}
