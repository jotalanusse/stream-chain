package types_test

import (
	"testing"

	types "github.com/StreamFinance-Protocol/stream-chain/protocol/x/perpetuals/types"
	"github.com/stretchr/testify/require"
)

func TestMsgSetCollateralPool_ValidateBasic(t *testing.T) {
	tests := map[string]struct {
		msg         types.MsgSetCollateralPool
		expectedErr string
	}{
		"Success": {
			msg: types.MsgSetCollateralPool{
				Authority: validAuthority,
				CollateralPool: types.CollateralPool{
					CollateralPoolId:                        1,
					MaxCumulativeInsuranceFundDeltaPerBlock: 100_000,
					MarketMultiCollateralAssets:             &types.MultiCollateralAssetsArray{MultiCollateralAssets: []uint32{1, 2, 3}},
					QuoteAssetId:                            2,
				},
			},
		},
		"Failure: Invalid authority": {
			msg: types.MsgSetCollateralPool{
				Authority: "",
			},
			expectedErr: "Authority is invalid",
		},
		"Failure: Isolated Market Max Cumulative Insurance Fund Delta Per Block is zero": {
			msg: types.MsgSetCollateralPool{
				Authority: validAuthority,
				CollateralPool: types.CollateralPool{
					CollateralPoolId:                        1,
					MaxCumulativeInsuranceFundDeltaPerBlock: 0,
					MarketMultiCollateralAssets:             &types.MultiCollateralAssetsArray{MultiCollateralAssets: []uint32{1, 2, 3}},
					QuoteAssetId:                            2,
				},
			},
			expectedErr: "Isolated Market Max Cumulative Insurance Fund Delta Per Block is zero",
		},
		"Failure: Isolated Market Multi Collateral Assets is empty": {
			msg: types.MsgSetCollateralPool{
				Authority: validAuthority,
				CollateralPool: types.CollateralPool{
					CollateralPoolId:                        1,
					MaxCumulativeInsuranceFundDeltaPerBlock: 100_000,
					MarketMultiCollateralAssets:             &types.MultiCollateralAssetsArray{},
					QuoteAssetId:                            2,
				},
			},
			expectedErr: "Isolated Market Multi Collateral Assets is empty",
		},
		"Failure: Isolated Market Multi Collateral Asset Does Not Contain Quote Asset": {
			msg: types.MsgSetCollateralPool{
				Authority: validAuthority,
				CollateralPool: types.CollateralPool{
					CollateralPoolId:                        1,
					MaxCumulativeInsuranceFundDeltaPerBlock: 100_000,
					MarketMultiCollateralAssets:             &types.MultiCollateralAssetsArray{MultiCollateralAssets: []uint32{1, 2, 3}},
					QuoteAssetId:                            4,
				},
			},
			expectedErr: "Isolated Market Multi Collateral Asset Does Not Contain Quote Asset",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, tc.expectedErr)
			}
		})
	}
}
