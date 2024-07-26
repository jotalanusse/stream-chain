package lendingpool

import (
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingaccount/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingaccount/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the lendingpool module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.InitializeForGenesis(ctx)
}

// ExportGenesis returns the lending module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {

	return nil
}
