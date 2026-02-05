package notes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
	"go.mattglei.ch/timber"
)

func Move() error {
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
		newPath := filepath.Join(ROUTES_FOLDER, slug, "+page.md")

		err = os.MkdirAll(filepath.Dir(newPath), 0755)
		if err != nil {
			return fmt.Errorf("%w failed to make parent directory for: %s", err, newPath)
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("%w failed to read binary from: %s", err, path)
		}

		err = os.WriteFile(newPath, b, 0655)
		if err != nil {
			return fmt.Errorf("%w failed to copy data to %s", err, path)
		}

		timber.Done("created", newPath)

		var matter FrontMatter
		_, err = frontmatter.Parse(bytes.NewReader(b), &matter)
		if err != nil {
			return fmt.Errorf("%w failed to read frontmatter for %s", err, path)
		}

		notes = append(notes, Note{
			Title: matter.Title,
			Slug:  slug,
			Date:  matter.Date,
		})
		return nil
	})
	if err != nil {
		return fmt.Errorf("walking notes directory: %w", err)
	}

	b, err := json.Marshal(notes)
	if err != nil {
		timber.Fatal(err, "failed to json marshal documents")
	}

	path := "src/routes/notes.json"
	err = os.WriteFile(path, b, 0655)
	if err != nil {
		timber.Fatal(err, "failed to write json data to notes.json file")
	}
	fmt.Println()
	timber.Done("wrote json data to", path)
	return nil
}
