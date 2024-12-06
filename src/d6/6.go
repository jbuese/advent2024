package day6

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	fmt.Println("--- Day 6: Guard Gallivant ---")
	fmt.Printf("C1: %d\n", c1("../data/6.txt"))
	fmt.Printf("C2: %d\n", c2("../data/6.txt"))
}

func c1(filename string) (sum int) {
	lmap := readMap(filename)
	x, y := findStartingPoint(lmap)
	lmap = move(lmap, x, y)
	sum = countOccupied(lmap)

	return
}

func c2(filename string) (count int) {
	return
}

func readMap(filename string) (lmap [][]string) {
	bytes, _ := os.ReadFile(filename)
	lines := strings.Split(strings.ToUpper(string(bytes)), "\n")
	lmap = make([][]string, len(lines))
	for i, line := range lines {
		lmap[i] = strings.Split(line, "")
	}
	return
}

func findStartingPoint(lmap [][]string) (x, y int) {
	for i, row := range lmap {
		for j, cell := range row {
			if cell == "^" {
				return i, j
			}
		}
	}
	return
}

func countOccupied(lmap [][]string) (count int) {
	for _, row := range lmap {
		for _, cell := range row {
			if cell == "X" {
				count++
			}
		}
	}
	return
}

type Direction struct {
	dx, dy int
	symbol string
	turn   string
}

var directions = map[string]Direction{
	"^": {-1, 0, "^", ">"},
	">": {0, 1, ">", "v"},
	"v": {1, 0, "v", "<"},
	"<": {0, -1, "<", "^"},
}

func move(lmap [][]string, x, y int) [][]string {
	dir := directions[lmap[x][y]]
	nx, ny := x+dir.dx, y+dir.dy

	if nx < 0 || nx >= len(lmap) || ny < 0 || ny >= len(lmap[0]) {
		lmap[x][y] = "X" // forgot this :(
		return lmap
	}

	if lmap[nx][ny] == "#" {
		lmap[x][y] = dir.turn
		return move(lmap, x, y)
	}

	lmap[x][y] = "X"
	lmap[nx][ny] = dir.symbol
	return move(lmap, nx, ny)
}
