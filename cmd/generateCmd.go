package cmd

import (
	"fmt"

	"github.com/jhlq/go-flare/gflr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new wallet.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		address, secret, err := gflr.GenerateWallet()
		er(err)
		fmt.Println("Address: ", address)
		fmt.Println("Private key: ", secret)
	},
}
