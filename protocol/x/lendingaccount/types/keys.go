package types

// Module name and store keys
const (
	// ModuleName defines the module name
	ModuleName = "lendingaccount"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	LendingAccountKeyPrefix = "lending_account"

	LendingManagerKeyPrefix = "lending_manager"

	LendingInterfaceKeyPrefix = "lending_interface"

	TotalDebtKeyPrefix = "total_debt"

	BorrowedInBlockKeyPrefix = "borrowed_in_block"

	BlockLastBorrowedKeyPrefix = "block_last_borrowed"

	LendingManagerSybilResistanceKeyPrefix = "lending_manager_sybil_resistance"
)
