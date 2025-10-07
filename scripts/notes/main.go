package main

import (
	"os"
	"time"

	"go.mattglei.ch/timber"
)

func main() {
	timber.TimeFormat("03:04:05")
	timber.Timezone(time.Local)

	if len(os.Args) < 2 {
		timber.FatalMsg("please provide an argument")
	}
	switch os.Args[1] {
	case "new":
		ask()
	}
}
