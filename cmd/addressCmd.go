package cmd

import (
	"fmt"

	"github.com/jhlq/go-flare/gflr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addressCmd)
	addressCmd.Flags().StringVarP(&secret, "secret", "s", "", "Private key (without 0x), this flag is for use in scripts.")
	addressCmd.Flags().StringVarP(&ks, "keystore", "k", "", "Keystore filename (put it in the go-flare-config folder)")
	addressCmd.Flags().StringVarP(&passphrase, "password", "p", "", "Passphrase for keystore file.")
}

var secret string
var ks string
var passphrase string

var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "Print the address associated with a secret key.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if ks != "" {
			if passphrase == "" {
				passphrase, err = gflr.InputHidden("Enter password: ")
				er(err)
			}
			secret, err = gflr.Unlock(ks, passphrase)
			er(err)
		}
		if secret == "" {
			secret, err = gflr.InputHidden("Enter private key (without 0x): ")
			er(err)
		}
		address, err := gflr.ToAddress(secret)
		er(err)
		fmt.Println(address)
	},
}
