package cli

import (
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

// CmdDepositToLendingPool deposits funds from a liquidity provider's x/bank account to a lending pool.
func CmdDepositToLendingPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit-to-lending-pool [liquidity_provider] [token_denom] [amount]",
		Short: "Deposit funds from an account to a lending pool.",
		Long: `Deposit funds from an account to a lending pool.
Note, the '--from' flag is ignored as it is implied from [liquidity_provider].
[token_denom] specifies the token of the pool you are depositing to.
[amount] specifies the amount to deposit.
`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Liquidity provider address validation done in `ValidateBasic()` below.
			argLiquidityProvider := args[0]
			err = cmd.Flags().Set(flags.FlagFrom, argLiquidityProvider)
			if err != nil {
				return err
			}

			argTokenDenom := args[1]
			argAmount := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDepositLiquidityIntoPool(
				clientCtx.GetFromAddress().String(),
				argTokenDenom,
				argAmount,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
