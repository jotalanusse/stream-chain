package types

import (
	errorsmod "cosmossdk.io/errors"
)

// ValidateFromParam checks that the MarketPrice is valid and that it corresponds to the given MarketParam.
func (mp *MarketPrice) ValidateFromParam(marketParam MarketParam) error {
	if marketParam.Id != mp.Id {
		return errorsmod.Wrapf(
			ErrInvalidInput,
			"market param id %d does not match market price id %d",
			marketParam.Id,
			mp.Id,
		)
	}
	if marketParam.Exponent != mp.Exponent {
		return errorsmod.Wrapf(
			ErrInvalidInput,
			"market param %d exponent %d does not match market price %d exponent %d",
			marketParam.Id,
			marketParam.Exponent,
			mp.Id,
			mp.Exponent,
		)
	}
	return nil
}

func NewMarketPriceUpdate(id uint32, spotPrice uint64, pnlPrice uint64) *MarketPriceUpdate {
	return &MarketPriceUpdate{
		MarketId:  id,
		SpotPrice: spotPrice,
		PnlPrice:  pnlPrice,
	}
}

func NewMarketSpotPriceUpdate(id uint32, spotPrice uint64) *MarketSpotPriceUpdate {
	return &MarketSpotPriceUpdate{
		MarketId:  id,
		SpotPrice: spotPrice,
	}
}
