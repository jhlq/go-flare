package cmd

import (
	"fmt"
	"strconv"

	"github.com/jhlq/go-flare/gflr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(blockCmd)
}

var blockCmd = &cobra.Command{
	Use:   "block [number]",
	Short: "Query information about a block.",
	Long:  `Query information about a block. Omit number to fetch latest block.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		b := -1
		if len(args) > 0 {
			b, err = strconv.Atoi(args[0])
			er(err)
		}
		block, err := gflr.Block(b)
		er(err)
		fmt.Println("Block number: ", block.Number())
		fmt.Println("Difficulty: ", block.Difficulty().Uint64())
		fmt.Println("Hash: ", block.Hash().Hex())
		fmt.Println("Transactions: ", len(block.Transactions()))
		if len(block.Transactions()) > 0 {
			v, err := block.Transactions()[0].MarshalJSON()
			er(err)
			fmt.Println("Transaction 1 info:", string(v))
		}
		fmt.Println("Gas used: ", block.GasUsed())
	},
}
