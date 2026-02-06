package cli

import (
	"fmt"

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

		notes, err := note.Notes()
		if err != nil {
			timber.Fatal(err, "failed to get notes")
		}

		note, err := note.Select("What note should the drawing be applied to?", notes)
		if err != nil {
			timber.Fatal(err, "failed to ask for which note to apply drawing to")
		}

		folder, err := drawing.CreateAssetsFolder(note)
		if err != nil {
			timber.Fatal(err, "failed to create assets folder for drawing")
		}

		movedPath, err := drawing.MovePDF(path, folder)
		if err != nil {
			timber.Fatal(err, "failed to move", path, "to", folder)
		}

		err = drawing.ConvertPDF(movedPath, folder)
		if err != nil {
			timber.Fatal(err, "failed to convert pdf to webp images")
		}
		fmt.Println()
		timber.Done("created SVGs from PDF. Will be injected into markdown files when moved")
	},
}

func init() {
	RootCommand.AddCommand(drawingCommand)
}
