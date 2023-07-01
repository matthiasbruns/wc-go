package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	exitCodeOK = iota
	exitCodeError
	exitCodeFileError
)

// read file into byte array
func readFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func main() {
	fileName := ""
	countPtr := flag.Bool("c", false, "count bytes")
	linePtr := flag.Bool("l", false, "count lines")
	wordPtr := flag.Bool("w", false, "count words")
	charsPtr := flag.Bool("m", false, "count chars")

	// last arg is file name
	if len(os.Args) > 0 {
		argsWithoutProg := os.Args[1:]
		if len(argsWithoutProg) > 0 {

			lastArg := argsWithoutProg[len(argsWithoutProg)-1]
			if !strings.HasPrefix(lastArg, "-") {
				fileName = lastArg
			}
		}
	}

	// react cli flags
	flag.Parse()

	// if no flags, default to -c -l -w
	if !*countPtr && !*linePtr && !*wordPtr && !*charsPtr {
		*countPtr = true
		*linePtr = true
		*wordPtr = true
	}

	var file []byte
	if fileName == "" {
		// read from stdin
		stdin, _ := io.ReadAll(os.Stdin)
		if stdin != nil {
			file = stdin
		}
	}

	if len(file) == 0 {
		if contents, err := readFile(fileName); err != nil {
			fmt.Printf("Error reading file: %v", err)
			os.Exit(exitCodeFileError)
		} else {
			file = contents
		}
	}

	if *linePtr {
		// get lines in file
		lines := 0
		for _, b := range file {
			if b == '\n' {
				lines++
			}
		}
		fmt.Printf("\t%d\t", lines)
	}
	if *wordPtr {
		words := len(bytes.Fields(file))
		fmt.Printf("\t%d\t", words)
	}

	if *charsPtr {
		chars := len(bytes.Runes(file))
		fmt.Printf("\t%d\t", chars)
	}

	if *countPtr {
		fmt.Printf("\t%d\t", len(file))
	}

	fmt.Println(fileName)
}
