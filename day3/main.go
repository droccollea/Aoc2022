package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// var inputFile string = "sample.txt"

var inputFile string = "day3.txt"

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

func calcPriority(b byte) int {
	// int A=65 int a=97
	if int(b) >= 97 {
		return int(b) - 96
	} else {
		return int(b) - 38
	}
}

func commonChars(left, right string) (commons string) {
	// var commons = []string
	for _, c := range left {
		for _, d := range right {
			if c == d {
				if len(commons) > 0 && (commons[0]) == byte(c) {
					continue
				} else {
					commons += string(c)
				}
			}
		}
	}
	return commons
}

func main() {

	input := readInputByLine()

	total := 0

	// Iterate in blocks of 3.
	for i := 0; i < len(input); i += 3 {
		remaining := commonChars(commonChars(input[i], input[i+1]), input[i+2])

		if len(remaining) > 1 {
			log.Fatal("Found too many common chars in: $s", remaining)
		}

		total += calcPriority(remaining[0])

	}

	fmt.Printf("Priority total: %d\n", total)
}
