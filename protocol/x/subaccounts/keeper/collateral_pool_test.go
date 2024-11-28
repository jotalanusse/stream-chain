package keeper_test

import (
	"math/big"
	"testing"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/dtypes"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/constants"
	assettypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/assets/types"
	perptypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/perpetuals/types"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/subaccounts/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/subaccounts/types"
	"github.com/stretchr/testify/require"
)

func TestGetCollateralPoolStateTransition(t *testing.T) {
	tests := map[string]struct {
		// parameters
		settledUpdateWithUpdatedSubaccount keeper.SettledUpdate
		perpetuals                         []perptypes.Perpetual

		// expectation
		expectedStateTransition *types.IsolatedPerpetualPositionStateTransition
		expectedErr             error
	}{
		`If no perpetual updates, nil state transition is returned`: {
			settledUpdateWithUpdatedSubaccount: keeper.SettledUpdate{
				SettledSubaccount: types.Subaccount{
					Id:                 &constants.Alice_Num0,
					PerpetualPositions: nil,
					AssetPositions:     nil,
				},
				PerpetualUpdates: nil,
				AssetUpdates:     nil,
			},
			perpetuals:              nil,
			expectedStateTransition: nil,
		},
		`A perpetual position is closed`: {
			settledUpdateWithUpdatedSubaccount: keeper.SettledUpdate{
				SettledSubaccount: types.Subaccount{
					Id:                 &constants.Alice_Num0,
					PerpetualPositions: nil,
					AssetPositions:     nil,
				},
				PerpetualUpdates: []types.PerpetualUpdate{
					{
						PerpetualId:      uint32(0),
						BigQuantumsDelta: big.NewInt(-100),
					},
				},
				AssetUpdates: nil,
			},
			perpetuals: []perptypes.Perpetual{
				constants.BtcUsd_100PercentMarginRequirement,
			},
			expectedStateTransition: &types.IsolatedPerpetualPositionStateTransition{
				SubaccountId: &constants.Alice_Num0,
				PerpetualId:  uint32(0),
				AssetIds:     []uint32{},
				BigQuantums:  []*big.Int{},
				Transition:   types.Closed,
			},
		},
		`Multiple perpetual positions are closed`: {
			settledUpdateWithUpdatedSubaccount: keeper.SettledUpdate{
				SettledSubaccount: types.Subaccount{
					Id:                 &constants.Alice_Num0,
					PerpetualPositions: nil,
					AssetPositions:     nil,
				},
				PerpetualUpdates: []types.PerpetualUpdate{
					{
						PerpetualId:      uint32(0),
						BigQuantumsDelta: big.NewInt(-100),
					},
					{
						PerpetualId:      uint32(1),
						BigQuantumsDelta: big.NewInt(-200),
					},
				},
				AssetUpdates: nil,
			},
			perpetuals: []perptypes.Perpetual{
				constants.BtcUsd_100PercentMarginRequirement,
				constants.EthUsd_NoMarginRequirement,
			},
			expectedStateTransition: &types.IsolatedPerpetualPositionStateTransition{
				SubaccountId: &constants.Alice_Num0,
				PerpetualId:  uint32(0),
				AssetIds:     []uint32{},
				BigQuantums:  []*big.Int{},
				Transition:   types.Closed,
			},
		},
		`Subaccount already has perpetual positions`: {
			settledUpdateWithUpdatedSubaccount: keeper.SettledUpdate{
				SettledSubaccount: types.Subaccount{
					Id: &constants.Alice_Num0,
					PerpetualPositions: []*types.PerpetualPosition{
						&constants.PerpetualPosition_OneBTCLong,
						&constants.PerpetualPosition_OneTenthEthLong,
					},
					AssetPositions: nil,
				},
				PerpetualUpdates: []types.PerpetualUpdate{
					{
						PerpetualId:      uint32(0),
						BigQuantumsDelta: big.NewInt(-100),
					},
				},
				AssetUpdates: nil,
			},
			perpetuals: []perptypes.Perpetual{
				constants.BtcUsd_100PercentMarginRequirement,
				constants.EthUsd_NoMarginRequirement,
			},
			expectedStateTransition: nil,
		},
		`Position closed with asset updates`: {
			settledUpdateWithUpdatedSubaccount: keeper.SettledUpdate{
				SettledSubaccount: types.Subaccount{
					Id:                 &constants.Alice_Num0,
					PerpetualPositions: nil,
					AssetPositions: []*types.AssetPosition{
						&constants.TDai_Asset_10_000,
					},
				},
				PerpetualUpdates: []types.PerpetualUpdate{
					{
						PerpetualId:      uint32(3),
						BigQuantumsDelta: big.NewInt(-1_000_000_000),
					},
				},
				AssetUpdates: []types.AssetUpdate{
					{
						AssetId:          assettypes.AssetTDai.Id,
						BigQuantumsDelta: big.NewInt(100_000_000),
					},
				},
			},
			perpetuals: []perptypes.Perpetual{
				constants.IsoUsd_IsolatedMarket,
			},
			expectedStateTransition: &types.IsolatedPerpetualPositionStateTransition{
				SubaccountId: &constants.Alice_Num0,
				PerpetualId:  uint32(3),
				AssetIds:     []uint32{assettypes.AssetTDai.Id},
				BigQuantums:  []*big.Int{constants.TDai_Asset_10_000.GetBigQuantums()},
				Transition:   types.Closed,
			},
		},
		`Position opened with asset updates`: {
			settledUpdateWithUpdatedSubaccount: keeper.SettledUpdate{
				SettledSubaccount: types.Subaccount{
					Id: &constants.Alice_Num0,
					PerpetualPositions: []*types.PerpetualPosition{
						&constants.PerpetualPosition_OneISOLong,
					},
					AssetPositions: []*types.AssetPosition{
						{
							AssetId:  assettypes.AssetTDai.Id,
							Quantums: dtypes.NewInt(-40_000_000), // -$40
						},
					},
				},
				PerpetualUpdates: []types.PerpetualUpdate{
					{
						PerpetualId:      uint32(3),
						BigQuantumsDelta: big.NewInt(1_000_000_000), // 1 ISO
					},
				},
				AssetUpdates: []types.AssetUpdate{
					{
						AssetId:          assettypes.AssetTDai.Id,
						BigQuantumsDelta: big.NewInt(-50_000_000), // -$50
					},
				},
			},
			perpetuals: []perptypes.Perpetual{
				constants.IsoUsd_IsolatedMarket,
			},
			expectedStateTransition: &types.IsolatedPerpetualPositionStateTransition{
				SubaccountId: &constants.Alice_Num0,
				PerpetualId:  uint32(3),
				AssetIds:     []uint32{assettypes.AssetTDai.Id},
				BigQuantums:  []*big.Int{big.NewInt(10_000_000)}, // $-40 - (-$50)
				Transition:   types.Opened,
			},
		},
		`Position size in increassed but already has perpetual position`: {
			settledUpdateWithUpdatedSubaccount: keeper.SettledUpdate{
				SettledSubaccount: types.Subaccount{
					Id: &constants.Alice_Num0,
					PerpetualPositions: []*types.PerpetualPosition{
						&constants.PerpetualPosition_OneISOLong,
					},
					AssetPositions: []*types.AssetPosition{
						{
							AssetId:  assettypes.AssetTDai.Id,
							Quantums: dtypes.NewInt(-40_000_000), // -$65
						},
					},
				},
				PerpetualUpdates: []types.PerpetualUpdate{
					{
						PerpetualId:      uint32(3),
						BigQuantumsDelta: big.NewInt(500_000_000), // 0.5 ISO
					},
				},
				AssetUpdates: []types.AssetUpdate{
					{
						AssetId:          assettypes.AssetTDai.Id,
						BigQuantumsDelta: big.NewInt(-25_000_000), // -$25
					},
				},
			},
			perpetuals: []perptypes.Perpetual{
				constants.IsoUsd_IsolatedMarket,
			},
			expectedStateTransition: nil,
		},
		`Returns error if perpetual position was opened with no asset updates`: {
			settledUpdateWithUpdatedSubaccount: keeper.SettledUpdate{
				SettledSubaccount: types.Subaccount{
					Id: &constants.Alice_Num0,
					PerpetualPositions: []*types.PerpetualPosition{
						&constants.PerpetualPosition_OneISOLong,
					},
					AssetPositions: []*types.AssetPosition{
						{
							AssetId:  assettypes.AssetTDai.Id,
							Quantums: dtypes.NewInt(50_000_000), // $50
						},
					},
				},
				PerpetualUpdates: []types.PerpetualUpdate{
					{
						PerpetualId:      uint32(3),
						BigQuantumsDelta: big.NewInt(1_000_000_000), // 1 ISO
					},
				},
				AssetUpdates: nil,
			},
			perpetuals: []perptypes.Perpetual{
				constants.IsoUsd_IsolatedMarket,
			},
			expectedStateTransition: nil,
			expectedErr:             types.ErrFailedToUpdateSubaccounts,
		},
		`Returns error if perpetual position was opened with multiple asset updates`: {
			settledUpdateWithUpdatedSubaccount: keeper.SettledUpdate{
				SettledSubaccount: types.Subaccount{
					Id: &constants.Alice_Num0,
					PerpetualPositions: []*types.PerpetualPosition{
						&constants.PerpetualPosition_OneISOLong,
					},
					AssetPositions: []*types.AssetPosition{
						{
							AssetId:  assettypes.AssetTDai.Id,
							Quantums: dtypes.NewInt(-40_000_000), // -$40
						},
						{
							AssetId:  constants.BtcUsd.Id,
							Quantums: dtypes.NewInt(100_000_000), // 1 BTC
						},
					},
				},
				PerpetualUpdates: []types.PerpetualUpdate{
					{
						PerpetualId:      uint32(3),
						BigQuantumsDelta: big.NewInt(1_000_000_000), // 1 ISO
					},
				},
				AssetUpdates: []types.AssetUpdate{
					{
						AssetId:          assettypes.AssetTDai.Id,
						BigQuantumsDelta: big.NewInt(-50_000_000), // -$50
					},
					{
						AssetId:          constants.BtcUsd.Id,
						BigQuantumsDelta: big.NewInt(100_000_000), // 1 BTC
					},
				},
			},
			perpetuals: []perptypes.Perpetual{
				constants.IsoUsd_IsolatedMarket,
			},
			expectedStateTransition: nil,
			expectedErr:             types.ErrFailedToUpdateSubaccounts,
		},
		`Returns error if perpetual position was opened with non-TDai asset update`: {
			settledUpdateWithUpdatedSubaccount: keeper.SettledUpdate{
				SettledSubaccount: types.Subaccount{
					Id: &constants.Alice_Num0,
					PerpetualPositions: []*types.PerpetualPosition{
						&constants.PerpetualPosition_OneISOLong,
					},
					AssetPositions: []*types.AssetPosition{
						{
							AssetId:  assettypes.AssetTDai.Id,
							Quantums: dtypes.NewInt(50_000_000), // $50
						},
						{
							AssetId:  constants.BtcUsd.Id,
							Quantums: dtypes.NewInt(100_000_000), // 1 BTC
						},
					},
				},
				PerpetualUpdates: []types.PerpetualUpdate{
					{
						PerpetualId:      uint32(3),
						BigQuantumsDelta: big.NewInt(1_000_000_000), // 1 ISO
					},
				},
				AssetUpdates: []types.AssetUpdate{
					{
						AssetId:          constants.BtcUsd.Id,
						BigQuantumsDelta: big.NewInt(100_000_000), // 1 BTC
					},
				},
			},
			perpetuals: []perptypes.Perpetual{
				constants.IsoUsd_IsolatedMarket,
			},
			expectedStateTransition: nil,
			expectedErr:             types.ErrFailedToUpdateSubaccounts,
		},
	}

	for name, tc := range tests {
		t.Run(
			name, func(t *testing.T) {
				stateTransition, err := keeper.GetCollateralPoolStateTransition(
					tc.settledUpdateWithUpdatedSubaccount,
					0,
				)
				if tc.expectedErr != nil {
					require.Error(t, tc.expectedErr, err)
				} else {
					require.NoError(t, err)
					require.Equal(t, tc.expectedStateTransition, stateTransition)
				}
			},
		)
	}
}
