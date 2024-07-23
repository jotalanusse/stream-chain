package cli

import (
	"context"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdGetPoolParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-pool-params [token_denom]",
		Short: "shows pool_params",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argName := args[0]

			params := &types.QueryGetPoolParamsRequest{
				TokenDenom: argName,
			}

			res, err := queryClient.PoolParams(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
