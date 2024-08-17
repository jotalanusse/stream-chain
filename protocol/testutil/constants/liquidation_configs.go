package constants

import (
	"math/big"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/dtypes"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
	big_testutil "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/big"
	clobtypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/clob/types"
)

var (
	// Block limits.
	FillablePriceConfig_Default = clobtypes.FillablePriceConfig{
		BankruptcyAdjustmentPpm:           lib.OneMillion,
		SpreadToMaintenanceMarginRatioPpm: 100_000,
	}
	FillablePriceConfig_Max_Smmr = clobtypes.FillablePriceConfig{
		BankruptcyAdjustmentPpm:           lib.OneMillion,
		SpreadToMaintenanceMarginRatioPpm: lib.OneMillion,
	}
	PositionBlockLimits_Default = clobtypes.PositionBlockLimits{
		MinPositionNotionalLiquidated: dtypes.NewIntFromBigInt(
			big_testutil.MustFirst(new(big.Int).SetString("1000", 10)), // $0.001
		),
		MaxPositionPortionLiquidatedPpm: 1_000_000,
	}
	SubaccountBlockLimits_Default = clobtypes.SubaccountBlockLimits{
		MaxNotionalLiquidated: dtypes.NewIntFromBigInt(
			big_testutil.MustFirst(new(big.Int).SetString("100000000000000", 10)),
		),
		MaxQuantumsInsuranceLost: dtypes.NewIntFromBigInt(
			big_testutil.MustFirst(new(big.Int).SetString("100000000000000", 10)),
		),
	}
	PositionBlockLimits_No_Limit = clobtypes.PositionBlockLimits{
		MinPositionNotionalLiquidated: dtypes.NewIntFromBigInt(
			big_testutil.MustFirst(new(big.Int).SetString("1", 10)), // $0.000001
		),
		MaxPositionPortionLiquidatedPpm: lib.OneMillion,
	}
	SubaccountBlockLimits_No_Limit = clobtypes.SubaccountBlockLimits{
		MaxNotionalLiquidated:    dtypes.MaxUint256SerializableInt(),
		MaxQuantumsInsuranceLost: dtypes.MaxUint256SerializableInt(),
	}
	// Liquidation Configs.
	LiquidationsConfig_No_Limit = clobtypes.LiquidationsConfig{
		MaxLiquidationFeePpm:  5_000,
		FillablePriceConfig:   FillablePriceConfig_Default,
		PositionBlockLimits:   PositionBlockLimits_No_Limit,
		SubaccountBlockLimits: SubaccountBlockLimits_No_Limit,
	}
	LiquidationsConfig_FillablePrice_Max_Smmr = clobtypes.LiquidationsConfig{
		MaxLiquidationFeePpm:  5_000,
		FillablePriceConfig:   FillablePriceConfig_Max_Smmr,
		PositionBlockLimits:   PositionBlockLimits_No_Limit,
		SubaccountBlockLimits: SubaccountBlockLimits_No_Limit,
	}
	LiquidationsConfig_Position_Min10m_Max05mPpm = clobtypes.LiquidationsConfig{
		MaxLiquidationFeePpm: 5_000,
		FillablePriceConfig:  FillablePriceConfig_Default,
		PositionBlockLimits: clobtypes.PositionBlockLimits{
			MinPositionNotionalLiquidated: dtypes.NewIntFromBigInt(
				big_testutil.MustFirst(new(big.Int).SetString("10000000", 10)), // $10
			),
			MaxPositionPortionLiquidatedPpm: 500_000,
		},
		SubaccountBlockLimits: SubaccountBlockLimits_No_Limit,
	}
	LiquidationsConfig_Subaccount_Max10bNotionalLiquidated_Max10bInsuranceLost = clobtypes.LiquidationsConfig{
		MaxLiquidationFeePpm: 5_000,
		FillablePriceConfig:  FillablePriceConfig_Default,
		PositionBlockLimits:  PositionBlockLimits_No_Limit,
		SubaccountBlockLimits: clobtypes.SubaccountBlockLimits{
			MaxNotionalLiquidated: dtypes.NewIntFromBigInt(
				big_testutil.MustFirst(new(big.Int).SetString("10000000000", 10)), // $10,000
			),
			MaxQuantumsInsuranceLost: dtypes.NewIntFromBigInt(
				big_testutil.MustFirst(new(big.Int).SetString("10000000000", 10)), // $10,000
			),
		},
	}
)
