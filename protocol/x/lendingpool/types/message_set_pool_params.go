package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgSetPoolParams{}

func (msg *MsgSetPoolParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic runs validation on the fields of a MsgSetPoolParams.
func (msg *MsgSetPoolParams) ValidateBasic() error {

	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return ErrInvalidAuthorityAddress
	}

	_, err := msg.PoolParams.Validate()
	if err != nil {
		return err
	}

	return nil
}
