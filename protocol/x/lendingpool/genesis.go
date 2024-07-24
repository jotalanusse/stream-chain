package lendingpool

import (
	"math/big"

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
		k.SetLastUpdatedTime(ctx, params.TokenDenom, uint64(ctx.BlockTime().Unix()))
		k.SetCumulativeInterestRate(ctx, params.TokenDenom, big.NewInt(types.TWENTY_SEVEN_DECIMALS))
		k.SetTotalBorrowed(ctx, params.TokenDenom, big.NewInt(0))
		k.SetCurrentBorrowAPY(ctx, params.TokenDenom, big.NewInt(0))
		k.SetLastUpdatedTotalLiquidity(ctx, params.TokenDenom, big.NewInt(0))
	}
}

// ExportGenesis returns the lending module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {

	poolParams := k.GetAllPoolParams(ctx)
	genesisPoolParams := make([]types.PoolParams, len(poolParams))

	for i, params := range poolParams {
		genesisPoolParams[i] = types.ConvertInternalToPoolParams(params)
	}

	return &types.GenesisState{
		PoolParams: genesisPoolParams,
	}
}
