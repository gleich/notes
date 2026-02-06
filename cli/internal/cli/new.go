package cli

import (
	"github.com/spf13/cobra"
	"go.mattglei.ch/notes/cli/internal/conf"
	"go.mattglei.ch/notes/cli/internal/note"
	"go.mattglei.ch/timber"
)

var newCommand = &cobra.Command{
	Use:   "new",
	Short: "Create a new note",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := conf.Read()
		if err != nil {
			timber.Fatal(err, "failed to read config")
		}
		err = config.GoToPath()
		if err != nil {
			timber.Fatal(err, "failed to validate path")
		}

		newNote, err := note.Ask()
		if err != nil {
			timber.Fatal(err, "failed to ask for new note")
		}
		timber.Debug(newNote.Title)
		timber.Debug(newNote.Path)
	},
}

func init() {
	RootCommand.AddCommand(newCommand)
}
