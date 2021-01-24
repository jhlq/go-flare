package cmd

import (
	"fmt"

	"github.com/jhlq/go-flare/gflr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(balanceERC20Cmd)
}

var balanceERC20Cmd = &cobra.Command{
	Use:   "balanceERC20 [contractAddress] [account]",
	Short: "Check the balance of a ERC-20 token in an account.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			er("Not enough arguments.")
		}
		amount, err := gflr.BalanceERC20(args[0], args[1])
		er(err)
		fmt.Println("Balance: ", amount)
	},
}
