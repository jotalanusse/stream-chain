package types

// DepositType defines a type for different deposit types.
type DepositType string

const (
	// BTCDeposit represents a deposit type of BTC
	BTCDeposit DepositType = "BTC"
	// ETHDeposit represents a deposit type of ETH
	ETHDeposit DepositType = "ETH"
	// USDCDeposit represents a deposit type of USDC
	USDCDeposit DepositType = "USDC"
	// Future deposit types can be added here
)

// BorrowType defines a type for different borrow types.
type BorrowType string

const (
	// BTCBorrow represents a borrow type of BTC
	BTCBorrow BorrowType = "BTC"
	// ETHBorrow represents a borrow type of ETH
	ETHBorrow BorrowType = "ETH"
	// USDCBorrow represents a borrow type of USDC
	USDCBorrow BorrowType = "USDC"
	// Future borrow types can be added here
)
