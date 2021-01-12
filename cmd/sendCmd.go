package cmd

import (
	"fmt"
	"strconv"

	"github.com/jhlq/go-flare/gflr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVarP(&secret, "secret", "s", "", "Private key (without 0x), this flag is for use in scripts.")
	sendCmd.Flags().StringVarP(&ks, "keystore", "k", "", "Keystore filename (put it in the go-flare-config folder)")
	sendCmd.Flags().StringVarP(&passphrase, "password", "p", "", "Passphrase for keystore file.")
}

var sendCmd = &cobra.Command{
	Use:   "send [account] [amount]",
	Short: "Send FXRP to an account.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		amount, err := strconv.ParseFloat(args[1], 10)
		er(err)
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
		tx, err := gflr.Send(secret, args[0], amount)
		er(err)
		fmt.Println("Sent transaction ", tx)
	},
}
