package cmd

import (
  "fmt"
  "strconv"

  "github.com/spf13/cobra"
  "github.com/jhlq/go-flare/gflr"
)

func init() {
  rootCmd.AddCommand(sendCmd)
}

var sendCmd = &cobra.Command{
  Use:   "send [account] [amount]",
  Short: "Send FxRP to an account.",
  Long:  `Currently this is limited to about 10000000 and doesn't take decimals into account.`,
  Run: func(cmd *cobra.Command, args []string) {
    amount, err := strconv.ParseFloat(args[1], 10)
    er(err)
    secret, err := gflr.InputSecret()
    er(err)
    err = gflr.Send(secret, args[0], amount)
    er(err)
    fmt.Println("Sent transaction.")
  },
}
