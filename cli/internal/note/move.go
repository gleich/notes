package note

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

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
		_, err = os.Stat(destination)
		if !errors.Is(err, fs.ErrNotExist) {
			existingBin, err := os.ReadFile(destination)
			if err != nil {
				return fmt.Errorf("reading %s failed: %w", destination, err)
			}
			if string(bin) != string(existingBin) {
				err = os.WriteFile(destination, bin, 0655)
				if err != nil {
					return fmt.Errorf("failed to copy file to %s: %w", destination, err)
				}
				timber.Done("moved", filepath.Base(path))
				moved++
			}
		}

	}
	if moved != 0 {
		fmt.Println()
	}
	timber.Done("Moved", moved, "notes")
	return nil
}
