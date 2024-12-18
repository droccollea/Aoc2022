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

func isVisibleFromN(x, y int) bool {
	for i := 0; i < y; i++ {
		if input[i][x] >= input[y][x] {
			return false
		}
	}
	return true
}

func isVisibleFromE(x, y int) bool {
	for i := len(input[y]) - 1; i > x; i-- {
		if input[y][i] >= input[y][x] {
			return false
		}
	}
	return true
}

func isVisibleFromS(x, y int) bool {
	for i := len(input) - 1; i > y; i-- {
		if input[i][x] >= input[y][x] {
			return false
		}
	}
	return true
}

func isVisibleFromW(x, y int) bool {
	for i := 0; i < x; i++ {
		if input[y][i] >= input[y][x] {
			return false
		}
	}
	return true
}

func viewToN(x, y int) int {
	view := 0
	for i := y - 1; i >= 0; i-- {
		view += 1
		if input[i][x] >= input[y][x] {
			return view
		}
	}
	return view
}

func viewToE(x, y int) int {
	view := 0
	for i := x + 1; i < len(input[y]); i++ {
		view += 1
		if input[y][i] >= input[y][x] {
			return view
		}
	}
	return view
}

func viewToS(x, y int) int {
	view := 0
	for i := y + 1; i < len(input); i++ {
		view += 1
		if input[i][x] >= input[y][x] {
			return view
		}
	}
	return view
}

func viewToW(x, y int) int {
	view := 0
	for i := x - 1; i >= 0; i-- {
		view += 1
		if input[y][i] >= input[y][x] {
			return view
		}
	}
	return view
}

func main() {

	input = readInputByLine()

	// range over all "trees"
	// If visible from horizontally or vertically +1

	visible := 0

	for y, row := range input {

		for x, _ := range row {

			if isVisibleFromN(x, y) || isVisibleFromE(x, y) || isVisibleFromS(x, y) || isVisibleFromW(x, y) {
				visible += 1
			}
		}
	}
	fmt.Printf("Visible trees: %d\n", visible)

	// Scenic view.
	highestView := 0

	for y, row := range input {

		for x, _ := range row {

			thisView := viewToN(x, y) * viewToE(x, y) * viewToS(x, y) * viewToW(x, y)
			// fmt.Printf("%d,%d sees %d\n", x, y, thisView)
			if thisView > highestView {
				highestView = thisView
			}

		}
	}

	// fmt.Printf("2,3 should be 8: %d * %d * %d * %d\n", viewToN(2, 3), viewToE(2, 3), viewToS(2, 3), viewToW(2, 3))
	fmt.Printf("Highest scene: %d\n", highestView)

}
