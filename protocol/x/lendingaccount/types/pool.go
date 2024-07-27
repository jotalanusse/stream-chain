package types

import "math/big"

const ()

var (
	PERCENTAGE_FACTOR = big.NewInt(1000)
)

// GetLendingTokenDenom concatenates the input denom with "-lp" to form the lending token denom.
func GetLendingInterface(managerName string) string {
	return managerName + "-interface"
}
