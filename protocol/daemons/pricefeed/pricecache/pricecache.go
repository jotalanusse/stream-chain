package pricecache

import (
	"math/big"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// this cache is used to set prices from vote extensions in processProposal
// which are fetched in ExtendVoteHandler and PreBlocker. This is to avoid
// redundant computation on calculating stake weighthed median prices in VEs
type PriceCache struct {
	priceUpdates  PriceUpdates
	height        int64
	round         int32
	consAddresses map[string]struct{}
	mu            sync.RWMutex
}

type PriceUpdate struct {
	MarketId  uint32
	SpotPrice *big.Int
	PnlPrice  *big.Int
}

type PriceUpdates []PriceUpdate

func (pc *PriceCache) SetPriceUpdates(
	ctx sdk.Context,
	updates PriceUpdates,
	round int32,
) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.priceUpdates = updates
	pc.height = ctx.BlockHeight()
	pc.round = round
}

func (pc *PriceCache) SetConsAddresses(
	height int64,
	round int32,
	consAddresses map[string]struct{},
) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	if pc.height != height || pc.round != round {
		return
	}
	pc.consAddresses = consAddresses
}

func (pc *PriceCache) GetPriceUpdates() PriceUpdates {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	return pc.priceUpdates
}

func (pc *PriceCache) GetConsAddresses() map[string]struct{} {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	return pc.consAddresses
}

func (pc *PriceCache) GetHeight() int64 {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	return pc.height
}

func (pc *PriceCache) GetRound() int32 {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	return pc.round
}

func (pc *PriceCache) HasValidPrices(currBlock int64, round int32) bool {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	return (pc.height == currBlock && pc.round == round)
}
