package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func IsTree(input string) bool {
	return input == "#"
}

func TraverseLine(treeLine string, curPos int, gradient int, counter int) (int, int){
	// Function should return the last index
	newPos := curPos + gradient // x coordinate
	// newPos must be "reinitialized" in order to account for "out of bounds"
	if newPos >= len(treeLine) {
		surplus := newPos - len(treeLine)
		newPos = surplus
	}
	stringToCheck := treeLine[newPos:newPos + 1]
	if IsTree(stringToCheck) {
		counter ++
	}
	return counter, newPos
}

func ReadFile(fileName string) []string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	treeLines := strings.Split(string(data), "\n")
	return treeLines
}

func main()  {
	gradient := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	// Initialize variables
	newPos := 0
	counter := 0

	// Initialize multiplied output
	answer := uint64(1)

	treeLines := ReadFile("input")
	//testString := "..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#"
	//treeLines := strings.Split(testString, "\n")
	for _, b := range gradient {
		counter = 0
		newPos = 0
		for treeIndex := b[1]; treeIndex < len(treeLines); treeIndex += b[1] {
			counter, newPos = TraverseLine(treeLines[treeIndex], newPos, b[0], counter)
		}
		answer = answer * uint64(counter)
	}

	// Compatible with down 1.
	//for _, line := range treeLines {
	//	counter, newPos = TraverseLine(line, newPos, gradient[1], counter)
	//}
	fmt.Println(answer)
}