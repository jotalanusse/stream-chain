package types

// DONTCOVER

import errorsmod "cosmossdk.io/errors"

var (
	ErrUnableToDecodeBigInt = errorsmod.Register(
		ModuleName,
		2,
		"Unable to decode big int",
	)
	ErrInvalidTokenDenom = errorsmod.Register(
		ModuleName,
		3,
		"Invalid token denomination",
	)
	ErrMaxPoolLiquidityIsZeroOrLess = errorsmod.Register(
		ModuleName,
		4,
		"Max pool liquidity is zero or less",
	)
	ErrWithdrawFeeOutOfRange = errorsmod.Register(
		ModuleName,
		5,
		"Withdraw fee is out of range",
	)
	ErrOptimalUtilizationRatioOutOfRange = errorsmod.Register(
		ModuleName,
		6,
		"Optimal utilization ratio is out of range",
	)
	ErrBaseRateOutOfRange = errorsmod.Register(
		ModuleName,
		7,
		"Base rate is out of range",
	)
	ErrSlopeOneRateOutOfRange = errorsmod.Register(
		ModuleName,
		8,
		"Slope one rate is out of range",
	)
	ErrSlopeTwoRateOutOfRange = errorsmod.Register(
		ModuleName,
		9,
		"Slope two rate is out of range",
	)
	ErrPoolParamsAlreadyExists = errorsmod.Register(
		ModuleName,
		10,
		"Pool parameters already exist",
	)
)
