package cmd

import (
	"fmt"

	"github.com/jhlq/go-flare/gflr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addressCmd)
}

var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "Print the address associated with a secret key.",
	Long:  `Print the address associated with a secret key or keystore.`,
	Run: func(cmd *cobra.Command, args []string) {
		secret, err := gflr.InputHidden("Enter private key: ")
		er(err)
		address, err := gflr.ToAddress(secret)
		er(err)
		fmt.Println(address)
	},
}
