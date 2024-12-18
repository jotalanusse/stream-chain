package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/store/prefix"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/names/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO(CLOB-863) Add tests for these endpoints.
func (k Keeper) AllNames(
	c context.Context,
	req *types.QueryAllNamesRequest,
) (*types.QueryAllNamesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var names []types.Name
	ctx := lib.UnwrapSDKContext(c, types.ModuleName)

	store := ctx.KVStore(k.storeKey)
	nameStore := prefix.NewStore(store, []byte(types.NameKeyPrefix))

	pageRes, err := query.Paginate(nameStore, req.Pagination, func(key []byte, value []byte) error {
		var name types.Name
		if err := k.cdc.Unmarshal(value, &name); err != nil {
			return err
		}

		names = append(names, name)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNamesResponse{Name: names, Pagination: pageRes}, nil
}

func (k Keeper) Name(c context.Context, req *types.QueryNameRequest) (*types.QueryNameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := lib.UnwrapSDKContext(c, types.ModuleName)

	val, exists := k.GetName(
		ctx,
		req.Id,
	)

	if !exists {
		return nil,
			status.Error(
				codes.NotFound,
				fmt.Sprintf(
					"Name id %+v not found.",
					req.Id,
				),
			)
	}

	return &types.QueryNameResponse{Name: val}, nil
}
