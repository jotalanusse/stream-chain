package types

import (
	errorsmod "cosmossdk.io/errors"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
)

func (collateralPool CollateralPool) Validate() error {
	if collateralPool.MaxCumulativeInsuranceFundDeltaPerBlock == 0 {
		return errorsmod.Wrap(ErrMaxCumulativeInsuranceFundDeltaPerBlockZero, lib.UintToString(collateralPool.MaxCumulativeInsuranceFundDeltaPerBlock))
	}

	if len(collateralPool.MarketMultiCollateralAssets.MultiCollateralAssets) == 0 {
		return errorsmod.Wrap(ErrMarketMultiCollateralAssetsEmpty, "")
	}

	quoteAssetFound := false
	for _, asset := range collateralPool.MarketMultiCollateralAssets.MultiCollateralAssets {
		if asset == collateralPool.QuoteAssetId {
			quoteAssetFound = true
			break
		}
	}

	if !quoteAssetFound {
		return errorsmod.Wrap(ErrIsolatedMarketMultiCollateralAssetDoesNotContainQuoteAsset, "")
	}

	return nil
}
