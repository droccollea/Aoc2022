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

type Monkey struct {
	Items       []int
	Op          string
	DivisableBy int
	T           int
	F           int
	Inspection  int
}

func readMonkeys() []Monkey {
	var m Monkey
	var monkeys []Monkey
	for _, line := range input {
		//Complete
		if line == "" {
			monkeys = append(monkeys, m)
			m = Monkey{}
		}

		if strings.Contains(line, "Starting") {
			itemsString1 := strings.SplitAfter(line, "Starting items: ")
			itemsString2 := strings.Split(itemsString1[1], ", ")
			for _, item := range itemsString2 {
				i, _ := strconv.Atoi(item)
				m.Items = append(m.Items, i)
			}
		}
		if strings.Contains(line, "Operation") {
			itemsString1 := strings.SplitAfter(line, "Operation: new = old ")
			m.Op = itemsString1[1]
		}
		if strings.Contains(line, "Test: divisible by") {
			itemsString1 := strings.SplitAfter(line, "Test: divisible by ")
			asInt, _ := strconv.Atoi(itemsString1[1])
			m.DivisableBy = asInt
		}
		if strings.Contains(line, "If true: throw to monkey ") {
			itemsString1 := strings.SplitAfter(line, "If true: throw to monkey ")
			m.T, _ = strconv.Atoi(itemsString1[1])
		}
		if strings.Contains(line, "If false: throw to monkey ") {
			itemsString1 := strings.SplitAfter(line, "If false: throw to monkey ")
			m.F, _ = strconv.Atoi(itemsString1[1])
		}

	}
	// append last one.
	monkeys = append(monkeys, m)

	return monkeys
}

func playRound(monkeys []Monkey, hcf int) []Monkey {
	for m, monkey := range monkeys {
		for _, item := range monkey.Items {
			operation := strings.Split(monkey.Op, " ")
			newItem := item
			switch operation[0] {
			case ("+"):
				if operand, err := strconv.Atoi(operation[1]); err != nil {
					if operation[1] != "old" {
						log.Fatal("unknown op", err)
					}
					newItem += item
				} else {
					newItem += operand
				}
			case ("*"):
				if operand, err := strconv.Atoi(operation[1]); err != nil {
					if operation[1] != "old" {
						log.Fatal("unknown op", err)
					}
					newItem *= item
				} else {
					newItem *= operand
				}
			default:
				log.Fatal("Unknown case: ", operation[0])
			}
			// Monkey bored. Divide by 3
			// newItem /= 3

			// Keep it sensible?
			newItem = (newItem % hcf)

			// Throw to next monkey.
			if newItem%monkey.DivisableBy == 0 {
				monkeys[monkey.T].Items = append(monkeys[monkey.T].Items, newItem)
			} else {
				monkeys[monkey.F].Items = append(monkeys[monkey.F].Items, newItem)
			}
		}
		monkeys[m].Inspection += len(monkey.Items)
		// No more items, clear list.
		monkeys[m].Items = []int{}
	}
	return monkeys
}

func findHCF(monkeys []Monkey) int {
	hcf := 19
	agreed := 0
	for ; hcf < 9999999; hcf += 19 {
		for _, monkey := range monkeys {
			if hcf%monkey.DivisableBy != 0 {
				agreed = 0
				break
			} else {
				agreed++
			}
		}
		if agreed == len(monkeys) {
			fmt.Printf("HCF is: %d\n", hcf)
			break
		}
	}
	return hcf
}

func main() {

	input = readInputByLine()

	// Read the monkeys.
	monkeys := readMonkeys()

	fmt.Printf("Monkeys %v\n", monkeys)

	hcf := findHCF(monkeys)
	for round := 0; round < 10000; round++ {
		//play a round
		monkeys = playRound(monkeys, hcf)
	}

	fmt.Printf("Monkeys %v\n", monkeys)
	first := 0
	for _, monkey := range monkeys {
		if monkey.Inspection > first {
			first = monkey.Inspection
		}
	}
	second := 0
	for _, monkey := range monkeys {
		if monkey.Inspection > second && monkey.Inspection < first {
			second = monkey.Inspection
		}
	}

	fmt.Printf("Total: %d * %d = %d\n", first, second, first*second)
	//14893585630 - low
	//14271202792 - low
	//14952185856
}
