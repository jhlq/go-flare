package cmd

/*
import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(transactionsCmd)
}

var transactionsCmd = &cobra.Command{
	Use:   "txs [cutoff] [backtrack]",
	Short: "Query the blockchain for large transactions.",
	Long:  `Go through a number of blocks specified by backtrack, defaults to 100, to look for transactions that are larger than cutoff which defaults to one million.`,
	Run: func(cmd *cobra.Command, args []string) {
		co := 1000000.0
		bt := 100
		client, err := ethclient.Dial(host)
		er(err)
		block, err := client.BlockByNumber(context.Background(), nil)
		number := block.Number()
		for i := 1; i < bt; i++ {
			//to be completed after Costwo
			block, err := client.BlockByNumber(context.Background(), nil)
		}
	},
}*/
