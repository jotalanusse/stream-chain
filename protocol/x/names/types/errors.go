package types

// DONTCOVER

import errorsmod "cosmossdk.io/errors"

// x/names module sentinel errors
var (
	ErrAssetDoesNotExist            = errorsmod.Register(ModuleName, 1, "Asset does not exist")
	ErrNoAssetWithDenom             = errorsmod.Register(ModuleName, 3, "No asset found associated with given denom")
	ErrNameNameAlreadyExists        = errorsmod.Register(ModuleName, 4, "Existing name found with the same name")
	ErrNameIdAlreadyExists          = errorsmod.Register(ModuleName, 5, "Existing name found with the same name id")
	ErrGapFoundInNameId             = errorsmod.Register(ModuleName, 6, "Found gap in name Id")
	ErrJotaMustBeNameZero           = errorsmod.Register(ModuleName, 7, "Jota must be name 0")
	ErrNoNameInGenesis              = errorsmod.Register(ModuleName, 8, "No name found in genesis state")
	ErrInvalidMarketId              = errorsmod.Register(ModuleName, 9, "Found market id for asset without market")
	ErrInvalidAssetAtomicResolution = errorsmod.Register(ModuleName, 10, "Invalid asset atomic resolution")
	ErrInvalidDenomExponent         = errorsmod.Register(ModuleName, 11, "Invalid denom exponent")
	ErrAssetAlreadyExists           = errorsmod.Register(ModuleName, 13, "Asset already exists")
	ErrUnexpectedTDaiDenomExponent  = errorsmod.Register(ModuleName, 14, "TDai denom exponent is unexpected")

	// Errors for Not Implemented
	ErrNotImplementedMulticollateral = errorsmod.Register(ModuleName, 401, "Not Implemented: Multi-Collateral")
	ErrNotImplementedMargin          = errorsmod.Register(ModuleName, 402, "Not Implemented: Margin-Trading of Assets")
)
