package types

import (
	"github.com/StreamFinance-Protocol/stream-chain/protocol/dtypes"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	LiquidationsConfig_Default = LiquidationsConfig{
		MaxLiquidationFeePpm: 5_000,
		FillablePriceConfig: FillablePriceConfig{
			BankruptcyAdjustmentPpm:           lib.OneMillion,
			SpreadToMaintenanceMarginRatioPpm: 100_000,
		},
		PositionBlockLimits: PositionBlockLimits{
			MinPositionNotionalLiquidated:   dtypes.NewIntFromString("1000"), // $0.001
			MaxPositionPortionLiquidatedPpm: 1_000_000,
		},
		SubaccountBlockLimits: SubaccountBlockLimits{
			MaxNotionalLiquidated:    dtypes.NewIntFromString("100_000_000_000_000"), // $100,000,000
			MaxQuantumsInsuranceLost: dtypes.NewIntFromString("100_000_000_000_000"), // $100,000,000
		},
	}
)

// LiquidationsConfigKeeper is an interface that encapsulates all reads and writes to the
// liquidation configuration values written to state.
type LiquidationsConfigKeeper interface {
	GetLiquidationsConfig(
		ctx sdk.Context,
	) LiquidationsConfig
}
