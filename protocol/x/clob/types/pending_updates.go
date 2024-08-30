package types

import (
	"math/big"
	"sort"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
	assettypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/assets/types"
	satypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/subaccounts/types"
)

// pendingUpdates is a utility struct used for storing the working updates to all Subaccounts.
type PendingUpdates struct {
	subaccountAssetUpdates              map[satypes.SubaccountId]map[uint32]*big.Int
	subaccountPerpetualOrBalanceUpdates map[satypes.SubaccountId]map[uint32]*big.Int
	subaccountFee                       map[satypes.SubaccountId]*big.Int
}

// newPendingUpdates returns a new `pendingUpdates`.
func NewPendingUpdates() *PendingUpdates {
	return &PendingUpdates{
		subaccountAssetUpdates:              make(map[satypes.SubaccountId]map[uint32]*big.Int),
		subaccountPerpetualOrBalanceUpdates: make(map[satypes.SubaccountId]map[uint32]*big.Int),
		subaccountFee:                       make(map[satypes.SubaccountId]*big.Int),
	}
}

// ConvertToUpdates converts a `pendingUpdates` struct to a slice of Subaccount Updates.
func (p *PendingUpdates) ConvertToUpdates() []satypes.Update {
	// Build a slice of all subaccounts which were updated.
	var allSubaccounts = make([]satypes.SubaccountId, 0, len(p.subaccountAssetUpdates))
	for subaccountId := range p.subaccountAssetUpdates {
		allSubaccounts = append(allSubaccounts, subaccountId)
	}

	// Sort the subaccounts for determinism.
	sort.Sort(satypes.SortedSubaccountIds(allSubaccounts))

	// Iterate over all subaccounts to convert `*PendingUpdates` to `[]satypes.Update` for use in the
	// subaccounts module.
	var updates = make([]satypes.Update, 0, len(allSubaccounts))
	for _, subaccountId := range allSubaccounts {
		// Create an empty slice to store the asset updates for this subaccount.
		assetUpdates := make(
			[]satypes.AssetUpdate,
			0,
			len(p.subaccountAssetUpdates[subaccountId]),
		)

		pendingAssetUpdates := p.subaccountAssetUpdates[subaccountId]
		for assetId, bigQuantumsDelta := range pendingAssetUpdates {
			assetUpdate := satypes.AssetUpdate{
				AssetId:          assetId,
				BigQuantumsDelta: bigQuantumsDelta,
			}
			assetUpdates = append(assetUpdates, assetUpdate)
		}

		if _, exists := pendingAssetUpdates[assettypes.AssetUsdc.Id]; !exists {
			pendingAssetUpdates[assettypes.AssetUsdc.Id] = new(big.Int)
		}

		// Subtract quote balance delta with total fees paid by subaccount.
		pendingAssetUpdates[assettypes.AssetUsdc.Id].Sub(
			pendingAssetUpdates[assettypes.AssetUsdc.Id],
			p.subaccountFee[subaccountId],
		)

		// Panic if there is more than one asset updates since we only support
		// USDC asset at the moment.
		if len(assetUpdates) > 1 {
			panic(ErrAssetUpdateNotImplemented)
		}

		// Create an empty slice to store the perpetual updates for this subaccount.
		perpetualUpdates := make(
			[]satypes.PerpetualUpdate,
			0,
			len(p.subaccountPerpetualOrBalanceUpdates[subaccountId]),
		)

		for perpetualId, bigQuantumsDelta := range p.subaccountPerpetualOrBalanceUpdates[subaccountId] {
			perpetualUpdate := satypes.PerpetualUpdate{
				PerpetualId:      perpetualId,
				BigQuantumsDelta: bigQuantumsDelta,
			}
			perpetualUpdates = append(perpetualUpdates, perpetualUpdate)
		}

		// Sort the perpetualIds in ascending order for determinism.
		sort.Slice(perpetualUpdates, func(i, j int) bool {
			return perpetualUpdates[i].PerpetualId < perpetualUpdates[j].PerpetualId
		})

		// Create the update.
		update := satypes.Update{
			AssetUpdates:     assetUpdates,
			PerpetualUpdates: perpetualUpdates,
			SubaccountId:     subaccountId,
		}

		updates = append(updates, update)
	}

	return updates
}

// AddPerpetualFill adds a new fill to the PendingUpdate object, by
// updating quoteBalanceDelta, perpetualUpdate and fees paid or received by a subaccount.
func (p *PendingUpdates) AddPerpetualFill(
	subaccountId satypes.SubaccountId,
	perpetualId uint32,
	isBuy bool,
	feePpm int32,
	bigFillBaseQuantums *big.Int,
	bigFillQuoteQuantums *big.Int,
) {
	var quoteBalanceUpdate *big.Int
	var subaccountPerpetualOrBalanceUpdates map[uint32]*big.Int
	var perpetualOrBalanceUpdate *big.Int

	subaccountAssetUpdates, exists := p.subaccountAssetUpdates[subaccountId]
	if !exists {
		subaccountAssetUpdates = make(map[uint32]*big.Int)
		p.subaccountAssetUpdates[subaccountId] = subaccountAssetUpdates
	}
	quoteBalanceUpdate, exists = subaccountAssetUpdates[assettypes.AssetUsdc.Id]
	if !exists {
		quoteBalanceUpdate = big.NewInt(0)
		subaccountAssetUpdates[assettypes.AssetUsdc.Id] = quoteBalanceUpdate
	}

	subaccountPerpetualOrBalanceUpdates, exists = p.subaccountPerpetualOrBalanceUpdates[subaccountId]
	if !exists {
		subaccountPerpetualOrBalanceUpdates = make(map[uint32]*big.Int)
		p.subaccountPerpetualOrBalanceUpdates[subaccountId] = subaccountPerpetualOrBalanceUpdates
	}

	perpetualOrBalanceUpdate, exists = subaccountPerpetualOrBalanceUpdates[perpetualId]
	if !exists {
		perpetualOrBalanceUpdate = big.NewInt(0)
		subaccountPerpetualOrBalanceUpdates[perpetualId] = perpetualOrBalanceUpdate
	}

	p.handlePerpetualOrBalanceChanges(
		isBuy,
		quoteBalanceUpdate,
		perpetualOrBalanceUpdate,
		bigFillQuoteQuantums,
		bigFillBaseQuantums,
	)

	totalFee, exists := p.subaccountFee[subaccountId]
	if !exists {
		totalFee = big.NewInt(0)
	}
	// TODO(SCL) - figure out how to handle fees on the subaccount side here
	bigFeeQuoteQuantums := lib.BigIntMulSignedPpm(bigFillQuoteQuantums, feePpm, true)

	totalFee.Add(
		totalFee,
		bigFeeQuoteQuantums,
	)
	p.subaccountFee[subaccountId] = totalFee
}

func (p *PendingUpdates) handlePerpetualOrBalanceChanges(
	isBuy bool,
	quoteBalanceUpdate *big.Int,
	perpetualOrBalanceUpdate *big.Int,
	bigFillQuoteQuantums *big.Int,
	bigFillBaseQuantums *big.Int,
) {
	if isBuy {
		quoteBalanceUpdate.Sub(
			quoteBalanceUpdate,
			bigFillQuoteQuantums,
		)

		perpetualOrBalanceUpdate.Add(
			perpetualOrBalanceUpdate,
			bigFillBaseQuantums,
		)
	} else {
		quoteBalanceUpdate.Add(
			quoteBalanceUpdate,
			bigFillQuoteQuantums,
		)

		perpetualOrBalanceUpdate.Sub(
			perpetualOrBalanceUpdate,
			bigFillBaseQuantums,
		)
	}
}
