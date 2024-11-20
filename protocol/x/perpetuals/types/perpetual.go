package types

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
	"github.com/pkg/errors"
)

func (p *Perpetual) GetId() uint32 {
	return p.Params.Id
}

// Stateless validation on Perpetual params.
func (p *PerpetualParams) Validate() error {

	// Validate `ticker`.
	if len(p.Ticker) == 0 {
		return errors.WithStack(ErrTickerEmptyString)
	}

	// Validate `defaultFundingPpm`
	defaultFundingPpm := lib.AbsInt32(p.DefaultFundingPpm)
	if defaultFundingPpm > MaxDefaultFundingPpmAbs {
		return errorsmod.Wrap(
			ErrDefaultFundingPpmMagnitudeExceedsMax,
			lib.IntToString(p.DefaultFundingPpm))
	}

	if p.IsolatedMarketMaxCumulativeInsuranceFundDeltaPerBlock == 0 {
		return errorsmod.Wrap(
			ErrIsolatedMarketMaxCumulativeInsuranceFundDeltaPerBlockZero,
			lib.UintToString(p.IsolatedMarketMaxCumulativeInsuranceFundDeltaPerBlock),
		)
	}

	if len(p.IsolatedMarketMultiCollateralAssets.MultiCollateralAssets) == 0 {
		return errorsmod.Wrap(
			ErrIsolatedMarketMultiCollateralAssetsEmpty,
			"In validate perpetual params",
		)
	} else {
		// Check that MultiCollateralAssets contains the quote asset
		containsQuoteAsset := false
		for _, asset := range p.IsolatedMarketMultiCollateralAssets.MultiCollateralAssets {
			if asset == p.QuoteAssetId {
				containsQuoteAsset = true
				break
			}
		}
		if !containsQuoteAsset {
			return errorsmod.Wrap(
				ErrIsolatedMarketMultiCollateralAssetDoesNotContainQuoteAsset,
				"MultiCollateralAssets does not contain quote asset",
			)
		}
	}

	return nil
}
