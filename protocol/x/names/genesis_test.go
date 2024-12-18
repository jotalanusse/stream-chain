package names_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	keepertest "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/names"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/names/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/names/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	expected := types.DefaultGenesis()
	ctx, k, _, _, _, _ := keepertest.NamesKeepers(t, true)
	names.InitGenesis(ctx, *k, *expected)
	assertNameCreateEventsInIndexerBlock(t, k, ctx, len(expected.Names))
	actual := names.ExportGenesis(ctx, *k)
	require.NotNil(t, actual)
	require.ElementsMatch(t, actual.Names, expected.Names)
}

// assertNameCreateEventsInIndexerBlock checks that the number of name create events
// included in the Indexer block kafka message.
func assertNameCreateEventsInIndexerBlock(
	t *testing.T,
	k *keeper.Keeper,
	ctx sdk.Context,
	numNames int,
) {
	nameEvents := keepertest.GetNameCreateEventsFromIndexerBlock(ctx, k)
	require.Len(t, nameEvents, numNames)
}
