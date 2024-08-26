package types

import (
	blocktypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/blocktime/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//nolint:staticcheck
)

type BlockTimeKeeper interface {
	GetPreviousBlockInfo(ctx sdk.Context) blocktypes.BlockInfo
}
