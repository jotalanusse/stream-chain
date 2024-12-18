package keeper_test

import (
	"testing"

	keepertest "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	ctx, keeper, _, _, _, _ := keepertest.NamesKeepers(t, true)
	logger := keeper.Logger(ctx)
	require.NotNil(t, logger)
}
