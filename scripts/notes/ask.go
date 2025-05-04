package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
	"go.mattglei.ch/timber"
)

func ask() {
	var (
		name string
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Note Name?").Value(&name),
		),
	).WithTheme(huh.ThemeBase())
	err := form.Run()
	if err != nil {
		timber.Fatal(err, "failed to ask form for new note")
	}

	timber.Debug(fmt.Sprintf("%s.md", strings.ReplaceAll(strings.ToLower(name), " ", "-")))
}
