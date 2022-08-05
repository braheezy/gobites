package main

import (
	"flag"
	"fmt"

	"github.com/braheezy/gobites/pkg/_2_scrabble"
	"github.com/braheezy/gobites/pkg/_3_tag_analysis"
)

func main() {
	runArg := flag.String("run", "", "Choose what to run.")

	flag.Parse()

	switch *runArg {
	case "":
		fmt.Println("No command provided. Try 'scrabble'.")
	case "scrabble":
		_2_scrabble.PlayScrabble()
	case "tags":
		_3_tag_analysis.Run()
	default:
		fmt.Println("Unknown command.")
	}
}
