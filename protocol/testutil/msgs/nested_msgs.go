package msgs

import (
	upgrade "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/constants"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/encoding"
	sending "github.com/StreamFinance-Protocol/stream-chain/protocol/x/sending/types"
)

func init() {
	testEncodingCfg := encoding.GetTestEncodingCfg()
	testTxBuilder := testEncodingCfg.TxConfig.NewTxBuilder()

	_ = testTxBuilder.SetMsgs(MsgSoftwareUpgrade)
	MsgSoftwareUpgradeTxBytes, _ = testEncodingCfg.TxConfig.TxEncoder()(testTxBuilder.GetTx())

	_ = testTxBuilder.SetMsgs(MsgCancelUpgrade)
	MsgCancelUpgradeTxBytes, _ = testEncodingCfg.TxConfig.TxEncoder()(testTxBuilder.GetTx())

	_ = testTxBuilder.SetMsgs(&MsgExecWithKlyraMessage)
	MsgExecWithKlyraMessageTxBytes, _ = testEncodingCfg.TxConfig.TxEncoder()(testTxBuilder.GetTx())
}

var (
	// Inner msgs.
	MsgSoftwareUpgrade = &upgrade.MsgSoftwareUpgrade{
		Authority: constants.Bob_Num0.Owner,
		Plan: upgrade.Plan{
			Name:   "test-plan",
			Height: 10,
			Info:   "test-info",
		},
	}
	MsgSoftwareUpgradeTxBytes []byte

	MsgCancelUpgrade = &upgrade.MsgCancelUpgrade{
		Authority: constants.Bob_Num0.Owner,
	}
	MsgCancelUpgradeTxBytes []byte

	MsgExecWithKlyraMessage = authz.NewMsgExec(
		constants.AliceAccAddress,
		[]sdk.Msg{&sending.MsgCreateTransfer{}},
	)
	MsgExecWithKlyraMessageTxBytes []byte
)
