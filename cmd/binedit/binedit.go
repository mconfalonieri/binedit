package main

import (
	"log"

	"github.com/pborman/getopt/v2"
)

// Main program. It can be called with a file to be opened.
func main() {
	getopt.Parse()
	args := getopt.Args()
	numArgs := len(args)
	if numArgs > 1 {
		log.Fatalf("This command expects one command line argument. %d given.", numArgs)
	}
}
