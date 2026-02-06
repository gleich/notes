package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.mattglei.ch/notes/cli/internal/command"
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

		err = newNote.Create()
		if err != nil {
			timber.Fatal(err, "failed to create new note")
		}
		timber.Done("created", newNote.Path)

		err = command.Run("code", "--goto", fmt.Sprintf("%s:%d", newNote.Path, 6), newNote.Path)
		if err != nil {
			timber.Fatal(err, "failed to open file in vscode")
		}
		timber.Done("opened file in vscode")
	},
}

func init() {
	RootCommand.AddCommand(newCommand)
}
