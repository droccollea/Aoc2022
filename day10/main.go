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

func exec(inst string, X int) int {
	// Add whatever to X

	// fmt.Printf("goingto split >%s<\n", inst)
	splitIns := strings.SplitAfter(inst, "addx ")
	operand, err := strconv.Atoi(splitIns[1])
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("From %v, split %v, adding %d to x:%d\n", inst, splitIns, operand, X)
	return (X + operand)
}

func main() {

	input = readInputByLine()

	// Map input to cycles.
	instructions := []string{}
	for _, in := range input {

		// Append another noop and the actual instruction.
		if in != "noop" {
			instructions = append(instructions, "executing", in)
			// instructions = append(instructions, in)
		} else {
			instructions = append(instructions, "noop")
		}
	}

	total := 0
	X := 1

	crt := [6]string{}

	for cycle, inst := range instructions {
		if (cycle+21)%40 == 0 && cycle > 0 {
			total += X * (cycle + 1)
			fmt.Printf("at cycle %d, X is %d, total %d\n", cycle, X, total)
		}
		if cycle > 240 {
			break
		}

		// Might need moved up above exec?
		row := cycle / 40
		if X == (cycle%40) || (X+1) == (cycle%40) || (X-1) == (cycle%40) {
			crt[row] += "#"
		} else {
			crt[row] += "."
		}
		if inst != "noop" && inst != "executing" {
			X = exec(inst, X)
		}
	}

	fmt.Printf("Total: %d\n", total)

	fmt.Printf("CRT: %s\n", crt)
	for _, i := range crt {
		fmt.Println(i)
	}
}
