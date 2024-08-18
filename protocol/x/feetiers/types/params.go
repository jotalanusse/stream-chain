package types

import (
	"math"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/dtypes"
)

func (m *PerpetualFeeParams) Validate() error {
	if len(m.Tiers) == 0 {
		return ErrNoTiersExist
	}

	if m.Tiers[0].AbsoluteVolumeRequirement.IsNil() {
		m.Tiers[0].AbsoluteVolumeRequirement = dtypes.NewInt(0)
	}

	if m.Tiers[0].AbsoluteVolumeRequirement.Cmp(dtypes.NewInt(0)) != 0 ||
		m.Tiers[0].TotalVolumeShareRequirementPpm != 0 ||
		m.Tiers[0].MakerVolumeShareRequirementPpm != 0 {
		return ErrInvalidFirstTierRequirements
	}

	for i := 1; i < len(m.Tiers); i++ {
		prevTier := m.Tiers[i-1]
		currTier := m.Tiers[i]
		if currTier.AbsoluteVolumeRequirement.IsNil() {
			currTier.AbsoluteVolumeRequirement = dtypes.NewInt(0)
		}
		if prevTier.AbsoluteVolumeRequirement.Cmp(currTier.AbsoluteVolumeRequirement) == 1 ||
			prevTier.TotalVolumeShareRequirementPpm > currTier.TotalVolumeShareRequirementPpm ||
			prevTier.MakerVolumeShareRequirementPpm > currTier.MakerVolumeShareRequirementPpm {
			return ErrTiersOutOfOrder
		}
	}

	lowestMakerFee := int32(math.MaxInt32)
	lowestTakerFee := int32(math.MaxInt32)
	for _, tier := range m.Tiers {
		if tier.MakerFeePpm < lowestMakerFee {
			lowestMakerFee = tier.MakerFeePpm
		}
		if tier.TakerFeePpm < lowestTakerFee {
			lowestTakerFee = tier.TakerFeePpm
		}
	}

	// Prevent overflow
	if int64(lowestMakerFee)+int64(lowestTakerFee) < 0 {
		return ErrInvalidFee
	}

	return nil
}
