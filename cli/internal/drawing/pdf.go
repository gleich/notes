package drawing

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/beevik/etree"
	"go.mattglei.ch/notes/cli/internal/command"
	"go.mattglei.ch/timber"
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
	if file == "" {
		return "", errors.New("no drawing found")
	}
	return file, nil
}

func MovePDF(path string, folder string) (string, error) {
	destination := filepath.Join(folder, filepath.Base(path))
	err := os.Rename(path, destination)
	if err != nil {
		return "", fmt.Errorf("moving pdf from %s: %w", path, err)
	}
	return destination, nil
}

func ConvertPDF(path string, folder string) error {
	startingDirectory, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting working directory: %w", err)
	}

	err = os.Chdir(folder)
	if err != nil {
		return fmt.Errorf("changing directory to %s: %w", folder, err)
	}

	pdfFilename := filepath.Base(path)

	err = command.Run("pdf2svg", pdfFilename, "%d.svg", "all")
	if err != nil {
		return err
	}
	timber.Done("converted", pdfFilename, "to SVG")

	entries, err := os.ReadDir(".")
	if err != nil {
		return fmt.Errorf("reading current directory: %w", err)
	}

	for _, entry := range entries {
		name := entry.Name()
		if !strings.HasSuffix(name, ".svg") {
			continue
		}

		doc := etree.NewDocument()
		err = doc.ReadFromFile(name)
		if err != nil {
			return fmt.Errorf("parsing xml from svg: %w", err)
		}
		root := doc.SelectElement("svg")
		if root != nil {
			for pathElement := range root.ChildElementsSeq() {
				for _, attribute := range pathElement.Attr {
					if attribute.Key == "fill" &&
						(attribute.Value == "rgb(100%, 100%, 100%)" || attribute.Value == "#fff") {
						root.RemoveChild(pathElement)
					}
				}
			}
		}
		svg, err := doc.WriteToString()
		if err != nil {
			return fmt.Errorf("writing xml changes to svg: %w", err)
		}
		err = os.WriteFile(name, []byte(svg), 0655)
		if err != nil {
			return fmt.Errorf("writing removed background to svg: %w", err)
		}
		timber.Done("removed background from", name)

		err = command.Run("inkscape", "-o", name, "-D", name)
		if err != nil {
			return fmt.Errorf(": %w", err)
		}
		timber.Done("cropped", name)
	}

	err = command.Run("svgo", "-f", ".", "-o", ".")
	if err != nil {
		return err
	}

	timber.Done("optimized all SVGs")

	err = os.Remove(pdfFilename)
	if err != nil {
		return fmt.Errorf("removing %s: %w", path, err)
	}

	err = os.Chdir(startingDirectory)
	if err != nil {
		return fmt.Errorf("changing directory to %s: %w", startingDirectory, err)
	}
	return nil
}
