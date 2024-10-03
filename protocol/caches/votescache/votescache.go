package votescache

import (
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// this cache is used to set prices from vote extensions in processProposal
// which are fetched in ExtendVoteHandler and PreBlocker. This is to avoid
// redundant computation on calculating stake weighthed median prices in VEs
type VotesCache struct {
	height        int64
	consAddresses map[string]struct{}
	mu            sync.RWMutex
}

func NewVotesCache() *VotesCache {
	return &VotesCache{
		height:        0,
		consAddresses: make(map[string]struct{}),
	}
}

func (pc *VotesCache) SetSeenVotesInCache(
	ctx sdk.Context,
	consAddresses map[string]struct{},
) {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.height = ctx.BlockHeight()
	pc.consAddresses = consAddresses
}

func (pc *VotesCache) GetSeenVotesInCache() map[string]struct{} {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	return pc.consAddresses
}

func (pc *VotesCache) GetHeight() int64 {
	pc.mu.RLock()
	defer pc.mu.RUnlock()
	return pc.height
}
