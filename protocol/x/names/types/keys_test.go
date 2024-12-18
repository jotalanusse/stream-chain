package types_test

import (
	"testing"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/names/types"
	"github.com/stretchr/testify/require"
)

func TestModuleKeys(t *testing.T) {
	require.Equal(t, "names", types.ModuleName)
	require.Equal(t, "names", types.StoreKey)
}

func TestStateKeys(t *testing.T) {
	require.Equal(t, "Name:", types.NameKeyPrefix)
}
