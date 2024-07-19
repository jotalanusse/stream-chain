package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// CreateLendingAccount handles the message to create a lending account.
func (k msgServer) CreateLendingAccount(
	goCtx context.Context,
	msg *types.MsgCreateLendingAccount,
) (*types.MsgCreateLendingAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the account already exists
	exists, err := k.Keeper.DoesLendingAccountExist(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"account with address %s already exists",
			msg.Creator,
		)
	}

	// Create the lending account
	_, err = k.Keeper.CreateLendingAccount(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateLendingAccountResponse{}, nil
}
