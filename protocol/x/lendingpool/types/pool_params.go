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
	// withdraw_fee is the fee charged for withdrawals from the pool send to the insurance fund.
	WithdrawFee *big.Int
	// optimal_utilization_ratio is the optimal ratio of utilization for the pool.
	OptimalUtilizationRatio *big.Int
	// base_rate is the base interest rate for the pool.
	BaseRate *big.Int
	// slope_one_rate is the interest rate slope for the first segment.
	SlopeOneRate *big.Int
	// slope_two_rate is the interest rate slope for the second segment.
	SlopeTwoRate               *big.Int
	PermissionedCreditAccounts []string
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

	bigWithdrawFee, err := ConvertStringToBigInt(params.WithdrawFee)
	if err != nil {
		return InternalPoolParams{}, err
	}

	if bigWithdrawFee.Cmp(big.NewInt(0)) < 0 || bigWithdrawFee.Cmp(big.NewInt(PERCENTAGE_PRECISION)) >= 0 {
		return InternalPoolParams{}, ErrWithdrawFeeOutOfRange
	}

	bigOptimalUtilizationRatio, err := ConvertStringToBigInt(params.OptimalUtilizationRatio)
	if err != nil {
		return InternalPoolParams{}, err
	}

	if bigOptimalUtilizationRatio.Cmp(big.NewInt(0)) < 0 || bigOptimalUtilizationRatio.Cmp(big.NewInt(PERCENTAGE_PRECISION)) >= 0 {
		return InternalPoolParams{}, ErrOptimalUtilizationRatioOutOfRange
	}

	bigBaseRate, err := ConvertStringToBigInt(params.BaseRate)
	if err != nil {
		return InternalPoolParams{}, err
	}

	if bigBaseRate.Cmp(big.NewInt(0)) < 0 || bigBaseRate.Cmp(big.NewInt(PERCENTAGE_PRECISION)) >= 0 {
		return InternalPoolParams{}, ErrBaseRateOutOfRange
	}

	bigSlopeOneRate, err := ConvertStringToBigInt(params.SlopeOneRate)
	if err != nil {
		return InternalPoolParams{}, err
	}

	if bigSlopeOneRate.Cmp(big.NewInt(0)) < 0 || bigSlopeOneRate.Cmp(big.NewInt(PERCENTAGE_PRECISION)) >= 0 {
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
		WithdrawFee:                bigWithdrawFee,
		OptimalUtilizationRatio:    bigOptimalUtilizationRatio,
		BaseRate:                   bigBaseRate,
		SlopeOneRate:               bigSlopeOneRate,
		SlopeTwoRate:               bigSlopeTwoRate,
		PermissionedCreditAccounts: creditAccounts,
	}

	return internalParams, nil
}

// ConvertInternalToPoolParams converts InternalPoolParams to PoolParams
func ConvertInternalToPoolParams(internalParams InternalPoolParams) PoolParams {
	return PoolParams{
		TokenDenom:                 internalParams.TokenDenom,
		MaxPoolLiquidity:           internalParams.MaxPoolLiquidity.String(),
		WithdrawFee:                internalParams.WithdrawFee.String(),
		OptimalUtilizationRatio:    internalParams.OptimalUtilizationRatio.String(),
		BaseRate:                   internalParams.BaseRate.String(),
		SlopeOneRate:               internalParams.SlopeOneRate.String(),
		SlopeTwoRate:               internalParams.SlopeTwoRate.String(),
		PermissionedCreditAccounts: internalParams.PermissionedCreditAccounts,
	}
}

// ApplyDecimalConversions converts the pool params to the correct decimal places
func (params *InternalPoolParams) ApplyDecimalConversions() error {

	bigEighteenDecimals := big.NewInt(EIGHTEEN_DECIMALS)
	bigTwentySevenDecimals := big.NewInt(TWENTY_SEVEN_DECIMALS)

	params.OptimalUtilizationRatio = PercentMul(params.OptimalUtilizationRatio, bigEighteenDecimals)

	params.BaseRate = PercentMul(params.BaseRate, bigTwentySevenDecimals)

	params.SlopeOneRate = PercentMul(params.SlopeOneRate, bigTwentySevenDecimals)

	params.SlopeTwoRate = PercentMul(params.SlopeTwoRate, bigTwentySevenDecimals)

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

func PercentMul(value, percentage *big.Int) *big.Int {
	if value.Cmp(big.NewInt(0)) == 0 || percentage.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	halfPercent := big.NewInt(HALF_PERCENT)
	percentageFactor := big.NewInt(PERCENTAGE_PRECISION)

	result := new(big.Int).Mul(value, percentage)
	result = result.Add(result, halfPercent)
	result = result.Div(result, percentageFactor)

	return result
}
