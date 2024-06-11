package types

import "github.com/StreamFinance-Protocol/stream-chain/protocol/lib"

const (
	// UusdcDenom is the precomputed denom for IBC Micro USDC.
	UusdcDenom         = "ibc/8E27BA2D5493AF5636760E354E46004562C46AB7EC0CC4C1CA14E9E20E2545B5"
	UusdcDenomExponent = -6

	// WethWeiDenom is the precomputed denom for IBC Wrapped Ethereum
	WethWeiDenom = "ibc/EA1D43981D5C9A1C4AAEA9C23BB1D4FA126BA9BC7020A25E0AE4AA841EA25DC5"
	WethWeiDenomExponent = -18
)

var (
	AssetUsdc Asset = Asset{
		Id:               0,
		Symbol:           "USDC",
		DenomExponent:    UusdcDenomExponent,
		Denom:            UusdcDenom,
		HasMarket:        false,
		AtomicResolution: lib.QuoteCurrencyAtomicResolution,
	}

	AssetEth Asset = Asset{
		Id:               1,
		Symbol:           "ETH",
		DenomExponent:    WethWeiDenomExponent,
		Denom:            WethWeiDenom,
		HasMarket:        true,
		MarketId:		  lib.EthMarketId,
		AtomicResolution: lib.EthAtomicResolution, // TODO: This needs to be price adjusted. 
	}
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Assets: []Asset{
			AssetUsdc,
			AssetEth,
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Genesis state should contain at least one asset.
	if len(gs.Assets) == 0 {
		return ErrNoAssetInGenesis
	}

	// The first asset should always be USDC.
	if gs.Assets[0] != AssetUsdc {
		return ErrUsdcMustBeAssetZero
	}

	// The second asset should always be ETH
	if gs.Assets[1] != AssetEth {
		return ErrEthMustBeAssetOne
	}

	// Provided assets should not contain duplicated asset ids, and denoms.
	// Asset ids should be sequential.
	// MarketId should be 0 if HasMarket is false.
	assetIdSet := make(map[uint32]struct{})
	denomSet := make(map[string]struct{})
	expectedId := uint32(0)

	for _, asset := range gs.Assets {
		if _, exists := assetIdSet[asset.Id]; exists {
			return ErrAssetIdAlreadyExists
		}
		if _, exists := denomSet[asset.Denom]; exists {
			return ErrAssetDenomAlreadyExists
		}
		if asset.Id != expectedId {
			return ErrGapFoundInAssetId
		}
		if !asset.HasMarket && asset.MarketId > 0 {
			return ErrInvalidMarketId
		}
		assetIdSet[asset.Id] = struct{}{}
		denomSet[asset.Denom] = struct{}{}
		expectedId = expectedId + 1
	}
	return nil
}
