package keeper

import (
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/types"
)

var _ types.QueryServer = Keeper{}
