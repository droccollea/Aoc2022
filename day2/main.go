package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// var inputFile string = "sample.txt"

var inputFile string = "day2.txt"

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

const Win = 6
const Draw = 3

func scoreRound(e, m byte) int {

	// A/X == Rock
	// B/Y == Paper
	// C/Z == Scissors
	switch m {
	// lose
	case 'X':
		if e == 'A' {
			return scoreHand('Z')
		}
		if e == 'B' {
			return scoreHand('X')
		}
		if e == 'C' {
			return scoreHand('Y')
		}
	// Draw
	case 'Y':
		return Draw + scoreHand(e)
	// Win
	case 'Z':
		if e == 'A' {
			return Win + scoreHand('Y')
		}
		if e == 'B' {
			return Win + scoreHand('Z')
		}
		if e == 'C' {
			return Win + scoreHand('X')
		}
	}

	return 0 // Lose
}

func scoreHand(m byte) int {

	// A/X == Rock
	// B/Y == Paper
	// C/Z == Scissors
	switch m {
	case 'X', 'A':
		return 1
	case 'Y', 'B':
		return 2
	case 'Z', 'C':
		return 3
	}

	return 0 // Lose
}

func main() {

	input := readInputByLine()
	currentScore := 0

	// read and score line
	for _, v := range input {
		elf := v[0]
		me := v[2]

		currentScore += scoreRound(elf, me)
		fmt.Printf("Current score is %d\n", currentScore)

	}

	fmt.Printf("Final score is %d\n", currentScore)

	// 10235 - too low

}
