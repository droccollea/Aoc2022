package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// var inputFile string = "sample.txt"

var inputFile string = "day6.txt"

func readInputByLine() []string {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var content []string
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return content
}

func hasDupChar(s string) bool {

	for _, c := range s {
		count := 0

		for i := 0; i < 14; i++ {
			if c == rune(s[i]) {
				count++
			}
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func main() {
	//One line
	input := readInputByLine()[0]

	// iterate line from char 4.
	// Check prev 4 char for dups and continue until none.
	// Return the last index before breaking.
	for i := 14; i < len(input); i++ {
		s := input[i-14 : i]
		if !hasDupChar(s) {
			fmt.Printf("found marker %s at index %d\n", s, i)
			break
		}
	}
}
