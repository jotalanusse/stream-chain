package types

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
)

type InternalPoolParams struct {
	// token_denom is the denomination of the token in the pool.
	TokenDenom string
	// max_pool_liquidity is the maximum liquidity allowed in the pool.
	MaxPoolLiquidity *big.Int
	// optimal_utilization_ratio is the optimal ratio of utilization for the pool.
	OptimalUtilizationRatio *big.Int
	// base_rate is the base interest rate for the pool.
	BaseRate *big.Int
	// slope_one_rate is the interest rate slope for the first segment.
	SlopeOneRate *big.Int
	// slope_two_rate is the interest rate slope for the second segment.
	SlopeTwoRate *big.Int
	// permissioned_credit_accounts are the accounts that are allowed to borrow from the pool.
	PermissionedCreditAccounts []string
	// is_isolated is a flag to determine if the pool is isolated.
	IsIsolated bool
}

// Validate validates perpetual module's parameters.
func (params PoolParams) Validate() (internalParams InternalPoolParams, err error) {

	bigMaxPoolLiquidity, err := ConvertStringToBigInt(params.MaxPoolLiquidity)
	if err != nil {
		return InternalPoolParams{}, err
	}

	if bigMaxPoolLiquidity.Cmp(big.NewInt(0)) <= 0 {
		return InternalPoolParams{}, ErrMaxPoolLiquidityIsZeroOrLess
	}

	bigOptimalUtilizationRatio, err := ConvertStringToBigInt(params.OptimalUtilizationRatio)
	if err != nil {
		return InternalPoolParams{}, err
	}

	if bigOptimalUtilizationRatio.Cmp(big.NewInt(0)) < 0 || bigOptimalUtilizationRatio.Cmp(PERCENTAGE_PRECISION) >= 0 {
		return InternalPoolParams{}, ErrOptimalUtilizationRatioOutOfRange
	}

	bigBaseRate, err := ConvertStringToBigInt(params.BaseRate)
	if err != nil {
		return InternalPoolParams{}, err
	}

	if bigBaseRate.Cmp(big.NewInt(0)) < 0 || bigBaseRate.Cmp(PERCENTAGE_PRECISION) >= 0 {
		return InternalPoolParams{}, ErrBaseRateOutOfRange
	}

	bigSlopeOneRate, err := ConvertStringToBigInt(params.SlopeOneRate)
	if err != nil {
		return InternalPoolParams{}, err
	}

	if bigSlopeOneRate.Cmp(big.NewInt(0)) < 0 || bigSlopeOneRate.Cmp(PERCENTAGE_PRECISION) >= 0 {
		return InternalPoolParams{}, ErrSlopeOneRateOutOfRange
	}

	bigSlopeTwoRate, err := ConvertStringToBigInt(params.SlopeTwoRate)
	if err != nil {
		return InternalPoolParams{}, err
	}

	// we dont upper bound R_2
	if bigSlopeTwoRate.Cmp(big.NewInt(0)) < 0 {
		return InternalPoolParams{}, ErrSlopeTwoRateOutOfRange
	}

	creditAccounts := params.PermissionedCreditAccounts
	if len(creditAccounts) == 0 {
		return InternalPoolParams{}, ErrPermissionedCreditAccountsEmpty
	}

	internalParams = InternalPoolParams{
		TokenDenom:                 params.TokenDenom,
		MaxPoolLiquidity:           bigMaxPoolLiquidity,
		OptimalUtilizationRatio:    bigOptimalUtilizationRatio,
		BaseRate:                   bigBaseRate,
		SlopeOneRate:               bigSlopeOneRate,
		SlopeTwoRate:               bigSlopeTwoRate,
		PermissionedCreditAccounts: creditAccounts,
		IsIsolated:                 params.IsIsolated,
	}

	return internalParams, nil
}

// ConvertInternalToPoolParams converts InternalPoolParams to PoolParams
func ConvertInternalToPoolParams(internalParams InternalPoolParams) PoolParams {
	return PoolParams{
		TokenDenom:                 internalParams.TokenDenom,
		MaxPoolLiquidity:           internalParams.MaxPoolLiquidity.String(),
		OptimalUtilizationRatio:    internalParams.OptimalUtilizationRatio.String(),
		BaseRate:                   internalParams.BaseRate.String(),
		SlopeOneRate:               internalParams.SlopeOneRate.String(),
		SlopeTwoRate:               internalParams.SlopeTwoRate.String(),
		PermissionedCreditAccounts: internalParams.PermissionedCreditAccounts,
		IsIsolated:                 internalParams.IsIsolated,
	}
}

// ApplyDecimalConversions converts the pool params to the correct decimal places
func (params *InternalPoolParams) ApplyDecimalConversions() error {

	params.OptimalUtilizationRatio = PercentMultiply(params.OptimalUtilizationRatio, EIGHTEEN_DECIMALS)

	params.BaseRate = PercentMultiply(params.BaseRate, TWENTY_SEVEN_DECIMALS)

	params.SlopeOneRate = PercentMultiply(params.SlopeOneRate, TWENTY_SEVEN_DECIMALS)

	params.SlopeTwoRate = PercentMultiply(params.SlopeTwoRate, TWENTY_SEVEN_DECIMALS)

	return nil
}

func ConvertStringToBigInt(str string) (*big.Int, error) {

	bigint, ok := new(big.Int).SetString(str, 10)
	if !ok {
		return nil, errorsmod.Wrap(
			ErrUnableToDecodeBigInt,
			"Unable to convert the sDAI conversion rate to a big int",
		)
	}

	return bigint, nil
}

func PercentMultiply(value, percentage *big.Int) (result *big.Int) {

	if value.Cmp(big.NewInt(0)) == 0 || percentage.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	result = new(big.Int).Mul(value, percentage)
	result = result.Add(result, HALF_PERCENT) // to round up
	return result.Div(result, PERCENTAGE_PRECISION)
}
