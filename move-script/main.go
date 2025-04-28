package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"go.mattglei.ch/timber"
)

type Document struct {
	Title string    `json:"title"`
	Slug  string    `json:"slug"`
	Date  time.Time `json:"date"`
}

type FrontMatter struct {
	Title string    `yaml:"title"`
	Date  time.Time `yaml:"date"`
}

func main() {
	timber.TimeFormat("03:04:05")

	documents := []Document{}
	err := filepath.WalkDir("./notes/", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("%w failed to walk: %s", err, path)
		}
		if d.IsDir() || d.Name() == "README.md" {
			return nil
		}

		// given a path like notes/college/sofa-103/syllabus.md it should generate
		// src/routes/college/sofa-103/syllabus/+page.md
		dir, filename := filepath.Split(path)
		slug := filepath.Join(
			strings.TrimPrefix(dir, "notes/"),
			strings.TrimSuffix(filename, ".md"),
		)
		newPath := filepath.Join("src/routes/(notes)/", slug, "+page.md")

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

		documents = append(documents, Document{
			Title: matter.Title,
			Slug:  slug,
			Date:  matter.Date,
		})
		return nil
	})
	if err != nil {
		timber.Fatal(err, "failed to walk current working directory")
	}

	b, err := json.Marshal(documents)
	if err != nil {
		timber.Fatal(err, "failed to json marshal documents")
	}

	path := "src/routes/notes.json"
	err = os.WriteFile(path, b, 0655)
	if err != nil {
		timber.Fatal(err, "failed to write json data to notes.json file")
	}
	timber.Done("wrote json data to", path)
}
