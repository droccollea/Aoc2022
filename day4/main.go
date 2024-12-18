package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// var inputFile string = "sample.txt"

var inputFile string = "day4.txt"

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

func hasOverlap(pair string) bool {
	pairs := strings.Split(pair, ",")
	lrange := strings.Split(pairs[0], "-")
	rrange := strings.Split(pairs[1], "-")

	return isContained(lrange, rrange) || isContained(rrange, lrange)
}

func asInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func isContained(this, that []string) bool {
	// if start within the full range or end within full range
	return (asInt(this[0]) <= asInt(that[0]) && asInt(this[1]) >= asInt(that[0])) ||
		(asInt(this[0]) <= asInt(that[1]) && asInt(this[1]) >= asInt(that[1]))
}

func main() {

	input := readInputByLine()

	total := 0

	for _, pair := range input {

		if hasOverlap(pair) {
			total += 1
		}

	}

	fmt.Printf("Priority total: %d\n", total)
}
