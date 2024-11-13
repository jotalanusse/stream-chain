package types

import "fmt"

func (multiCollateralAssetsArray MultiCollateralAssetsArray) Validate() error {

	// Validate multi collateral assets
	// 1. IDs are sequential and unique
	// 2. Contains tdai
	if len(multiCollateralAssetsArray.MultiCollateralAssets) == 0 {
		return fmt.Errorf("no supported multi collateral assets")
	}
	assetId := uint32(0)
	for j, id := range multiCollateralAssetsArray.MultiCollateralAssets {
		if j == 0 {
			if id != 0 {
				return fmt.Errorf("tDai is a required multi collateral asset")
			}
		} else {
			if id <= assetId {
				return fmt.Errorf("multi collateral assets must be unique and sequential")
			}
			assetId = id
		}
	}

	return nil
}
