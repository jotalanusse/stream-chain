package types

import "math/big"

const (
	// address of the lending pool insurance fund
	LENDING_POOL_INSURANCE_FUND = "lendingpoolinsurancefund"
)

var (
	// all percentage values are in the range [0, PERCENTAGE_PRECISION]
	PERCENTAGE_PRECISION = big.NewInt(10000)

	HALF_PERCENT = big.NewInt(5000)

	EIGHTEEN_DECIMALS = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	TWENTY_SEVEN_DECIMALS = new(big.Int).Exp(big.NewInt(10), big.NewInt(27), nil)

	SECONDS_PER_YEAR = big.NewInt(31536000)
)

// GetLendingTokenDenom concatenates the input denom with "-lp" to form the lending token denom.
func GetLendingTokenDenom(denom string) string {
	return denom + "-lp"
}
