package ve

import (
	"fmt"
	"math/big"
	"sort"

	"cosmossdk.io/log"
	codec "github.com/StreamFinance-Protocol/stream-chain/protocol/app/ve/codec"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/app/ve/types"
	veutils "github.com/StreamFinance-Protocol/stream-chain/protocol/app/ve/utils"
	clobtypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/clob/types"
	pricetypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/prices/types"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type VoteExtensionHandler struct {
	logger log.Logger

	// encoding and decoding vote extensions
	voteCodec codec.VoteExtensionCodec

	// fetching valid price updates and current markets
	pricesKeeper PreBlockExecPricesKeeper

	// fetching last funding rates for price calc
	perpetualsKeeper ExtendVotePerpetualsKeeper

	// fetching mid price for price calc
	clobKeeper ExtendVoteClobKeeper

	// writing prices to the prices module store
	priceApplier VEPriceApplier
}

var (
	acceptResponse = &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_ACCEPT}
	rejectResponse = &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_REJECT}
)

func NewVoteExtensionHandler(
	logger log.Logger,
	voteCodec codec.VoteExtensionCodec,
	pricesKeeper PreBlockExecPricesKeeper,
	perpetualsKeeper ExtendVotePerpetualsKeeper,
	clobKeeper ExtendVoteClobKeeper,
	priceApplier VEPriceApplier,
) *VoteExtensionHandler {
	return &VoteExtensionHandler{
		logger:           logger,
		voteCodec:        voteCodec,
		pricesKeeper:     pricesKeeper,
		perpetualsKeeper: perpetualsKeeper,
		clobKeeper:       clobKeeper,
		priceApplier:     priceApplier,
	}
}

// Returns a handler that extends pre-commit votes with the current
// prices pulled from the perpetually running price daemon
// In the case of an error, the handler will return an empty vote extension
// ensuring liveness in the case of a price daemon failure
func (h *VoteExtensionHandler) ExtendVoteHandler() sdk.ExtendVoteHandler {
	return func(ctx sdk.Context, request *abci.RequestExtendVote) (resp *abci.ResponseExtendVote, err error) {
		defer func() {
			if recovery := recover(); recovery != nil {
				h.logger.Error(
					"recovered from panic in ExtendVoteHandler",
					"err", recovery,
				)
				resp = &abci.ResponseExtendVote{VoteExtension: []byte{}}
				err = ErrPanic{fmt.Errorf("%v", recovery)}
			}
		}()

		if request == nil {
			err = fmt.Errorf("nil request for extend vote")
			return nil, err
		}

		reqFinalizeBlock := &abci.RequestFinalizeBlock{
			Txs:    request.Txs,
			Height: request.Height,
			DecidedLastCommit: abci.CommitInfo{
				Round: request.ProposedLastCommit.Round,
				Votes: []abci.VoteInfo{},
			},
		}

		// apply prices from prev block to ensure that the prices are up to date
		if err := h.priceApplier.ApplyPricesFromVE(ctx, reqFinalizeBlock); err != nil {
			h.logger.Error(
				"failed to aggregate oracle votes",
				"height", request.Height,
				"err", err,
			)
			err = PreBlockError{err}

			return &abci.ResponseExtendVote{VoteExtension: []byte{}}, err
		}

		veBytes, err := h.GetVEBytesFromCurrPrices(ctx)
		if err != nil {
			h.logger.Error(
				"failed to get vote extension bytes from current prices",
				"height", request.Height,
				"err", err,
			)
			return &abci.ResponseExtendVote{VoteExtension: []byte{}}, err
		}

		return &abci.ResponseExtendVote{VoteExtension: veBytes}, nil
	}
}

func (h *VoteExtensionHandler) VerifyVoteExtensionHandler() sdk.VerifyVoteExtensionHandler {
	return func(
		ctx sdk.Context,
		req *abci.RequestVerifyVoteExtension,
	) (resp *abci.ResponseVerifyVoteExtension, err error) {

		defer func() {
			if recovery := recover(); recovery != nil {
				h.logger.Error(
					"recovered from panic in VerifyVoteExtensionHandler",
					"err", recovery,
				)
				resp = rejectResponse
				err = ErrPanic{fmt.Errorf("%v", recovery)}

			}
		}()

		if req == nil {
			err = fmt.Errorf("nil request for verify vote")
			return nil, err
		}

		if len(req.VoteExtension) == 0 {
			h.logger.Info(
				"empty vote extension",
				"height", req.Height,
			)

			return acceptResponse, nil
		}

		if err := ValidateVEMarketsAndPrices(
			ctx,
			h.pricesKeeper,
			req.VoteExtension,
			h.voteCodec,
		); err != nil {
			h.logger.Error(
				"failed to decode and validate vote extension",
				"height", req.Height,
				"err", err,
			)
			return rejectResponse, err
		}

		return acceptResponse, nil
	}
}

func (h *VoteExtensionHandler) GetVEBytesFromCurrPrices(ctx sdk.Context) ([]byte, error) {
	priceUpdates := h.getCurrentPrices(ctx)

	if len(priceUpdates.MarketPriceUpdates) == 0 {
		return nil, fmt.Errorf("no valid median prices")
	}

	// turn prices from daemon into a VE
	voteExt, err := h.transformDaemonPricesToVE(priceUpdates.MarketPriceUpdates)
	if err != nil {
		return nil, err
	}

	veBytes, err := h.voteCodec.Encode(voteExt)
	if err != nil {
		return nil, err
	}

	return veBytes, nil
}

func (h *VoteExtensionHandler) transformDaemonPricesToVE(
	priceupdates []*pricetypes.MarketPriceUpdates_MarketPriceUpdate,
) (types.DaemonVoteExtension, error) {
	vePrices := make(map[uint32][]byte)

	for _, priceUpdate := range priceupdates {
		// check if the marketId is valid
		encodedPrice, err := h.GetEncodedPriceFromPriceUpdate(priceUpdate)
		if err != nil {
			continue
		}
		marketId := priceUpdate.GetMarketId()
		vePrices[marketId] = encodedPrice
	}

	return types.DaemonVoteExtension{
		Prices: vePrices,
	}, nil
}

func (h *VoteExtensionHandler) GetEncodedPriceFromPriceUpdate(
	priceUpdate *pricetypes.MarketPriceUpdates_MarketPriceUpdate,
) ([]byte, error) {
	price := new(big.Int).SetUint64(priceUpdate.GetPrice())

	encodedPrice, err := veutils.GetVEEncodedPrice(price)
	if err != nil {
		return nil, err
	}

	return encodedPrice, nil
}

func (h *VoteExtensionHandler) getCurrentPrices(
	ctx sdk.Context,
) *pricetypes.MarketPriceUpdates {
	indexPrices := h.pricesKeeper.GetValidMarketPriceUpdates(ctx)

	for i, market := range indexPrices.MarketPriceUpdates {

		clobMidPrice := h.getClobMidPrice(ctx, market.MarketId)
		if clobMidPrice == nil {
			continue
		}
		smoothedPrice := h.getSmoothedPrice(market.MarketId)
		if smoothedPrice == nil {
			continue
		}
		lastFundingRate := h.getLastFundingRate(ctx, market.MarketId)
		if lastFundingRate == nil {
			continue
		}

		medianPrice := h.getMedianPrice(
			new(big.Int).SetUint64(market.Price),
			clobMidPrice,
			smoothedPrice,
			lastFundingRate,
		)

		indexPrices.MarketPriceUpdates[i].Price = medianPrice.Uint64()
	}

	return indexPrices
}

func (h *VoteExtensionHandler) getMedianPrice(
	indexPrice *big.Int,
	clobMidPrice *big.Int,
	smoothedPrice *big.Int,
	lastFundingRate *big.Int,
) *big.Int {
	adjustedFundingRate := new(big.Int).Add(lastFundingRate, big.NewInt(1))
	fundingWeightedPrice := new(big.Int).Mul(indexPrice, adjustedFundingRate)

	prices := []*big.Int{clobMidPrice, smoothedPrice, fundingWeightedPrice}
	sort.Slice(prices, func(i, j int) bool {
		return prices[i].Cmp(prices[j]) < 0
	})

	return prices[1]
}

func (h *VoteExtensionHandler) getClobMidPrice(
	ctx sdk.Context,
	marketId uint32,
) *big.Int {
	clobPair, found := h.clobKeeper.GetClobPair(ctx, clobtypes.ClobPairId(marketId))

	if !found {
		return nil
	}

	clobMetadata := h.clobKeeper.GetSingleMarketClobMetadata(ctx, clobPair)

	return clobMetadata.MidPrice.ToBigInt()
}

func (h *VoteExtensionHandler) getSmoothedPrice(
	marketId uint32,
) *big.Int {
	smoothedPrice, exists := h.pricesKeeper.GetSmoothedPrice(marketId)
	if !exists || smoothedPrice == 0 {
		return nil
	}

	return new(big.Int).SetUint64(smoothedPrice)
}

func (h *VoteExtensionHandler) getLastFundingRate(
	ctx sdk.Context,
	marketId uint32,
) *big.Int {
	perpetual, err := h.perpetualsKeeper.GetPerpetual(ctx, marketId)
	if err != nil {
		return nil
	}

	return perpetual.LastFundingRate.BigInt()
}
