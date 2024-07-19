package keeper

import (
	"context"
	"errors"

	errorsmod "cosmossdk.io/errors"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// CreateLendingPosition handles the message to create a lending position.
func (k msgServer) CreateLendingPosition(
	goCtx context.Context,
	msg *types.MsgCreateLendingPosition,
) (*types.MsgCreateLendingAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create the lending position
	_, err := k.Keeper.CreateLendingPosition(ctx, msg.Creator, msg.Amount)
	if err != nil {
		if errors.Is(err, types.ErrAccountNotFound) {
			return nil, errorsmod.Wrapf(
				govtypes.ErrInvalidSigner,
				"account with address %s not found",
				msg.Creator,
			)
		}
		return nil, err
	}

	return &types.MsgCreateLendingAccountResponse{}, nil
}
