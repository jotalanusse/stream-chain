package keeper_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/dtypes"
	testapp "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/app"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/constants"
	epochstypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/epochs/types"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/stats/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()

	logger := tApp.App.StatsKeeper.Logger(ctx)
	require.NotNil(t, logger)
}

type recordFillArgs struct {
	taker    string
	maker    string
	notional *big.Int
}

func TestRecordFill(t *testing.T) {
	tests := map[string]struct {
		args               []recordFillArgs
		expectedBlockStats *types.BlockStats
	}{
		"no fills": {
			[]recordFillArgs{},
			&types.BlockStats{
				Fills: nil,
			},
		},
		"single fill": {
			[]recordFillArgs{
				{"taker", "maker", new(big.Int).SetUint64(123)},
			},
			&types.BlockStats{
				Fills: []*types.BlockStats_Fill{
					{
						Taker:    "taker",
						Maker:    "maker",
						Notional: constants.Serializable_Int_123,
					},
				},
			},
		},
		"multiple fills": {
			[]recordFillArgs{
				{"alice", "bob", new(big.Int).SetUint64(123)},
				{"bob", "alice", new(big.Int).SetUint64(321)},
			},
			&types.BlockStats{
				Fills: []*types.BlockStats_Fill{
					{
						Taker:    "alice",
						Maker:    "bob",
						Notional: constants.Serializable_Int_123,
					},
					{
						Taker:    "bob",
						Maker:    "alice",
						Notional: constants.Serializable_Int_321,
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tApp := testapp.NewTestAppBuilder(t).Build()
			ctx := tApp.InitChain()
			k := tApp.App.StatsKeeper

			for _, fill := range tc.args {
				k.RecordFill(ctx, fill.taker, fill.maker, fill.notional)
			}
			require.Equal(t, tc.expectedBlockStats, k.GetBlockStats(ctx))
		})
	}
}

func TestProcessBlockStats(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()

	// Epochs initialize at block height 2
	tApp.AdvanceToBlock(2, testapp.AdvanceToBlockOptions{
		BlockTime: time.Unix(1, 0).UTC(),
	})
	ctx := tApp.AdvanceToBlock(10, testapp.AdvanceToBlockOptions{
		BlockTime: time.Unix(int64(epochstypes.StatsEpochDuration)+1, 0).UTC(),
	})
	k := tApp.App.StatsKeeper

	k.SetBlockStats(ctx, &types.BlockStats{
		Fills: []*types.BlockStats_Fill{
			{
				Taker:    "alice",
				Maker:    "bob",
				Notional: constants.Serializable_Int_5,
			},
			{
				Taker:    "bob",
				Maker:    "alice",
				Notional: constants.Serializable_Int_10,
			},
		},
	})
	k.ProcessBlockStats(ctx)

	assert.Equal(t, &types.GlobalStats{
		NotionalTraded: constants.Serializable_Int_15,
	}, k.GetGlobalStats(ctx))
	assert.Equal(t, &types.UserStats{
		TakerNotional: constants.Serializable_Int_5,
		MakerNotional: constants.Serializable_Int_10,
	}, k.GetUserStats(ctx, "alice"))
	assert.Equal(t, &types.UserStats{
		TakerNotional: constants.Serializable_Int_10,
		MakerNotional: constants.Serializable_Int_5,
	}, k.GetUserStats(ctx, "bob"))
	assert.Equal(t, &types.EpochStats{
		EpochEndTime: time.Unix(7200, 0).UTC(),
		Stats: []*types.EpochStats_UserWithStats{
			{
				User: "alice",
				Stats: &types.UserStats{
					TakerNotional: constants.Serializable_Int_5,
					MakerNotional: constants.Serializable_Int_10,
				},
			},
			{
				User: "bob",
				Stats: &types.UserStats{
					TakerNotional: constants.Serializable_Int_10,
					MakerNotional: constants.Serializable_Int_5,
				},
			},
		},
	}, k.GetEpochStatsOrNil(ctx, 1))

	k.SetBlockStats(ctx, &types.BlockStats{
		Fills: []*types.BlockStats_Fill{
			{
				Taker:    "bob",
				Maker:    "alice",
				Notional: constants.Serializable_Int_10,
			},
		},
	})
	k.ProcessBlockStats(ctx)
	assert.Equal(t, &types.GlobalStats{
		NotionalTraded: constants.Serializable_Int_25,
	}, k.GetGlobalStats(ctx))
	assert.Equal(t, &types.UserStats{
		TakerNotional: constants.Serializable_Int_5,
		MakerNotional: constants.Serializable_Int_20,
	}, k.GetUserStats(ctx, "alice"))
	assert.Equal(t, &types.UserStats{
		TakerNotional: constants.Serializable_Int_20,
		MakerNotional: constants.Serializable_Int_5,
	}, k.GetUserStats(ctx, "bob"))
	assert.Equal(t, &types.EpochStats{
		EpochEndTime: time.Unix(7200, 0).UTC(),
		Stats: []*types.EpochStats_UserWithStats{
			{
				User: "alice",
				Stats: &types.UserStats{
					TakerNotional: constants.Serializable_Int_5,
					MakerNotional: constants.Serializable_Int_20,
				},
			},
			{
				User: "bob",
				Stats: &types.UserStats{
					TakerNotional: constants.Serializable_Int_20,
					MakerNotional: constants.Serializable_Int_5,
				},
			},
		},
	}, k.GetEpochStatsOrNil(ctx, 1))
}

func TestExpireOldStats(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()

	// Epochs start at block height 2
	ctx := tApp.AdvanceToBlock(2, testapp.AdvanceToBlockOptions{
		BlockTime: time.Unix(int64(1), 0).UTC(),
	})
	windowDuration := tApp.App.StatsKeeper.GetWindowDuration(ctx)
	// 5 epochs are out of the window
	tApp.AdvanceToBlock(3, testapp.AdvanceToBlockOptions{
		BlockTime: time.Unix(0, 0).
			Add(windowDuration).
			Add((time.Duration(5*epochstypes.StatsEpochDuration) + 1) * time.Second).
			UTC(),
	})
	ctx = tApp.AdvanceToBlock(100, testapp.AdvanceToBlockOptions{})
	k := tApp.App.StatsKeeper

	// Create a bunch of EpochStats.
	// Odd epochs don't have stats. 30 epochs total.
	for i := 0; i < 30; i++ {
		k.SetEpochStats(ctx, uint32(i*2), &types.EpochStats{
			EpochEndTime: time.Unix(0, 0).
				Add(time.Duration(i*int(epochstypes.StatsEpochDuration)) * time.Second).
				UTC(),
			Stats: []*types.EpochStats_UserWithStats{
				{
					User: "alice",
					Stats: &types.UserStats{
						TakerNotional: constants.Serializable_Int_1,
						MakerNotional: constants.Serializable_Int_2,
					},
				},
				{
					User: "bob",
					Stats: &types.UserStats{
						TakerNotional: constants.Serializable_Int_2,
						MakerNotional: constants.Serializable_Int_1,
					},
				},
			},
		})
	}
	k.SetUserStats(ctx, "alice", &types.UserStats{
		TakerNotional: constants.Serializable_Int_30,
		MakerNotional: constants.Serializable_Int_60,
	})
	k.SetUserStats(ctx, "bob", &types.UserStats{
		TakerNotional: constants.Serializable_Int_60,
		MakerNotional: constants.Serializable_Int_30,
	})
	k.SetGlobalStats(ctx, &types.GlobalStats{
		NotionalTraded: constants.Serializable_Int_90,
	})
	k.SetStatsMetadata(ctx, &types.StatsMetadata{
		TrailingEpoch: 0,
	})

	// Prune epochs in batches of 2. For each pair, the second epoch is nil.
	// Epochs 1-10 pruned.
	for i := 0; i < 6; i++ {
		// EpochStats exist before pruning
		require.NotNil(t, k.GetEpochStatsOrNil(ctx, uint32(i*2)))

		k.ExpireOldStats(ctx)
		require.Equal(t, &types.UserStats{
			TakerNotional: dtypes.NewIntFromUint64(30 - uint64(i+1)),
			MakerNotional: dtypes.NewIntFromUint64(60 - 2*uint64(i+1)),
		}, k.GetUserStats(ctx, "alice"))
		require.Equal(t, &types.UserStats{
			TakerNotional: dtypes.NewIntFromUint64(60 - 2*uint64(i+1)),
			MakerNotional: dtypes.NewIntFromUint64(30 - uint64(i+1)),
		}, k.GetUserStats(ctx, "bob"))
		require.Equal(t, &types.GlobalStats{
			NotionalTraded: dtypes.NewIntFromUint64(90 - 3*uint64(i+1)),
		}, k.GetGlobalStats(ctx))

		// EpochStats removed
		require.Nil(t, k.GetEpochStatsOrNil(ctx, uint32(i*2)))

		k.ExpireOldStats(ctx)

		// Unchanged after pruning nil epoch
		require.Equal(t, &types.UserStats{
			TakerNotional: dtypes.NewIntFromUint64(30 - uint64(i+1)),
			MakerNotional: dtypes.NewIntFromUint64(60 - 2*uint64(i+1)),
		}, k.GetUserStats(ctx, "alice"))
		require.Equal(t, &types.UserStats{
			TakerNotional: dtypes.NewIntFromUint64(60 - 2*uint64(i+1)),
			MakerNotional: dtypes.NewIntFromUint64(30 - uint64(i+1)),
		}, k.GetUserStats(ctx, "bob"))
		require.Equal(t, &types.GlobalStats{
			NotionalTraded: dtypes.NewIntFromUint64(90 - 3*uint64(i+1)),
		}, k.GetGlobalStats(ctx))
	}

	// Epoch 12 is within the window so it won't get pruned
	k.ExpireOldStats(ctx)
	k.ExpireOldStats(ctx)
	k.ExpireOldStats(ctx)
	require.NotNil(t, k.GetEpochStatsOrNil(ctx, uint32(12)))
}
