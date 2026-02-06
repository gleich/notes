package note

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.mattglei.ch/timber"
)

func Move(notes []Note) error {
	moved := 0
	for _, note := range notes {
		path := note.Path
		destination := note.DestinationPath
		bin, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading %s failed: %w", path, err)
		}
		markdown := string(bin)

		err = os.MkdirAll(filepath.Dir(destination), 0755)
		if err != nil {
			return fmt.Errorf("%w failed to make parent directory for: %s", err, destination)
		}

		// inject drawing links
		var (
			patchedMarkdown = strings.Builder{}
			drawingIndex    = 1
		)
		for line := range strings.Lines(markdown) {
			if line == "<!-- DRAWING -->\n" {
				path := filepath.Join(
					"/drawings/",
					note.Slug,
					fmt.Sprintf("%d.svg", drawingIndex),
				)
				systemPath := filepath.Join("static", path)
				// check to make sure file exists
				_, err = os.Stat(systemPath)
				if err != nil {
					return err
				}

				patchedMarkdown.WriteString(
					"\n<div class=\"drawing\"><div class=\"drawing-scale\">",
				)
				fmt.Fprintf(
					&patchedMarkdown,
					`<img alt="Drawing #%d" src="%s"/>`,
					drawingIndex,
					path,
				)
				patchedMarkdown.WriteString("</div></div>\n")
				drawingIndex++
			} else {

				patchedMarkdown.WriteString(line)
			}
		}

		err = os.WriteFile(destination, []byte(patchedMarkdown.String()), 0655)
		if err != nil {
			return fmt.Errorf("failed to copy file to %s: %w", destination, err)
		}
		timber.Done("moved", filepath.Base(path))
		moved++

	}
	if moved != 0 {
		fmt.Println()
	}
	timber.Done("Moved", moved, "notes")
	return nil
}
