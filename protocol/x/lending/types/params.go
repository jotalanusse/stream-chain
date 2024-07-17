package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeyPoolParams = []byte("PoolParams")
)

// ParamKeyTable for lending module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable(
		paramtypes.NewParamSetPair(KeyPoolParams, &PoolParams{}, validatePoolParams),
	)
}

// DefaultPoolParams returns the default pool parameters.
func DefaultPoolParams() PoolParams {
	return PoolParams{
		AssetDenom: "stake",
		InterestRateModel: &InterestRateModel{
			BaseRate:        0.02, // 2%
			Multiplier:      0.1,  // 10%
			JumpMultiplier:  0.5,  // 50%
			TargetThreshold: 0.9,  // 90% utilization
		},
	}
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of lending module's parameters.
func (p *PoolParams) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyPoolParams, p, validatePoolParams),
	}
}

// validatePoolParams validates the pool parameters.
func validatePoolParams(i interface{}) error {
	params, ok := i.(PoolParams)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if err := sdk.ValidateDenom(params.AssetDenom); err != nil {
		return err
	}

	if params.InterestRateModel.BaseRate < 0 {
		return fmt.Errorf("base rate should be non-negative: %f", params.InterestRateModel.BaseRate)
	}

	if params.InterestRateModel.Multiplier < 0 {
		return fmt.Errorf("multiplier should be non-negative: %f", params.InterestRateModel.Multiplier)
	}

	if params.InterestRateModel.JumpMultiplier < 0 {
		return fmt.Errorf("jump multiplier should be non-negative: %f", params.InterestRateModel.JumpMultiplier)
	}

	if params.InterestRateModel.TargetThreshold < 0 || params.InterestRateModel.TargetThreshold > 1 {
		return fmt.Errorf("target threshold should be between 0 and 1: %f", params.InterestRateModel.TargetThreshold)
	}

	return nil
}
