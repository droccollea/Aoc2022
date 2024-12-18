package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// var inputFile string = "sample.txt"

var inputFile string = "input.txt"

var input []string

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

func main() {

	input = readInputByLine()

	correctPairs := 0
	leftLevel := 0
	rightLevel := 0
	// 3 lines at a time, compare.
	for line := 0; line < len(input); line += 3 {
		lpos := 0
		rpos := 0

		// loop till this pair compared.
		for {
			if input[line][lpos] == '[' {
				leftLevel++
				lpos++
			}
			if input[line+1][rpos] == '[' {
				rightLevel++
				rpos++
			}
			if input[line][lpos] == ']' {
				leftLevel--
				lpos++
			}
			if input[line+1][rpos] == ']' {
				rightLevel--
				rpos++
			}
			if input[line][lpos] == ',' {
				lpos++
			}
			if input[line+1][rpos] == ',' {
				rpos++
			}

		}
	}

	fmt.Printf("Correct Pairs: %d", correctPairs)
}
