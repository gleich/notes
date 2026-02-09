package cli

import (
	"github.com/spf13/cobra"
	"go.mattglei.ch/timber"
)

var RootCommand = &cobra.Command{
	Use:   "notes",
	Short: "CLI for working with my notes",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			timber.Fatal(err, "failed to output help")
		}
	},
}
