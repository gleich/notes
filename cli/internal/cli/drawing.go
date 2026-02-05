package cli

import (
	"github.com/spf13/cobra"
	"go.mattglei.ch/notes/cli/internal/drawing"
	"go.mattglei.ch/notes/cli/internal/note"
	"go.mattglei.ch/timber"
)

var drawingCommand = &cobra.Command{
	Use:   "drawing",
	Short: "Insert the downloaded GoodNotes drawing into the current markdown file",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := drawing.Find()
		if err != nil {
			timber.Fatal(err, "failed to find drawing")
		}
		timber.Debug(path)

		notes, err := note.Notes()
		if err != nil {
			timber.Fatal(err, "failed to get notes")
		}

		note, err := note.Select("What note should the drawing be applied to?", notes)
		if err != nil {
			timber.Fatal(err, "failed to ask for which note to apply drawing to")
		}
		timber.Debug(note)
	},
}

func init() {
	RootCommand.AddCommand(drawingCommand)
}
