package cmd

import (
  "fmt"

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
    fmt.Println("This will print the address")
  },
}
