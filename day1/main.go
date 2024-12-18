package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

//var inputFile string = "sample.txt"
var inputFile string = "day1.txt"

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
	input := readInputByLine()

	bigElf := 0
	bigElfCals := 0
	currentElf := 0
	currentCals := 0

	var allCals []int

	for _, val := range input {
		//fmt.Printf("these cals %s\n", val)

		// Next elf, check if this one is biggest.
		if val == "" {
			allCals = append(allCals, currentCals)
			if currentCals > bigElfCals {
				bigElf = currentElf
				bigElfCals = currentCals
			}
			// Next and reset cals.
			currentElf++
			currentCals = 0
		} else {
			if i, err := strconv.Atoi(val); err == nil {
				currentCals += i
				//fmt.Printf("current cals %d\n", currentCals)
			} else {
				log.Fatal(err)
			}
		}
	}

	// Sort the cals and take top 3.
	sort.Sort(sort.Reverse(sort.IntSlice(allCals)))

	total := 0
	for _, v := range allCals[0:3] {
		total += v
		fmt.Printf("add cals %d\n", v)
	}

	fmt.Printf("Big Elf is %d, with %d cals\n", bigElf+1, bigElfCals)

	fmt.Printf("Top 3 total %d cals\n", total)
	//141817 too low

}
