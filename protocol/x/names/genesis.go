package names

import (
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/names/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/names/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.InitializeForGenesis(ctx)

	for _, name := range genState.Names {
		_, err := k.CreateName(
			ctx,
			name.Id,
			name.Name,
		)
		if err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Names = k.GetAllNames(ctx)
	return genesis
}
