package types_test

import (
	"testing"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/perpetuals/types"
	"github.com/stretchr/testify/require"
)

func TestPerpetualParams_Validate(t *testing.T) {
	tests := []struct {
		desc        string
		params      types.PerpetualParams
		expectedErr string
	}{
		{
			desc: "Valid param",
			params: types.PerpetualParams{
				Ticker:            "test",
				DefaultFundingPpm: 1_000_000,
			},
			expectedErr: "",
		},
		{
			desc: "Empty ticker",
			params: types.PerpetualParams{
				Ticker:            "",
				DefaultFundingPpm: 1_000_000,
			},
			expectedErr: "Ticker must be non-empty string",
		},
		{
			desc: "Invalid DefaultFundingPpm",
			params: types.PerpetualParams{
				Ticker:            "test",
				DefaultFundingPpm: 100_000_000,
			},
			expectedErr: "DefaultFundingPpm magnitude exceeds maximum value",
		},
		{
			desc: "Invalid Isolated market max cumulative insurance fund delta per block",
			params: types.PerpetualParams{
				Ticker:            "test",
				DefaultFundingPpm: 1_000_000,
				IsolatedMarketMultiCollateralAssets: &types.MultiCollateralAssetsArray{
					MultiCollateralAssets: []uint32{1},
				},
				QuoteAssetId: 1,
			},
			expectedErr: "isolated market max cumulative insurance fund delta per block is zero",
		},
		{
			desc: "Invalid Isolated market collateral assets",
			params: types.PerpetualParams{
				Ticker:            "test",
				DefaultFundingPpm: 1_000_000,
				IsolatedMarketMaxCumulativeInsuranceFundDeltaPerBlock: uint64(100),
				IsolatedMarketMultiCollateralAssets: &types.MultiCollateralAssetsArray{
					MultiCollateralAssets: []uint32{1},
				},
				QuoteAssetId:     0,
				CollateralPoolId: 0,
			},
			expectedErr: "multi collateral asset does not contain quote asset",
		},
		{
			desc: "Invalid Isolated market collateral assets",
			params: types.PerpetualParams{
				Ticker:            "test",
				DefaultFundingPpm: 1_000_000,
				IsolatedMarketMaxCumulativeInsuranceFundDeltaPerBlock: uint64(100),
				IsolatedMarketMultiCollateralAssets: &types.MultiCollateralAssetsArray{
					MultiCollateralAssets: []uint32{0},
				},
				QuoteAssetId:     1,
				CollateralPoolId: 0,
			},
			expectedErr: "multi collateral asset does not contain quote asset",
		},
		{
			desc: "Valid Isolated market",
			params: types.PerpetualParams{
				Ticker:            "test",
				DefaultFundingPpm: 1_000_000,
				IsolatedMarketMaxCumulativeInsuranceFundDeltaPerBlock: uint64(100),
				IsolatedMarketMultiCollateralAssets: &types.MultiCollateralAssetsArray{
					MultiCollateralAssets: []uint32{0},
				},
				QuoteAssetId:     0,
				CollateralPoolId: 0,
			},
			expectedErr: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.params.Validate()
			if tc.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tc.expectedErr)
			}
		})
	}
}
