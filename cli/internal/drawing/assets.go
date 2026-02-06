package drawing

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"go.mattglei.ch/notes/cli/internal/note"
)

func CreateAssetsFolder(document note.Note) (string, error) {
	folderPath := filepath.Join("static/drawings/", document.Slug)

	_, err := os.Stat(folderPath)
	if !errors.Is(err, fs.ErrNotExist) {
		err = os.RemoveAll(folderPath)
		if err != nil {
			return "", fmt.Errorf("removing %s: %w", folderPath, err)
		}
	}

	err = os.MkdirAll(folderPath, 0755)
	if err != nil {
		return "", fmt.Errorf("creating %s: %w", folderPath, err)
	}

	return folderPath, nil
}
