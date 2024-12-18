package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// var inputFile string = "sample.txt"

var inputFile string = "input.txt"

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

type Efile struct {
	name  string
	size  int
	isDir bool
	files []*Efile
}

func getSize(f *Efile) int {
	size := 0
	if f.isDir {
		for _, d := range f.files {
			size += getSize(d)
		}
		return size
	} else {
		return f.size
	}
}

func main() {
	//One line
	input := readInputByLine()

	var filesystem = make(map[string]*Efile)
	var dirStack []string
	var currentDir *Efile

	// add the files
	var root Efile
	root.name = "/"
	root.isDir = true
	filesystem[root.name] = &root

	for _, line := range input {

		if line == "$ cd .." {
			//pop stack
			dirStack = dirStack[:len(dirStack)-1]
			currentDir = filesystem[dirStack[len(dirStack)-1]]
			fmt.Printf(".. : . is now: %s\n", currentDir.name)
		} else if strings.HasPrefix(line, "$ ls") {
			continue
		} else if strings.HasPrefix(line, "$ cd ") {
			raw := strings.SplitAfter(line, "cd ")
			fmt.Printf("cd to >%s<\n", raw[1])
			//push stack
			fqn := ""
			if raw[1] == "/" {
				fqn = "/"
				dirStack = make([]string, 1)
			} else {
				fqn = currentDir.name + raw[1]
			}
			fmt.Print("appending %s\n", fqn)

			dirStack = append(dirStack, fqn)
			fmt.Println("appended")
			currentDir = filesystem[fqn]
			fmt.Printf(". is now: %s\n", fqn)

		} else if strings.HasPrefix(line, "dir") {
			// add dir
			raw := strings.SplitAfter(line, "dir ")
			var nfile Efile
			nfile.name = currentDir.name + raw[1]
			nfile.isDir = true
			currentDir.files = append(currentDir.files, &nfile)
			filesystem[nfile.name] = &nfile
		} else { // a regular file! Append to current efile
			parts := strings.Split(line, " ")
			var f Efile
			f.size, _ = strconv.Atoi(parts[0]) //Ignore Error!
			f.name = parts[1]
			// Update the files with current list, this one plus any others
			currentDir.files = append(currentDir.files, &f)
			fmt.Printf(". files: %v\n", currentDir.files)
		}
	}
	fmt.Printf("The FS: %v\n", filesystem["a"])

	// Now iterate root files and get sizes.
	totalUpTo100K := 0
	for k, v := range filesystem {

		s := getSize(v)
		if s < 100000 {
			totalUpTo100K += s
		}
		fmt.Printf("size of %s = %d\n", k, s)
	}
	fmt.Printf("totalUpTo100K %d\n", totalUpTo100K)
	// 1584809 - low
	// 1743217

	// get the unused size (70000000 - root)
	avail := 70000000 - getSize(filesystem[root.name])
	fmt.Printf("avail = %d\n", avail)
	// take 30000000 from this to get min
	needed := 30000000 - avail
	fmt.Printf("needed = %d\n", needed)
	// get dir >= the needed, append to list then sort.
	// Now iterate root files and get sizes.
	cands := make([]int, 8)

	for _, v := range filesystem {

		s := getSize(v)
		if s >= needed {
			cands = append(cands, s)
		}
	}
	sort.Ints(cands)

	fmt.Printf("cands  %v\n", cands)
}
