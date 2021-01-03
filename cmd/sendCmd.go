package cmd

import (
	"fmt"
	"strconv"

	"github.com/jhlq/go-flare/gflr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendCmd)
}

var sendCmd = &cobra.Command{
	Use:   "send [account] [amount]",
	Short: "Send FXRP to an account.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		amount, err := strconv.ParseFloat(args[1], 10)
		er(err)
		secret, err := gflr.InputSecret()
		er(err)
		tx, err := gflr.Send(secret, args[0], amount)
		er(err)
		fmt.Println("Sent transaction ", tx)
	},
}
