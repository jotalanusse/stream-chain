package keeper

import (
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/names/types"
)

var _ types.QueryServer = Keeper{}
