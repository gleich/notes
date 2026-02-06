package cli

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
	"go.mattglei.ch/notes/cli/internal/conf"
	"go.mattglei.ch/notes/cli/internal/note"
	"go.mattglei.ch/timber"
)

var (
	skipPathSet bool

	moveCommand = &cobra.Command{
		Use:   "move",
		Short: "Move the markdown files to their proper location",
		Run: func(cmd *cobra.Command, args []string) {
			if !skipPathSet {
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

			notesJSON, err := json.Marshal(notes)
			if err != nil {
				timber.Fatal(err, "failed to marshal notes json")
			}

			err = os.WriteFile("src/routes/notes.json", notesJSON, 0655)
			if err != nil {
				timber.Fatal(err, "failed to write to notes json file")
			}
			timber.Done("wrote", len(notes), "notes to notes.json")
		},
	}
)

func init() {
	moveCommand.PersistentFlags().
		BoolVarP(&skipPathSet, "skip-path-set", "s", false, "Skip moving to the path defined in the configuration file")
	RootCommand.AddCommand(moveCommand)
}
