package keeper

import (
	aggregator "github.com/StreamFinance-Protocol/stream-chain/protocol/app/ve/aggregator"
	voteweighted "github.com/StreamFinance-Protocol/stream-chain/protocol/app/ve/math"
	abcicomet "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PriceApplierInterface interface {
	WritePricesToStoreAndMaybeCache(ctx sdk.Context, prices map[string]voteweighted.AggregatorPricePair, round int32, writeToCache bool) error
	VoteAggregator() aggregator.VoteAggregator
	CacheSeenExtendedVotes(
		ctx sdk.Context,
		req *abcicomet.RequestCommit,
	) error
}
