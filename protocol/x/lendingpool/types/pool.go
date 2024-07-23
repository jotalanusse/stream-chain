package types

const (
	// all percentage values are in the range [0, PERCENTAGE_PRECISION]
	PERCENTAGE_PRECISION = 10000

	HALF_PERCENT = 5000

	EIGHTEEN_DECIMALS = 10 ^ 18

	TWENTY_SEVEN_DECIMALS = 10 ^ 27

	SECONDS_PER_YEAR = 31536000

	// todo placeholders for ibc denoms of btc, eth, dai
	BTC_POOL = "BTC"

	ETH_POOL = "ETH"

	DAI_POOL = "DAI"

	// address of the lending pool insurance fund
	LENDING_POOL_INSURANCE_FUND = "lendingpoolinsurancefund"
)

// GetLendingTokenDenom concatenates the input denom with "-lp" to form the lending token denom.
func GetLendingTokenDenom(denom string) string {
	return denom + "-lp"
}
