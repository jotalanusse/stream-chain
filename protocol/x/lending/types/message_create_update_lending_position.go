package types

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Ensure MsgUpdateLendingPosition implements the sdk.Msg interface
var _ sdk.Msg = &MsgUpdateLendingPosition{}

// ValidateBasic performs basic validation on the message fields.
func (msg *MsgUpdateLendingPosition) ValidateBasic() error {
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
	if msg.PositionId == "" {
		return errorsmod.Wrap(
			ErrInvalidPositionId,
			"position ID must not be empty",
		)
	}
	if msg.Asset.Amount.IsNegative() {
		return errorsmod.Wrap(
			ErrInvalidAmount,
			"amount must be non-negative",
		)
	}
	return nil
}
