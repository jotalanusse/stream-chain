package types

import (
	"math/big"

	errorsmod "cosmossdk.io/errors"
)

// Validate validates perpetual module's parameters.
func (params PoolParams) Validate() error {
	if params.TokenDenom != BTC_POOL && params.TokenDenom != ETH_POOL && params.TokenDenom != DAI_POOL {
		return ErrInvalidTokenDenom
	}

	bigMaxPoolLiquidity, err := ConvertStringToBigInt(params.MaxPoolLiquidity)
	if err != nil {
		return err
	}

	if bigMaxPoolLiquidity.Cmp(big.NewInt(0)) <= 0 {
		return ErrMaxPoolLiquidityIsZeroOrLess
	}

	bigWithdrawFee, err := ConvertStringToBigInt(params.WithdrawFee)
	if err != nil {
		return err
	}

	if bigWithdrawFee.Cmp(big.NewInt(0)) < 0 || bigWithdrawFee.Cmp(big.NewInt(PERCENTAGE_PRECISION)) >= 0 {
		return ErrWithdrawFeeOutOfRange
	}

	bigOptimalUtilizationRatio, err := ConvertStringToBigInt(params.OptimalUtilizationRatio)
	if err != nil {
		return err
	}

	if bigOptimalUtilizationRatio.Cmp(big.NewInt(0)) < 0 || bigOptimalUtilizationRatio.Cmp(big.NewInt(PERCENTAGE_PRECISION)) >= 0 {
		return ErrOptimalUtilizationRatioOutOfRange
	}

	bigBaseRate, err := ConvertStringToBigInt(params.BaseRate)
	if err != nil {
		return err
	}

	if bigBaseRate.Cmp(big.NewInt(0)) < 0 || bigBaseRate.Cmp(big.NewInt(PERCENTAGE_PRECISION)) >= 0 {
		return ErrBaseRateOutOfRange
	}

	bigSlopeOneRate, err := ConvertStringToBigInt(params.SlopeOneRate)
	if err != nil {
		return err
	}

	if bigSlopeOneRate.Cmp(big.NewInt(0)) < 0 || bigSlopeOneRate.Cmp(big.NewInt(PERCENTAGE_PRECISION)) >= 0 {
		return ErrSlopeOneRateOutOfRange
	}

	bigSlopeTwoRate, err := ConvertStringToBigInt(params.SlopeTwoRate)
	if err != nil {
		return err
	}

	// we dont upper bound R_2
	if bigSlopeTwoRate.Cmp(big.NewInt(0)) < 0 {
		return ErrSlopeTwoRateOutOfRange
	}

	return nil
}

// ApplyDecimalConversions converts the pool params to the correct decimal places
func (params *PoolParams) ApplyDecimalConversions() error {

	bigEighteenDecimals := big.NewInt(EIGHTEEN_DECIMALS)
	bigTwentySevenDecimals := big.NewInt(TWENTY_SEVEN_DECIMALS)

	bigOptimalUtilizationRatio, err := ConvertStringToBigInt(params.OptimalUtilizationRatio)
	if err != nil {
		return err
	}

	params.OptimalUtilizationRatio = PercentMul(bigOptimalUtilizationRatio, bigEighteenDecimals).String()

	bigBaseRate, err := ConvertStringToBigInt(params.BaseRate)
	if err != nil {
		return err
	}

	params.BaseRate = PercentMul(bigBaseRate, bigTwentySevenDecimals).String()

	bigSlopeOneRate, err := ConvertStringToBigInt(params.SlopeOneRate)
	if err != nil {
		return err
	}

	params.SlopeOneRate = PercentMul(bigSlopeOneRate, bigTwentySevenDecimals).String()

	bigSlopeTwoRate, err := ConvertStringToBigInt(params.SlopeTwoRate)
	if err != nil {
		return err
	}

	params.SlopeTwoRate = PercentMul(bigSlopeTwoRate, bigTwentySevenDecimals).String()

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
