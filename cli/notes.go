package main

import (
	"time"

	"go.mattglei.ch/notes/cli/internal/cli"
	"go.mattglei.ch/timber"
)

func main() {
	timber.Timezone(time.Local)
	timber.TimeFormat("03:04:05")

	err := cli.RootCommand.Execute()
	if err != nil {
		timber.Fatal(err, "failed to run root command")
	}
}
