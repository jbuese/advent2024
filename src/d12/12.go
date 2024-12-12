package day12

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("--- Day 12: Garden Groups ---")
	fmt.Printf("C1: %d\n", c1("../data/12.txt"))
	fmt.Printf("C2: %d\n", c2("../data/12.txt"))
}

func c1(filename string) (price int) {
	garden := readGarden(filename)
	visited := make(map[string]bool)
	patches := []map[string]bool{} // Store all patches

	// Iterate through each cell in the garden
	for i := range garden {
		for j := range garden[i] {
			pos := fmt.Sprintf("%d,%d", i, j)
			if !visited[pos] {
				// Found a new unvisited cell, start a new patch
				patch := bfsPatch(garden, i, j, visited)
				patches = append(patches, patch)
			}
		}
	}

	for _, patch := range patches {
		price += len(patch) * calculateCircumference(patch)
	}

	return
}

func c2(filename string) (sum int) {

	return
}

func bfsPatch(garden [][]string, startI, startJ int, visited map[string]bool) map[string]bool {
	// Define possible movements (up, right, down, left)
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	// Create queue for BFS and patch for tracking positions
	queue := [][2]int{{startI, startJ}}
	patch := make(map[string]bool)
	startChar := garden[startI][startJ]

	// Mark start position as visited and add to patch
	startPos := fmt.Sprintf("%d,%d", startI, startJ)
	visited[startPos] = true
	patch[startPos] = true

	// Process queue
	for len(queue) > 0 {
		// Pop first element
		curr := queue[0]
		queue = queue[1:]

		// Check all directions
		for _, dir := range dirs {
			newI, newJ := curr[0]+dir[0], curr[1]+dir[1]

			// Check bounds and if the new position is valid
			if newI >= 0 && newI < len(garden) &&
				newJ >= 0 && newJ < len(garden[0]) {

				pos := fmt.Sprintf("%d,%d", newI, newJ)
				// If unvisited and same character
				if !visited[pos] && garden[newI][newJ] == startChar {
					visited[pos] = true
					patch[pos] = true
					queue = append(queue, [2]int{newI, newJ})
				}
			}
		}
	}

	return patch
}

func calculateCircumference(patch map[string]bool) (circumference int) {
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for pos := range patch {
		posParts := strings.Split(pos, ",")
		posI, _ := strconv.Atoi(posParts[0])
		posJ, _ := strconv.Atoi(posParts[1])
		neighbors := 0
		for _, dir := range dirs {
			newPos := fmt.Sprintf("%d,%d", posI+dir[0], posJ+dir[1])
			if patch[newPos] {
				neighbors++
			}
		}
		circumference += (4 - neighbors)
	}
	return circumference
}

func readGarden(filename string) [][]string {
	bytes, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	garden := make([][]string, len(lines))

	for i, line := range lines {
		garden[i] = strings.Split(line, "")
	}
	return garden
}
