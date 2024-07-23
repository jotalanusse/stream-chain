package lendingpool

import (
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the lendingpool module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.InitializeForGenesis(ctx)

	// Set all the lending pool params
	for _, params := range genState.PoolParams {
		if err := k.CreatePoolParams(ctx, params); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the lending module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		PoolParams: k.GetAllPoolParams(ctx),
	}
}
