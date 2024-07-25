package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) DepositLiquidityIntoPool(ctx context.Context, msg *types.MsgDepositLiquidityIntoPool) (*types.MsgDepositLiquidityIntoPoolResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Validate account sender.
	liquidityProvider, err := sdk.AccAddressFromBech32(msg.LiquidityProvider)
	if err != nil {
		return nil, types.ErrInvalidAccountAddress
	}

	amount, err := types.ConvertStringToBigInt(msg.Amount)
	if err != nil {
		return nil, types.ErrInvalidDepositAmount
	}

	_, exists := k.Keeper.GetPoolParams(sdkCtx, msg.TokenDenom)
	if !exists {
		return nil, types.ErrInvalidTokenDenom
	}

	err = k.Keeper.DepositLiquidity(sdkCtx, amount, liquidityProvider, msg.TokenDenom)
	if err != nil {
		return nil, err
	}

	return &types.MsgDepositLiquidityIntoPoolResponse{}, nil
}
