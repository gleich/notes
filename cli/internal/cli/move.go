package cli

import (
	"github.com/spf13/cobra"
	"go.mattglei.ch/notes/cli/internal/conf"
	"go.mattglei.ch/notes/cli/internal/note"
	"go.mattglei.ch/timber"
)

var (
	skipPathValidation bool

	moveCommand = &cobra.Command{
		Use:   "move",
		Short: "Move the markdown files to their proper location",
		Run: func(cmd *cobra.Command, args []string) {
			if !skipPathValidation {
				config, err := conf.Read()
				if err != nil {
					timber.Fatal(err, "failed to read config")
				}
				err = config.GoToPath()
				if err != nil {
					timber.Fatal(err, "failed to validate path")
				}
			}

			notes, err := note.Notes()
			if err != nil {
				timber.Fatal(err, "failed to get list of notes")
			}

			err = note.Move(notes)
			if err != nil {
				timber.Fatal(err, "failed to move notes")
			}
		},
	}
)

func init() {
	moveCommand.PersistentFlags().
		BoolVarP(&skipPathValidation, "skip-path-validation", "s", false, "If path validation should be skipped. For use in build environment")
	RootCommand.AddCommand(moveCommand)
}
