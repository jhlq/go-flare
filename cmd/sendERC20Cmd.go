package cmd

import (
	"fmt"
	"strconv"

	"github.com/jhlq/go-flare/gflr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendERC20Cmd)
}

var sendERC20Cmd = &cobra.Command{
	Use:   "sendERC20 [contractAddress] [account] [amount]",
	Short: "Send ERC-20 tokens to an account.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		amount, err := strconv.ParseFloat(args[2], 10)
		er(err)
		secret, err := gflr.InputHidden("Enter private key (without 0x): ")
		er(err)
		tx, err := gflr.SendERC20(secret, args[0], args[1], amount)
		er(err)
		fmt.Println("Sent transaction ", tx)
	},
}
