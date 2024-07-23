package keeper

import (
	"context"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PoolParams(
	c context.Context, req *types.QueryGetPoolParamsRequest) (*types.QueryPoolParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := lib.UnwrapSDKContext(c, types.ModuleName)

	val, found := k.GetPoolParams(
		ctx,
		req.TokenDenom,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryPoolParamsResponse{PoolParams: val}, nil
}
