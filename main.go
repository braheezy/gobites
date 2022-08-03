package main

import (
	"flag"
	"fmt"

	"github.com/braheezy/gobites/pkg/_2_scrabble"
)

func main() {
	runArg := flag.String("run", "", "Choose what to run.")

	flag.Parse()

	switch *runArg {
	case "":
		fmt.Println("No command provided. Try 'scrabble'.")
	case "scrabble":
		_2_scrabble.PlayScrabble()
	default:
		fmt.Println("Unknown command.")
	}
}
