package cmd

import (
	"fmt"
	"math/big"

	"github.com/jhlq/go-flare/gflr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ownerERC721Cmd)
}

var ownerERC721Cmd = &cobra.Command{
	Use:   "ownerERC721 [contractAddress] [id]",
	Short: "Check the owner of a ERC-721 token.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			er("Not enough arguments.")
		}
		id, ok := new(big.Int).SetString(args[1], 10)
		if !ok {
			er("Could not set big.Int to " + args[1])
		}
		owner, err := gflr.OwnerERC721(args[0], id)
		er(err)
		fmt.Println("Owner: ", owner)
	},
}
