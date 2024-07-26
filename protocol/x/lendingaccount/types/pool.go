package types

const ()

var ()

// GetLendingTokenDenom concatenates the input denom with "-lp" to form the lending token denom.
func GetLendingInterface(managerName string) string {
	return managerName + "-interface"
}
