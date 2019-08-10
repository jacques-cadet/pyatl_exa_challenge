package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dhellmann/pyatl_exa_challenge/pkg/exa"
)

func main() {
	flag.Parse()

	programName := flag.Arg(0)
	if programName == "" {
		fmt.Fprintf(os.Stderr, "Please specify a program to run\n")
		os.Exit(1)
	}

	fd, err := os.Open(programName)
	defer fd.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open program file: %s\n", err)
		os.Exit(2)
	}

	program, err := exa.Load(fd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not load program: %s\n", err)
		os.Exit(2)
	}
	results, err := program.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run program: %s\n", err)
		os.Exit(3)
	}
	fmt.Printf("%v\n", results)
}
