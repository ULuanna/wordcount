package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}
	//fmt.Println(src)
	numOfWords := wordCount(src)
	fmt.Println(numOfWords)
}

// wordCount returns number of words (separated by space) in source string.
func wordCount(src string) int {
	return len(strings.Split(strings.TrimSpace(src), " "))
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string, err error) {
	argsWithoutProg := os.Args
	if len(argsWithoutProg) == 1 {
		return src, errors.New("missing string to match")
	} else if len(argsWithoutProg) > 2 {
		return src, errors.New("too many arguments")
	}
	src = argsWithoutProg[1]
	return src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("wordcount:", err)
	os.Exit(1)
}
