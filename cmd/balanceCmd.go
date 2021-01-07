package cmd

import (
	"fmt"

	"github.com/jhlq/go-flare/gflr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(balanceCmd)
}

var balanceCmd = &cobra.Command{
	Use:   "balance [account]",
	Short: "Check the balance of an account.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		amount, err := gflr.Balance(args[0])
		er(err)
		fmt.Println("Balance: ", gflr.From18zToFloat(amount))
	},
}
