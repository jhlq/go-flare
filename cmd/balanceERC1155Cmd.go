package cmd

import (
	"fmt"
	"math/big"

	"github.com/jhlq/go-flare/gflr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(balanceERC1155Cmd)
}

var balanceERC1155Cmd = &cobra.Command{
	Use:   "balanceERC1155 [contractAddress] [account] [id]",
	Short: "Check the balance of ERC-1155 tokens in an account.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			er("Not enough arguments.")
		}
		id, ok := new(big.Int).SetString(args[2], 10)
		if !ok {
			er("Could not set big.Int to " + args[2])
		}
		amount, err := gflr.BalanceERC1155(args[0], args[1], id)
		er(err)
		fmt.Println("Balance: ", amount)
	},
}
