package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	X int
	Y int
}

var height map[Coord]rune
var blocked map[Coord]bool

var start Coord
var end Coord
var routes [][]Coord

func buildHeightMap(input []string) {
	for y, row := range input {
		for x, h := range row {
			if h == rune('S') {
				start = Coord{x, y}
				height[Coord{x, y}] = 'a'
			} else if h == rune('E') {
				end = Coord{x, y}
				height[Coord{x, y}] = 'z'
			} else {
				height[Coord{x, y}] = h
			}
		}
	}
}

func buildDistanceMap(from Coord) map[Coord]int {

	distance := make(map[Coord]int)

	counted := append([]Coord{}, from)
	distance[from] = 0

	for step := 1; len(counted) > 0; step++ {

		justCounted := []Coord{}
		for _, current := range counted {
			// Right
			right := Coord{current.X + 1, current.Y}
			if right.X < len(input[0]) && !(distance[right] > 0) && height[right] <= height[current]+1 {
				distance[right] = step
				justCounted = append(justCounted, right)
			}

			// Down
			down := Coord{current.X, current.Y + 1}
			if down.Y < len(input) && !(distance[down] > 0) && height[down] <= height[current]+1 {
				distance[down] = step
				justCounted = append(justCounted, down)
			}

			// Left
			left := Coord{current.X - 1, current.Y}
			if left.X >= 0 && !(distance[left] > 0) && height[left] <= height[current]+1 {
				distance[left] = step
				justCounted = append(justCounted, left)
			}

			// Up
			up := Coord{current.X, current.Y - 1}
			if up.Y >= 0 && !(distance[up] > 0) && height[up] <= height[current]+1 {
				distance[up] = step
				justCounted = append(justCounted, up)
			}

		}
		// Prep next count.
		counted = justCounted
	}
	return distance
}

func distanceFrom(from Coord) int {
	distance := buildDistanceMap(from)
	return distance[end]
}

func distanceFromA() int {
	// Get the 'a' nodes
	aNodes := []Coord{}
	for k, v := range height {
		if v == 'a' {
			aNodes = append(aNodes, k)
		}
	}

	toEnd := 999
	// For each a node, get distance to end.
	for _, a := range aNodes {
		steps := distanceFrom(a)
		fmt.Printf("From %v it's %d\n", a, steps)
		if steps > 0 && steps < toEnd {
			toEnd = steps
		}
	}
	return toEnd

}

func plot(route []Coord) {
	for y, row := range input {
		for x := range row {
			if contains(route, Coord{x, y}) {
				fmt.Print(strings.ToUpper(fmt.Sprintf("%c", height[Coord{x, y}])))
				// } else if blocked[Coord{x, y}] {
				// 	fmt.Print("x")
			} else {
				fmt.Printf("%c", height[Coord{x, y}])
			}
		}
		fmt.Println()
	}
}

func contains(slice []Coord, element Coord) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func visit(current Coord, visited []Coord) bool {

	// fmt.Printf("Visiting %v, visited %d\n", current, len(visited))

	if len(visited)%200 == 0 {
		fmt.Printf("Length %d at %v, height %c\n", len(visited), current, height[current])
	}
	// if height[current] > 'm' {
	// 	fmt.Printf("At %c %v\n", height[current], current)
	// }

	if current == end {
		fmt.Printf("Found a route of length %d\n", len(visited))
		routes = append(routes, append(visited, current))
		return true
	}

	validPath := false
	// Visit neightbouring nodes if not yet visited and not beyond height restriction or edges.

	// Try staying at this level or +1 first.
	// Right
	right := Coord{current.X + 1, current.Y}
	if right.X < len(input[0]) && !contains(visited, right) && (height[right] == height[current] || height[right] == height[current]+1) && !blocked[right] {
		if visit(right, append(visited, current)) {
			validPath = true
		}
	}

	// Down
	down := Coord{current.X, current.Y + 1}
	if down.Y < len(input) && !contains(visited, down) && (height[down] == height[current] || height[down] == height[current]+1) && !blocked[down] {
		if visit(down, append(visited, current)) {
			validPath = true
		}
	}

	// Left
	left := Coord{current.X - 1, current.Y}
	if left.X >= 0 && !contains(visited, left) && (height[left] == height[current] || height[left] == height[current]+1) && !blocked[left] {
		if visit(left, append(visited, current)) {
			validPath = true
		}
	}

	// Up
	up := Coord{current.X, current.Y - 1}
	if up.Y >= 0 && !contains(visited, up) && (height[up] == height[current] || height[up] == height[current]+1) && !blocked[up] {
		if visit(up, append(visited, current)) {
			validPath = true
		}
	}

	// Now drop.
	if right.X < len(input[0]) && !contains(visited, right) && (height[right] < height[current]) && !blocked[right] {
		if visit(right, append(visited, current)) {
			validPath = true
		}
	}

	// Up
	if up.Y >= 0 && !contains(visited, up) && (height[up] < height[current]) && !blocked[up] {
		if visit(up, append(visited, current)) {
			validPath = true
		}
	}

	// Left
	if left.X >= 0 && !contains(visited, left) && (height[left] < height[current]) && !blocked[left] {
		if visit(left, append(visited, current)) {
			validPath = true
		}
	}

	// Down
	if down.Y < len(input) && !contains(visited, down) && (height[down] < height[current]) && !blocked[down] {
		if visit(down, append(visited, current)) {
			validPath = true
		}
	}

	// fmt.Printf("Done here, current %v, at length %d.\n", current, len(visited))
	if !validPath {
		blocked[current] = true
		// fmt.Printf("Blocking: %v", current)
	}
	return validPath

}

func main() {

	input = readInputByLine()

	height = make(map[Coord]rune)
	blocked = make(map[Coord]bool, 0)

	buildHeightMap(input)

	fmt.Printf("Start: %v, End %v\n", start, end)
	// visited := []Coord{}
	// visit(start, visited)

	// fmt.Printf("Routes: \n%v", routes)
	steps := 0
	for _, route := range routes {
		fmt.Printf("Route length: %d\n", len(route))
		if steps == 0 || len(route) < steps {

			steps = len(route)
		}
	}
	// Plot the shortest route.
	for _, v := range routes {
		if len(v) == steps {
			plot(v)
		}
	}
	// Route minus starting pos
	fmt.Printf("Min steps: %d\n", steps-1)

	//448 - high
	//422 - high
	//415 - high (random guess to see if even near)
	// Just guessing it now :(
	// 405 - wrong!

	// Recursive looping fails with larger puzzle input.
	// Alternative?
	// Map which a's can get to a b, b's get to c's etc - still navigate peaks.

	// New approach. Forget recursing and deadends, step back from start repeatedly until no more steps
	// Step back from start and count the steps. Will grow exponentially at first but then should reduce and find the shortest path.

	fmt.Printf("Distance to End is %d\n", distanceFrom(start))
	// plotDistances()
	//394 - correct!

	fmt.Printf("Shortest distance to End from nearest 'a' %d\n", distanceFromA())

}
