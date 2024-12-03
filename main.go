package main

import (
	"flag"
	"os"
)

type Flags struct {
	PrintLines bool
}

func main() {
	readFlags()
	os.Exit(1)
}

func readFlags() Flags {
	var printLines bool
	flag.BoolVar(&printLines, "lines", false, "print lines")
	flag.Parse()

	return Flags{
		PrintLines: printLines,
	}
}
