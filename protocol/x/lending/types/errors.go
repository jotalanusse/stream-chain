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
		1,
		"Invalid Address",
	)
	ErrAccountNotFound = errorsmod.Register(
		ModuleName,
		1,
		"Address not found",
	)
)
