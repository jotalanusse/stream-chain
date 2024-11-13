package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/perpetuals/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (k msgServer) SetMultiCollateralAssets(
	goCtx context.Context,
	msg *types.MsgSetMultiCollateralAssets,
) (*types.MsgSetMultiCollateralAssetsResponse, error) {
	if !k.Keeper.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)
	k.Keeper.SetMultiCollateralAssets(ctx, msg.MultiCollateralAssets)

	return &types.MsgSetMultiCollateralAssetsResponse{}, nil
}
