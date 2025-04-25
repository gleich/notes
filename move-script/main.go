package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"go.mattglei.ch/timber"
)

func main() {
	timber.TimeFormat("03:04:05")

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
		newPath := filepath.Join(
			"src/routes",
			strings.TrimPrefix(dir, "notes/"),
			fmt.Sprintf("%s/+page.md", strings.TrimSuffix(filename, ".md")),
		)

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

		return nil
	})
	if err != nil {
		timber.Fatal(err, "failed to walk current working directory")
	}

}
