package types_test

import (
	"testing"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/perpetuals/types"
	"github.com/stretchr/testify/require"
)

func TestCollateralPoolValidate(t *testing.T) {
	tests := map[string]struct {
		IsolatedMarketMaxCumulativeInsuranceFundDeltaPerBlock uint64
		multiCollateralAssets                                 *types.MultiCollateralAssetsArray
		quoteAssetId                                          uint32
		expectedError                                         error
	}{
		"Validates successfully": {
			IsolatedMarketMaxCumulativeInsuranceFundDeltaPerBlock: 100_000,
			multiCollateralAssets: &types.MultiCollateralAssetsArray{
				MultiCollateralAssets: []uint32{1, 2, 3},
			},
			quoteAssetId:  2,
			expectedError: nil,
		},
		"Failure: Isolated Market Max Cumulative Insurance Fund Delta Per Block is zero": {
			IsolatedMarketMaxCumulativeInsuranceFundDeltaPerBlock: 0,
			multiCollateralAssets: &types.MultiCollateralAssetsArray{
				MultiCollateralAssets: []uint32{1, 2, 3},
			},
			quoteAssetId:  2,
			expectedError: types.ErrIsolatedMarketMaxCumulativeInsuranceFundDeltaPerBlockZero,
		},
		"Failure: Isolated Market Multi Collateral Assets is empty": {
			IsolatedMarketMaxCumulativeInsuranceFundDeltaPerBlock: 100_000,
			multiCollateralAssets: &types.MultiCollateralAssetsArray{},
			quoteAssetId:          2,
			expectedError:         types.ErrIsolatedMarketMultiCollateralAssetsEmpty,
		},
		"Failure: Isolated Market Multi Collateral Asset Does Not Contain Quote Asset": {
			IsolatedMarketMaxCumulativeInsuranceFundDeltaPerBlock: 100_000,
			multiCollateralAssets: &types.MultiCollateralAssetsArray{MultiCollateralAssets: []uint32{1, 2, 3}},
			quoteAssetId:          4,
			expectedError:         types.ErrIsolatedMarketMultiCollateralAssetDoesNotContainQuoteAsset,
		},
	}

	// Run tests.
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			collateralPool := &types.CollateralPool{
				CollateralPoolId: 1,
				IsolatedMarketMaxCumulativeInsuranceFundDeltaPerBlock: tc.IsolatedMarketMaxCumulativeInsuranceFundDeltaPerBlock,
				IsolatedMarketMultiCollateralAssets:                   tc.multiCollateralAssets,
				QuoteAssetId:                                          tc.quoteAssetId,
			}

			err := collateralPool.Validate()
			if tc.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
