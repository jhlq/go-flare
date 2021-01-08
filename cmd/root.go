package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	keystore string

	rootCmd = &cobra.Command{
		Use:   "go-flare",
		Short: "Tools for the Flare network",
		Long:  `Flare leverages the Ethereum Virtual Machine without relying on either PoW or PoS.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	//rootCmd.PersistentFlags().StringVar(&keystore, "keystore", "", "filename of keystore (not implemented yet)")
}

func er(msg interface{}) {
	if msg != nil {
		fmt.Println("Error:", msg)
		os.Exit(1)
	}
}
