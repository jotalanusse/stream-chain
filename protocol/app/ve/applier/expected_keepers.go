package price_writer

import (
	clobKeeper "github.com/StreamFinance-Protocol/stream-chain/protocol/x/clob/keeper"
	clobtypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/clob/types"
	perptypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/perpetuals/types"
	pricestypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/prices/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PriceApplierPricesKeeper interface {
	PerformStatefulPriceUpdateValidation(
		ctx sdk.Context,
		marketPriceUpdates *pricestypes.MarketPriceUpdates,
	) error

	GetAllMarketParams(ctx sdk.Context) []pricestypes.MarketParam

	UpdateMarketPrice(
		ctx sdk.Context,
		update *pricestypes.MarketPriceUpdates_MarketPriceUpdate,
	) error
	GetMarketParam(
		ctx sdk.Context,
		id uint32,
	) (
		market pricestypes.MarketParam,
		exists bool,
	)

	GetSmoothedPrice(
		markedId uint32,
	) (uint64, bool)
}

type PriceApplierClobKeeper interface {
	GetClobMetadata(
		ctx sdk.Context,
	) (
		clobMetadata map[clobtypes.ClobPairId]clobKeeper.ClobMetadata,
	)
}

type PriceApplierPerpetualsKeeper interface {
	GetPerpetual(
		ctx sdk.Context,
		id uint32,
	) (val perptypes.Perpetual, err error)
}
