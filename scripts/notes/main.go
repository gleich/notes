package main

import "os"

func main() {
	switch os.Args[1] {
	case "new":
		ask()
	}
}
