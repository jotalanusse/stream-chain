package keeper

import (
	"fmt"
	"math/big"

	veaggregator "github.com/StreamFinance-Protocol/stream-chain/protocol/app/ve/aggregator"
	vecodec "github.com/StreamFinance-Protocol/stream-chain/protocol/app/ve/codec"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
	pricestypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/prices/types"
	satypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/subaccounts/types"
	abcicomet "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetLiquidatableSubaccountIds(ctx sdk.Context, extendedCommitInfo *abcicomet.ExtendedCommitInfo) []satypes.SubaccountId {

	subaccounts := k.subaccountsKeeper.GetAllSubaccount(ctx)

	perps := k.perpetualsKeeper.GetAllPerpetuals(ctx)

	liquidityTiers := k.perpetualsKeeper.GetAllLiquidityTiers(ctx)

	branchedCtx, ctxErr := ctx.CacheContext()

	// from cometbft so is either nil or is valid and > 2/3
	if (extendedCommitInfo != &abcicomet.ExtendedCommitInfo{}) && ctxErr == nil {
		votes, err := getDaemonVotesFromExtendedCommitInfo(extendedCommitInfo)
		if err == nil {
			prices, err := k.Aggregator.AggregateDaemonVEIntoFinalPrices(ctx, votes)
			if err == nil {
				err = k.writePricesToStore(branchedCtx, prices)
			}
		}
	}

	prices := []pricestypes.MarketPrice{}

	if ctxErr == nil {
		prices = k.pricesKeeper.GetAllMarketPrices(branchedCtx)

	} else {
		prices = k.pricesKeeper.GetAllMarketPrices(ctx)
	}

	marketPricesMap := lib.UniqueSliceToMap(prices, func(m pricestypes.MarketPrice) uint32 {
		return m.Id
	})

	// check what is liquidatable
}

func (k Keeper) writePricesToStore(
	ctx sdk.Context,
	prices map[string]*big.Int,
) error {
	marketParams := k.pricesKeeper.GetAllMarketParams(ctx)
	for _, market := range marketParams {
		pair := market.Pair
		price, ok := prices[pair]
		if !ok {
			continue
		}
		shouldWritePrice, price := k.shouldWritePriceToStore(ctx, price, market.Id)
		if !shouldWritePrice {
			continue
		}

		newPrice := pricestypes.MarketPriceUpdates_MarketPriceUpdate{
			MarketId: market.Id,
			Price:    price.Uint64(),
		}

		if err := k.pricesKeeper.UpdateMarketPrice(ctx, &newPrice); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) shouldWritePriceToStore(ctx sdk.Context, price *big.Int, marketId uint32) (bool, *big.Int) {

	if price.Sign() == -1 {
		return false, nil
	}
	priceUpdate := pricestypes.MarketPriceUpdates{
		MarketPriceUpdates: []*pricestypes.MarketPriceUpdates_MarketPriceUpdate{
			{
				MarketId: marketId,
				Price:    price.Uint64(),
			},
		},
	}

	if k.pricesKeeper.PerformStatefulPriceUpdateValidation(ctx, &priceUpdate) != nil {
		return false, nil
	}

	return true, price
}

func getDaemonVotesFromExtendedCommitInfo(extendedCommitInfo *abcicomet.ExtendedCommitInfo) ([]veaggregator.Vote, error) {
	veCodec := vecodec.NewDefaultVoteExtensionCodec()

	votes := make([]veaggregator.Vote, len(extendedCommitInfo.Votes))
	for i, voteInfo := range extendedCommitInfo.Votes {
		voteExtension, err := veCodec.Decode(voteInfo.VoteExtension)
		if err != nil {
			return nil, fmt.Errorf("error decoding vote-extension: %w", err)
		}

		votes[i] = veaggregator.Vote{
			ConsAddress:         voteInfo.Validator.Address,
			DaemonVoteExtension: voteExtension,
		}
	}
}
