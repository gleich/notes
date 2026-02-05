package main

import (
	"go.mattglei.ch/notes/cli/internal/cli"
	"go.mattglei.ch/timber"
)

func main() {
	err := cli.RootCommand.Execute()
	if err != nil {
		timber.Fatal(err, "failed to run root command")
	}
}
