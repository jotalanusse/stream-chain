package types

// Module name and store keys
const (
	// ModuleName defines the module name
	ModuleName = "lendingpool"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	PoolParamsKeyPrefix = "poolparams"

	LastUpdatedTimePrefix = "lastupdatedtime"

	LastUpdatedTotalLiquidityPrefix = "lastupdatedtotalliquidity"

	CumulativeInterestRatePrefix = "cumulativeinterestrate"

	TotalBorrowedPrefix = "totalborrowed"

	CurrentBorrowAPYEighteenDecimalsPrefix = "currentborrowapy"
)
