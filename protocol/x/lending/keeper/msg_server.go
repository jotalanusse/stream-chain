package keeper

import (
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/assets/types"
)

//TODO: implement the interface for the LendingKeeper

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
