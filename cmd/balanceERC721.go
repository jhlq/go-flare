package cmd

import (
	"fmt"

	"github.com/jhlq/go-flare/gflr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(balanceERC721Cmd)
}

var balanceERC721Cmd = &cobra.Command{
	Use:   "balanceERC721 [contractAddress] [account]",
	Short: "Check the balance of ERC-721 NFTs in an account.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			er("Not enough arguments.")
		}
		amount, err := gflr.BalanceERC721(args[0], args[1])
		er(err)
		fmt.Println("Balance: ", amount)
	},
}
