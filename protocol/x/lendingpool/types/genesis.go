package types

// DefaultGenesis returns the default lendingpool genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PoolParams: DefaultPoolParams(),
	}
}

// DefaultPoolParams returns the standard pool params for long-term operation of the network.
func DefaultPoolParams() []PoolParams {
	return []PoolParams{
		{
			TokenDenom:                 "ibc/BTC", // placeholders todo
			MaxPoolLiquidity:           "1000000000000",
			OptimalUtilizationRatio:    "8000",
			BaseRate:                   "200",
			SlopeOneRate:               "300",
			SlopeTwoRate:               "500",
			PermissionedCreditAccounts: []string{"defaultcreditaccount"},
			IsIsolated:                 false,
		},
		{
			TokenDenom:                 "ibc/ETH", // placeholders todo
			MaxPoolLiquidity:           "2000000000000",
			OptimalUtilizationRatio:    "7500",
			BaseRate:                   "250",
			SlopeOneRate:               "350",
			SlopeTwoRate:               "550",
			PermissionedCreditAccounts: []string{"defaultcreditaccount"},
			IsIsolated:                 false,
		},
		{
			TokenDenom:                 "ibc/DAI", // placeholders todo
			MaxPoolLiquidity:           "3000000000000",
			OptimalUtilizationRatio:    "7000",
			BaseRate:                   "300",
			SlopeOneRate:               "400",
			SlopeTwoRate:               "600",
			PermissionedCreditAccounts: []string{"defaultcreditaccount"},
			IsIsolated:                 false,
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	for _, pool := range gs.PoolParams {
		if _, err := pool.Validate(); err != nil {
			return err
		}
	}
	return nil
}
