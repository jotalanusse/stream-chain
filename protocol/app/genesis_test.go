package app_test

import (
	"encoding/json"
	"os"
	"testing"

	testapp "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/app"

	"github.com/stretchr/testify/require"
)

func TestDefaultGenesisState(t *testing.T) {
	app := testapp.DefaultTestApp(nil)
	defaultGenesisState := app.DefaultGenesis()
	humanReadableDefaultGenesisState, jsonUnmarshalErr := json.Marshal(&defaultGenesisState)

	expectedDefaultGenesisState, fileReadErr := os.ReadFile("testdata/default_genesis_state.json")

	require.NoError(t, fileReadErr)
	require.NoError(t, jsonUnmarshalErr)
	require.JSONEq(t, string(expectedDefaultGenesisState), string(humanReadableDefaultGenesisState))
}
