package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

type Flags struct {
	PrintLines      bool
	PrintWords      bool
	PrintCharacters bool
	PrintBytes      bool
}

func main() {
	readFlags()
	os.Exit(1)
}

func readFlags() Flags {
	var printLines, printWords, printCharacters, printBytes bool
	flag.BoolVar(&printLines, "l", false, "print lines")
	flag.BoolVar(&printWords, "w", false, "print words")
	flag.BoolVar(&printBytes, "c", false, "print bytes")
	flag.BoolVar(&printCharacters, "m", false, "print no of characters")
	flag.Parse()

	if !printLines && !printWords && !printCharacters && !printBytes {
		printLines = true
		printWords = true
		printCharacters = true
		printBytes = true
	}

	return Flags{
		PrintLines:      printLines,
		PrintWords:      printWords,
		PrintCharacters: printCharacters,
		PrintBytes:      printBytes,
	}
}

func readTxtFile(fp string) *bufio.Scanner {
	file, err := os.Open(fp)
	if err != nil {
		log.Fatalf("unable to open %s with err %s", fp, err.Error())
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalf("unable to close %s", err.Error())
		}
	}()

	return bufio.NewScanner(file)
}

func readLines(s *bufio.Scanner) int {
	numberOfLines := 0
	s.Split(bufio.ScanLines)
	for s.Scan() {
		numberOfLines++
	}

	if err := s.Err(); err != nil {
		log.Fatalf("error encountered %s", err.Error())
	}

	return numberOfLines
}
