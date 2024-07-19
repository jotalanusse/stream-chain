package types

// DONTCOVER

import errorsmod "cosmossdk.io/errors"

// x/lending module sentinel errors
var (
	ErrInvalidAddress = errorsmod.Register(
		ModuleName,
		1,
		"Invalid Address",
	)
	ErrAccountAlreadyExists = errorsmod.Register(
		ModuleName,
		2,
		"Invalid Address",
	)
	ErrAccountNotFound = errorsmod.Register(
		ModuleName,
		3,
		"Address not found",
	)
	ErrInvalidPositionId = errorsmod.Register(
		ModuleName,
		4,
		"Invalid Position ID",
	)
	ErrInvalidAmount = errorsmod.Register(
		ModuleName,
		5,
		"Invalid Amount",
	)
)
