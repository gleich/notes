package main

import (
	"os"

	"go.mattglei.ch/timber"
)

func main() {
	if len(os.Args) < 2 {
		timber.FatalMsg("please provide an argument")
	}
	switch os.Args[1] {
	case "new":
		ask()
	}
}
