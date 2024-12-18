package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// const cols = 3

// var inputFile string = "sample.txt"

const cols = 9

var inputFile string = "day5.txt"

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

func readCrates(input []string) [cols]string {

	var crates [cols]string
	for _, line := range input {
		// ignore non-crate. Break after crates " 1   2..."
		if line[1] == '1' {
			break
		}

		// Got a line of crates now.
		for i := 0; i < cols; i++ {
			// Is there a crate? Is so prepend it to the col.
			// From pos 1, crates are every 4 cols taking out [] and spaces.
			// 0 == 1, 1 == 5, 2 == 9 :
			cpos := (i * 4) + 1
			// fmt.Printf("col %d, cpos %d\n", i, cpos)
			if line[cpos] != ' ' {
				if crates[i] != "" {
					crates[i] = string(line[cpos]) + crates[i]
				} else {
					crates[i] = string(line[cpos])
				}
			}
		}

	}
	return crates
}

type Instruction struct {
	count       int
	source      int
	destination int
}

func asInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func readInstructions(input []string) (instructions []Instruction) {

	for _, line := range input {
		// ignore non-instruction.
		if len(line) > 5 && line[0] == 'm' {

			// Genearte an instruction.
			stuff := strings.Split(line, " ")
			// Fix source/dest as 0 based.
			ins := Instruction{asInt(stuff[1]), asInt(stuff[3]) - 1, asInt(stuff[5]) - 1}
			instructions = append(instructions, ins)
		}
	}
	return instructions
}

func applyInstructions(crates [cols]string, ins []Instruction) []string {

	updatedCrates := crates[:]
	for _, in := range ins {
		// fmt.Printf("inst: %v\n", in)
		// for i := 0; i < in.count; i++ {
		c := updatedCrates[in.source][len(updatedCrates[in.source])-in.count:]
		// fmt.Printf("moving %c\n", c)
		// Push
		updatedCrates[in.destination] = updatedCrates[in.destination] + c
		// Pop
		updatedCrates[in.source] = updatedCrates[in.source][0 : len(updatedCrates[in.source])-in.count]
		// }
	}
	return updatedCrates
}

func topCrates(stacks []string) string {

	var tops string
	for _, s := range stacks {
		c := s[len(s)-1]
		tops = tops + string(c)
	}
	return tops
}

func main() {

	input := readInputByLine()

	crates := readCrates(input)
	fmt.Printf("Crates are: %s\n", crates)

	instructions := readInstructions(input)
	fmt.Printf("Instructons are: %v\n", instructions)

	newcrates := applyInstructions(crates, instructions)
	fmt.Printf("Crates are now: %s\n", newcrates)

	tops := topCrates(newcrates)
	fmt.Printf("Tops are: %s\n", tops)
}
