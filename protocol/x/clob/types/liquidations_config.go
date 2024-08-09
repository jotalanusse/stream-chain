package types

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/dtypes"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
)

// Validate validates each individual field of the liquidations config for validity.
// It returns an error if any of the liquidation config fields fail the following validation:
// - `maxLiquidationFee == 0 || maxLiquidationFee > 1_000_000`.
// - `bankruptcyAdjustmentPpm < 1_000_000`.
// - `spreadToMaintenanceMarginRatioPpm == 0.
// - `minPositionNotionalLiquidated` < 0`.
// - `maxPositionPortionLiquidatedPpm == 0 || maxPositionPortionLiquidatedPpm > 1_000_000`.
// - `maxNotionalLiquidated == 0 || maxNotionalLiquidated < 0`.
// - `maxQuantumsInsuranceLost == 0 || maxNotionalLiquidated < 0`.

func (lc *LiquidationsConfig) Validate() error {
	// Validate the BankruptcyAdjustmentPpm.
	bankruptcyAdjustmentPpm := lc.FillablePriceConfig.BankruptcyAdjustmentPpm
	if bankruptcyAdjustmentPpm < lib.OneMillion {
		return errorsmod.Wrapf(
			ErrInvalidLiquidationsConfig,
			"%v is not a valid BankruptcyAdjustmentPpm",
			bankruptcyAdjustmentPpm,
		)
	}

	// Validate the SpreadToMaintenanceMarginRatioPpm.
	spreadToMaintenanceMarginRatioPpm := lc.FillablePriceConfig.SpreadToMaintenanceMarginRatioPpm
	if spreadToMaintenanceMarginRatioPpm == 0 {
		return errorsmod.Wrapf(
			ErrInvalidLiquidationsConfig,
			"%v is not a valid SpreadToMaintenanceMarginRatioPpm",
			spreadToMaintenanceMarginRatioPpm,
		)
	}

	// Validate the MaxLiquidationFeePpm.
	if lc.MaxLiquidationFeePpm == 0 || lc.MaxLiquidationFeePpm > lib.OneMillion {
		return errorsmod.Wrapf(
			ErrInvalidLiquidationsConfig,
			"%v is not a valid MaxLiquidationFeePpm",
			lc.MaxLiquidationFeePpm,
		)
	}

	// Validate MinPositionNotionalLiquidated
	minPositionNotionalLiquidated := lc.PositionBlockLimits.MinPositionNotionalLiquidated
	if minPositionNotionalLiquidated.Cmp(dtypes.NewInt(0)) == -1 {
		return errorsmod.Wrapf(
			ErrInvalidLiquidationsConfig,
			"%v is not a valid MinPositionNotionalLiquidated",
			minPositionNotionalLiquidated,
		)
	}

	// Validate the MaxPositionPortionLiquidatedPpm.
	maxPositionPortionLiquidatedPpm := lc.PositionBlockLimits.MaxPositionPortionLiquidatedPpm
	if maxPositionPortionLiquidatedPpm == 0 || maxPositionPortionLiquidatedPpm > lib.OneMillion {
		return errorsmod.Wrapf(
			ErrInvalidLiquidationsConfig,
			"%v is not a valid MaxPositionPortionLiquidatedPpm",
			maxPositionPortionLiquidatedPpm,
		)
	}

	// Validate the MaxNotionalLiquidated.
	maxNotionalLiquidated := lc.SubaccountBlockLimits.MaxNotionalLiquidated
	if maxNotionalLiquidated.Cmp(dtypes.NewInt(0)) == 0 || maxNotionalLiquidated.Cmp(dtypes.NewInt(0)) == -1 {
		return errorsmod.Wrapf(
			ErrInvalidLiquidationsConfig,
			"%v is not a valid MaxNotionalLiquidated",
			maxNotionalLiquidated,
		)
	}

	// Validate the MaxQuantumsInsuranceLost.
	maxQuantumsInsuranceLost := lc.SubaccountBlockLimits.MaxQuantumsInsuranceLost
	if maxQuantumsInsuranceLost.Cmp(dtypes.NewInt(0)) == 0 || maxQuantumsInsuranceLost.Cmp(dtypes.NewInt(0)) == -1 {
		return errorsmod.Wrapf(
			ErrInvalidLiquidationsConfig,
			"%v is not a valid MaxQuantumsInsuranceLost",
			maxQuantumsInsuranceLost,
		)
	}

	return nil
}
