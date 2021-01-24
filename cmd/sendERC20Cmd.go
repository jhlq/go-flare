package cmd

import (
	"fmt"
	"strconv"

	"github.com/jhlq/go-flare/gflr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendERC20Cmd)
	sendERC20Cmd.Flags().StringVarP(&secret, "secret", "s", "", "Private key (without 0x), this flag is for use in scripts.")
	sendERC20Cmd.Flags().StringVarP(&ks, "keystore", "k", "", "Keystore filename (put it in the go-flare-config folder)")
	sendERC20Cmd.Flags().StringVarP(&passphrase, "password", "p", "", "Passphrase for keystore file.")
}

var sendERC20Cmd = &cobra.Command{
	Use:   "sendERC20 [contractAddress] [account] [amount]",
	Short: "Send ERC-20 tokens to an account.",
	Long:  `Send ERC-20 tokens to an account. The ERC-20 interface is equivalent to "FLR-20". Do not send non-fungible tokens with this function.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			er("Not enough arguments.")
		}
		amount, err := strconv.ParseFloat(args[2], 10)
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
		tx, err := gflr.SendERC20(secret, args[0], args[1], amount)
		er(err)
		fmt.Println("Sent transaction ", tx)
	},
}
