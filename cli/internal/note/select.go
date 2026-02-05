package note

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
	"go.mattglei.ch/notes/cli/internal/prompt"
	"go.mattglei.ch/notes/cli/internal/styles"
)

func Select(title string, notes []Note) (Note, error) {
	options := []huh.Option[Note]{}
	for _, note := range notes {
		options = append(
			options,
			huh.NewOption(
				fmt.Sprintf(
					"%s %s",
					note.Title,
					styles.Grey.Render(
						fmt.Sprintf("[%s]", strings.TrimPrefix(note.Path, "notes/")),
					),
				),
				note,
			),
		)
	}

	var selected Note
	err := huh.NewSelect[Note]().Title(title).
		Options(options...).
		Value(&selected).
		WithTheme(prompt.Theme).
		Run()
	if err != nil {
		return Note{}, fmt.Errorf("failed to run selection prompt: %w", err)
	}
	return selected, nil
}
