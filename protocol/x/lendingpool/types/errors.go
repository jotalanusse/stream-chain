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
	ErrOptimalUtilizationRatioOutOfRange = errorsmod.Register(
		ModuleName,
		5,
		"Optimal utilization ratio is out of range",
	)
	ErrBaseRateOutOfRange = errorsmod.Register(
		ModuleName,
		6,
		"Base rate is out of range",
	)
	ErrSlopeOneRateOutOfRange = errorsmod.Register(
		ModuleName,
		7,
		"Slope one rate is out of range",
	)
	ErrSlopeTwoRateOutOfRange = errorsmod.Register(
		ModuleName,
		8,
		"Slope two rate is out of range",
	)
	ErrPoolParamsAlreadyExists = errorsmod.Register(
		ModuleName,
		9,
		"Pool parameters already exist",
	)
	ErrInvalidAddress = errorsmod.Register(
		ModuleName,
		10,
		"Invalid address",
	)
	ErrPoolMoreThanMaxLiquidityLimit = errorsmod.Register(
		ModuleName,
		11,
		"Pool exceed the max liquidity limit",
	)
	ErrPermissionedCreditAccountsEmpty = errorsmod.Register(
		ModuleName,
		12,
		"Permissioned credit accounts are empty",
	)
	ErrCreditAccountNotPermissioned = errorsmod.Register(
		ModuleName,
		13,
		"Credit account is not permissioned",
	)
	ErrInvalidRepayFromCreditAccount = errorsmod.Register(
		ModuleName,
		14,
		"Invalid repay from credit account",
	)
)
