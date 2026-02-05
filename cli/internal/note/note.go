package note

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
)

type Note struct {
	Title string    `json:"title"`
	Slug  string    `json:"slug"`
	Date  time.Time `json:"dates"`

	Path            string `json:"-"`
	DestinationPath string `json:"-"`
}

type metadata struct {
	Title string    `yaml:"title"`
	Date  time.Time `yaml:"date"`
}

func Notes() ([]Note, error) {
	notes := []Note{}
	err := filepath.WalkDir("./notes/", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walking %s: %w", path, err)
		}
		if d.IsDir() || d.Name() == "README.md" {
			return nil
		}

		dir, filename := filepath.Split(path)
		slug := filepath.Join(
			strings.TrimPrefix(dir, "notes/"),
			strings.TrimSuffix(filename, ".md"),
		)
		destination := filepath.Join("src/routes/(notes)/", slug, "+page.md")

		err = os.MkdirAll(filepath.Dir(destination), 0755)
		if err != nil {
			return fmt.Errorf("%w failed to make parent directory for: %s", err, destination)
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("%w failed to read binary from: %s", err, path)
		}

		var meta metadata
		_, err = frontmatter.Parse(bytes.NewReader(b), &meta)
		if err != nil {
			return fmt.Errorf("%w failed to read frontmatter for %s", err, path)
		}

		notes = append(notes, Note{
			Title:           meta.Title,
			Slug:            slug,
			Date:            meta.Date,
			Path:            path,
			DestinationPath: destination,
		})
		return nil
	})
	if err != nil {
		return []Note{}, fmt.Errorf("walking notes directory failed: %w", err)
	}

	sort.SliceStable(notes, func(i, j int) bool {
		return notes[i].Date.After(notes[j].Date)
	})

	return notes, nil
}
