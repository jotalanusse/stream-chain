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
	ErrInvalidAddress = errorsmod.Register(
		ModuleName,
		11,
		"Invalid address",
	)
	ErrPoolMoreThanMaxLiquidityLimit = errorsmod.Register(
		ModuleName,
		12,
		"Pool exceed the max liquidity limit",
	)
	ErrPermissionedCreditAccountsEmpty = errorsmod.Register(
		ModuleName,
		13,
		"Permissioned credit accounts are empty",
	)
	ErrCreditAccountNotPermissioned = errorsmod.Register(
		ModuleName,
		14,
		"Credit account is not permissioned",
	)
	ErrInvalidRepayFromCreditAccount = errorsmod.Register(
		ModuleName,
		15,
		"Invalid repay from credit account",
	)
)
