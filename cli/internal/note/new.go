package note

import (
	"cmp"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/goccy/go-yaml"
	"go.mattglei.ch/notes/cli/internal/prompt"
)

func Ask() (Note, error) {
	folders := []huh.Option[string]{}

	err := filepath.WalkDir(NOTES_DIRECTORY, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walking %s: %w", path, err)
		}
		if d.IsDir() && path != NOTES_DIRECTORY {
			path = strings.TrimPrefix(path, "notes/")
			folders = append(folders, huh.NewOption(path, path))
		}
		return nil
	})
	if err != nil {
		return Note{}, fmt.Errorf("walking %s: %w", NOTES_DIRECTORY, err)
	}

	slices.SortFunc(folders, func(a, b huh.Option[string]) int {
		if c := cmp.Compare(len(b.Key), len(a.Key)); c != 0 {
			return c
		}
		return cmp.Compare(a.Key, b.Key)
	})

	var (
		folder string
		note   Note
	)
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Title?").Value(&note.Title),
			huh.NewSelect[string]().Title("Folder?").
				Options(folders...).
				Value(&folder).
				Filtering(true),
		),
	).WithTheme(prompt.Theme)
	err = form.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			os.Exit(0)
		}
		return Note{}, fmt.Errorf("running from to create new note: %w", err)
	}

	filename := strings.ToLower(strings.ReplaceAll(note.Title, " ", "-")) + ".md"
	note.Path = filepath.Join(NOTES_DIRECTORY, folder, filename)
	note.Date = time.Now()

	return note, nil
}

func (n Note) Create() error {
	_, err := os.Stat(n.Path)
	if !errors.Is(err, fs.ErrNotExist) {
		return fs.ErrExist
	}

	meta, err := yaml.Marshal(metadata{
		Title: n.Title,
		Date:  n.Date,
	})
	if err != nil {
		return fmt.Errorf("encoding metadata to YAML: %w", err)
	}

	markdown := strings.Builder{}
	markdown.WriteString("---\n")
	markdown.WriteString(string(meta))
	markdown.WriteString("---\n\n")

	err = os.WriteFile(n.Path, []byte(markdown.String()), 0655)
	if err != nil {
		return fmt.Errorf("writing to %s: %w", n.Path, err)
	}

	return nil
}
