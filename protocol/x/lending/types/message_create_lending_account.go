package types

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Ensure MsgCreateLendingAccount implements the sdk.Msg interface
var _ sdk.Msg = &MsgCreateLendingAccount{}

// ValidateBasic performs basic validation on the message fields.
func (msg *MsgCreateLendingAccount) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return errorsmod.Wrap(
			ErrInvalidAddress,
			fmt.Sprintf(
				"creator '%s' must be a valid bech32 address, but got error '%v'",
				msg.Creator,
				err.Error(),
			),
		)
	}
	return nil
}
