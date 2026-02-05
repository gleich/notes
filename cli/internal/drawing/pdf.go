package drawing

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// finds the most recent drawing the ~/Documents folder
func Find() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getting user's home directory: %w", err)
	}
	documentsFolder := filepath.Join(home, "Documents")

	entries, err := os.ReadDir(documentsFolder)
	if err != nil {
		return "", fmt.Errorf("reading directory %s failed: %w", documentsFolder, err)
	}

	var (
		file    string
		modTime time.Time
	)
	for _, entry := range entries {
		name := entry.Name()
		if !entry.IsDir() && strings.HasSuffix(name, "drawing.pdf") {
			path := filepath.Join(documentsFolder, name)
			stat, err := os.Stat(path)
			if err != nil {
				return "", fmt.Errorf("getting status of %s: %w", path, err)
			}
			modified := stat.ModTime()
			if modified.After(modTime) {
				file = path
				modTime = modified
			}
		}
	}
	return file, nil
}
