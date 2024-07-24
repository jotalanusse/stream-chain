package price_writer

import (
	"cosmossdk.io/log"

	"math/big"

	aggregator "github.com/StreamFinance-Protocol/stream-chain/protocol/app/ve/aggregator"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/app/ve/codec"
	pricecache "github.com/StreamFinance-Protocol/stream-chain/protocol/daemons/pricefeed/pricecache"
	pricestypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/prices/types"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PriceWriter is an interface that defines the methods required to aggregate and apply prices from VE's
type PriceApplier struct {
	// used to aggregate votes into final prices
	voteAggregator aggregator.VoteAggregator

	// used to write prices to state.
	pricesKeeper PriceApplierPricesKeeper

	// used to get the mid price of clob markets
	clobKeeper PriceApplierClobKeeper

	// used to get funding rates
	perpetualKeeper PriceApplierPerpetualsKeeper

	// finalPriceCache is the cache that stores the final prices
	finalPriceCache pricecache.PriceCache

	// logger
	logger log.Logger

	// codecs
	voteExtensionCodec  codec.VoteExtensionCodec
	extendedCommitCodec codec.ExtendedCommitCodec
}

func NewPriceApplier(
	logger log.Logger,
	voteAggregator aggregator.VoteAggregator,
	pricesKeeper PriceApplierPricesKeeper,
	clobKeeper PriceApplierClobKeeper,
	perpetualKeeper PriceApplierPerpetualsKeeper,
	voteExtensionCodec codec.VoteExtensionCodec,
	extendedCommitCodec codec.ExtendedCommitCodec,
) *PriceApplier {
	return &PriceApplier{
		voteAggregator:      voteAggregator,
		pricesKeeper:        pricesKeeper,
		clobKeeper:          clobKeeper,
		perpetualKeeper:     perpetualKeeper,
		logger:              logger,
		voteExtensionCodec:  voteExtensionCodec,
		extendedCommitCodec: extendedCommitCodec,
	}
}

func (pa *PriceApplier) ApplyPricesFromVE(
	ctx sdk.Context,
	request *abci.RequestFinalizeBlock,
) error {
	prices, isCached, err := pa.writePricesToStore(ctx, request)
	if err != nil {
		return err
	}

	if !isCached && prices != nil {
		pa.writePricesToCache(ctx, request.DecidedLastCommit.Round, prices)
	}

	return nil
}

func (pa *PriceApplier) writePricesToStore(
	ctx sdk.Context,
	request *abci.RequestFinalizeBlock,
) (prices map[string]*big.Int, isCached bool, err error) {
	if pa.finalPriceCache.HasValidPrices(ctx.BlockHeight(), request.DecidedLastCommit.Round) {
		err := pa.writePricesToStoreFromCache(ctx)
		return nil, true, err
	} else {
		prices, err := pa.getPricesAndAggregateFromVE(ctx, request)
		if err != nil {
			return nil, false, err
		}

		pa.fallbackWritePricesToStore(ctx, prices)
		return prices, false, nil
	}
}

func (pa *PriceApplier) getPricesAndAggregateFromVE(
	ctx sdk.Context,
	request *abci.RequestFinalizeBlock,
) (map[string]*big.Int, error) {
	votes, err := aggregator.GetDaemonVotesFromBlock(request.Txs, pa.voteExtensionCodec, pa.extendedCommitCodec)

	if err != nil {
		pa.logger.Error(
			"failed to get extended commit info from proposal",
			"height", request.Height,
			"num_txs", len(request.Txs),
			"err", err,
		)

		return nil, err
	}

	clobMetadata := pa.getClobMidPrices(ctx)
	lastFundingIndexes := pa.getLastFundingRates(ctx)
	smoothedPrices := pa.getSmoothedPrices(ctx)

	prices, err := pa.voteAggregator.AggregateDaemonVEIntoFinalPrices(
		ctx,
		votes,
		clobMetadata,
		lastFundingIndexes,
		smoothedPrices,
	)

	if err != nil {
		pa.logger.Error(
			"failed to aggregate prices",
			"height", request.Height,
			"err", err,
		)

		return nil, err
	}

	return prices, nil
}

func (pa *PriceApplier) getClobMidPrices(
	ctx sdk.Context,
) map[string]*big.Int {
	midPrices := make(map[string]*big.Int)
	metadata := pa.clobKeeper.GetClobMetadata(ctx)
	for pair, data := range metadata {
		market, exists := pa.pricesKeeper.GetMarketParam(ctx, pair.ToUint32())
		if !exists || data.MidPrice == 0 {
			continue
		}
		midPrices[market.Pair] = data.MidPrice.ToBigInt()
	}

	return map[string]*big.Int{}
}

func (pa *PriceApplier) getSmoothedPrices(
	ctx sdk.Context,
) map[string]*big.Int {
	smoothedPrices := make(map[string]*big.Int)
	marketParams := pa.pricesKeeper.GetAllMarketParams(ctx)
	for _, market := range marketParams {
		marketPair := market.Pair
		smoothedPrice, exists := pa.pricesKeeper.GetSmoothedPrice(market.Id)
		if !exists || smoothedPrice == 0 {
			continue
		}
		smoothedPrices[marketPair] = new(big.Int).SetUint64(smoothedPrice)

	}
	return smoothedPrices
}

func (pa *PriceApplier) getLastFundingRates(
	ctx sdk.Context,
) map[string]*big.Int {
	fundingRates := make(map[string]*big.Int)
	marketParams := pa.pricesKeeper.GetAllMarketParams(ctx)
	for _, market := range marketParams {
		marketPair := market.Pair
		perpetual, err := pa.perpetualKeeper.GetPerpetual(ctx, market.Id)
		if err != nil {
			continue
		}
		fundingRates[marketPair] = perpetual.LastFundingRate.BigInt()
	}

	return fundingRates
}

func (pa *PriceApplier) GetCachedPrices() pricestypes.MarketPriceUpdates {
	return pa.finalPriceCache.GetPriceUpdates()
}

func (pa *PriceApplier) writePricesToCache(
	ctx sdk.Context,
	round int32,
	prices map[string]*big.Int,
) {
	marketParams := pa.pricesKeeper.GetAllMarketParams(ctx)
	var pricesToCache pricestypes.MarketPriceUpdates
	for _, market := range marketParams {
		shouldWritePrice, price := pa.shouldWritePriceToStore(ctx, prices, market)
		if !shouldWritePrice {
			continue
		}

		newPrice := pricestypes.MarketPriceUpdates_MarketPriceUpdate{
			MarketId: market.Id,
			Price:    price.Uint64(),
		}

		pricesToCache.MarketPriceUpdates = append(pricesToCache.MarketPriceUpdates, &newPrice)
	}
	pa.finalPriceCache.SetPriceUpdates(ctx, pricesToCache, round)
}

func (pa *PriceApplier) writePricesToStoreFromCache(ctx sdk.Context) error {
	pricesFromCache := pa.finalPriceCache.GetPriceUpdates()
	for _, price := range pricesFromCache.MarketPriceUpdates {
		if err := pa.pricesKeeper.UpdateMarketPrice(ctx, price); err != nil {
			pa.logger.Error(
				"failed to set price for currency pair",
				"market_id", price.MarketId,
				"err", err,
			)

			return err
		}

		pa.logger.Info(
			"set price for currency pair",
			"market_id", price.MarketId,
			"quote_price", price.Price,
		)
	}
	return nil
}

func (pa *PriceApplier) fallbackWritePricesToStore(ctx sdk.Context, prices map[string]*big.Int) error {
	marketParams := pa.pricesKeeper.GetAllMarketParams(ctx)
	for _, market := range marketParams {
		pair := market.Pair
		shouldWritePrice, price := pa.shouldWritePriceToStore(ctx, prices, market)
		if !shouldWritePrice {
			continue
		}

		newPrice := pricestypes.MarketPriceUpdates_MarketPriceUpdate{
			MarketId: market.Id,
			Price:    price.Uint64(),
		}

		if err := pa.pricesKeeper.UpdateMarketPrice(ctx, &newPrice); err != nil {
			pa.logger.Error(
				"failed to set price for currency pair",
				"currency_pair", pair,
				"err", err,
			)

			return err
		}

		pa.logger.Info(
			"set price for currency pair",
			"currency_pair", pair,
			"quote_price", newPrice.Price,
		)
	}
	return nil
}

func (pa *PriceApplier) shouldWritePriceToStore(
	ctx sdk.Context,
	prices map[string]*big.Int,
	marketToUpdate pricestypes.MarketParam,
) (bool, *big.Int) {
	marketPairToUpdate := marketToUpdate.Pair
	price, ok := prices[marketPairToUpdate]
	if !ok || price == nil {
		pa.logger.Debug(
			"no price for currency pair",
			"currency_pair", marketPairToUpdate,
		)
		return false, nil
	}

	if price.Sign() == -1 {
		pa.logger.Error(
			"price is negative",
			"currency_pair", marketPairToUpdate,
			"price", price.String(),
		)

		return false, nil
	}
	priceUpdate := pricestypes.MarketPriceUpdates{
		MarketPriceUpdates: []*pricestypes.MarketPriceUpdates_MarketPriceUpdate{
			{
				MarketId: marketToUpdate.Id,
				Price:    price.Uint64(),
			},
		},
	}

	if pa.pricesKeeper.PerformStatefulPriceUpdateValidation(ctx, &priceUpdate) != nil {
		pa.logger.Error(
			"price update validation failed",
			"currency_pair", marketPairToUpdate,
			"price", price.String(),
		)

		return false, nil
	}

	return true, price
}
