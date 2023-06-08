package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "ktucli",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().ParseErrorsWhitelist.UnknownFlags {
			return errors.New("Please enter right flag. ")
		}
		return cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("workspace-name", "", "specifying the workspaces name ")
	rootCmd.PersistentFlags().String("output-type", "table", "specifying the output type  [json, table]")
}
