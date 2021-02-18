package cmd

import (
	"fmt"
	"math/big"

	"github.com/jhlq/go-flare/gflr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sendERC1155Cmd)
	sendERC1155Cmd.Flags().StringVarP(&secret, "secret", "s", "", "Private key (without 0x), this flag is for use in scripts.")
	sendERC1155Cmd.Flags().StringVarP(&ks, "keystore", "k", "", "Keystore filename (put it in the go-flare-config folder)")
	sendERC1155Cmd.Flags().StringVarP(&passphrase, "password", "p", "", "Passphrase for keystore file.")
}

var sendERC1155Cmd = &cobra.Command{
	Use:   "sendERC1155 [contractAddress] [account] [id] [amount]",
	Short: "Send ERC-1155 tokens to an account.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 4 {
			er("Not enough arguments.")
		}
		id, ok := new(big.Int).SetString(args[2], 10)
		if !ok {
			er("Could not set big.Int to " + args[2])
		}
		amount, ok := new(big.Int).SetString(args[3], 10)
		if !ok {
			er("Could not set big.Int to " + args[3])
		}
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
		tx, err := gflr.SendERC1155(secret, args[0], args[1], id, amount)
		er(err)
		fmt.Println("Sent transaction ", tx)
	},
}
