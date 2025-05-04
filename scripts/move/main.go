package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"go.mattglei.ch/timber"
)

const ROUTES_FOLDER = "src/routes/(notes)/"

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

	cleanRoutes()

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

func cleanRoutes() {
	var dirs []string
	err := filepath.WalkDir(ROUTES_FOLDER, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walking %s: %w", path, err)
		}
		if d.IsDir() {
			dirs = append(dirs, path)
			return nil
		}
		if filepath.Ext(d.Name()) == ".md" {
			if err := os.Remove(path); err != nil {
				return fmt.Errorf("removing %s: %w", path, err)
			}
		}
		return nil
	})
	if err != nil {
		timber.Fatal(err, "failed cleaning out svelte files")
	}

	sort.Slice(dirs, func(i, j int) bool {
		return len(dirs[i]) > len(dirs[j])
	})

	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		if len(entries) == 0 {
			err = os.Remove(dir)
			if err != nil {
				timber.Fatal(err, "failed to delete", dir)
			}
		}
	}
}
