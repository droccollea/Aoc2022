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

type Coord struct {
	X, Y int
}

func dirsToHead(head, tail Coord) []string {
	dirs := make([]string, 1, 2)

	if head.X > tail.X {
		dirs = append(dirs, "R")
	}
	if head.X < tail.X {
		dirs = append(dirs, "L")
	}
	if head.Y > tail.Y {
		dirs = append(dirs, "U")
	}
	if head.Y < tail.Y {
		dirs = append(dirs, "D")
	}
	return dirs
}

func moveTail(head, tail Coord) Coord {

	// If head is +1/-1/0 places away, tail stays put.
	if (head.X-tail.X == -1 || head.X-tail.X == 0 || head.X-tail.X == 1) &&
		(head.Y-tail.Y == -1 || head.Y-tail.Y == 0 || head.Y-tail.Y == 1) {
		return tail
	}
	// Move to head one place or two.
	dirs := dirsToHead(head, tail)

	newPos := tail
	for _, dir := range dirs {
		newPos = moveHead(newPos, dir)
	}
	return newPos
}

func moveHead(head Coord, dir string) Coord {
	switch dir {
	case "U":
		return Coord{head.X, head.Y + 1}
	case "D":
		return Coord{head.X, head.Y - 1}
	case "L":
		return Coord{head.X - 1, head.Y}
	case "R":
		return Coord{head.X + 1, head.Y}
	default:
		fmt.Print("Returning head!")
		return head // Should never happen!
	}
}

func main() {

	input = readInputByLine()

	rope := [10]Coord{}

	visited := make(map[Coord]int)

	// Inc initial pos.
	visited[rope[len(rope)-1]]++

	for _, v := range input {
		inst := strings.Split(v, " ")
		dir := inst[0]
		steps, _ := strconv.Atoi(inst[1])

		fmt.Printf("Moving %d steps %s\n", steps, dir)
		for i := 0; i < steps; i++ {
			rope[0] = moveHead(rope[0], dir)
			for k := 1; k < len(rope); k++ {
				rope[k] = moveTail(rope[k-1], rope[k])
				// fmt.Printf("Rope: %v visitied %d\n", rope, len(visited))
			}
			visited[rope[len(rope)-1]]++
		}
	}

	fmt.Printf("Visitied count: %d\n", len(visited))
	// 5829 low, 6037
	// long tail sample 36

}
