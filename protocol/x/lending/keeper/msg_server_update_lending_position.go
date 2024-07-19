package keeper

import (
	"context"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateLendingPosition(
	goCtx context.Context,
	msg *types.MsgUpdateLendingPosition,
) (*types.MsgUpdateLendingPositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Update the lending position
	_, err := k.Keeper.UpdateLendingPosition(ctx, msg.Creator, msg.Asset, msg.PositionId)
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateLendingPositionResponse{}, nil
}
