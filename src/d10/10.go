package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("--- Day 10: Hoof It ---")
	fmt.Printf("C1: %d\n", c1("../data/10.txt"))
	fmt.Printf("C2: %d\n", c2("../data/10.txt"))
}

func c1(filename string) (sum int) {
	topo := readMap(filename)
	heads := findHeads(topo)

	for _, head := range heads {
		visited := make(map[string]bool)
		sum += findMountainTrails(topo, head, visited, true)
	}

	return
}

func c2(filename string) (sum int) {
	topo := readMap(filename)
	heads := findHeads(topo)

	for _, head := range heads {
		visited := make(map[string]bool)
		sum += findMountainTrails(topo, head, visited, false)
	}

	return
}

func readMap(filename string) [][]int {
	bytes, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	topo := make([][]int, len(lines))

	for i, line := range lines {
		nums := strings.Split(line, "")
		topo[i] = make([]int, len(nums))
		for j, num := range nums {
			n, _ := strconv.Atoi(num)
			topo[i][j] = n
		}
	}
	return topo
}

func findHeads(topo [][]int) (heads [][]int) {
	for i, row := range topo {
		for j, num := range row {
			if num == 0 {
				heads = append(heads, []int{i, j})
			}
		}
	}
	return
}

func findMountainTrails(topo [][]int, head []int, visited map[string]bool, unique bool) int {
	directions := [][]int{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}

	key := fmt.Sprintf("%d,%d", head[0], head[1])
	visited[key] = true
	if !unique {
		defer delete(visited, key)
	}

	if topo[head[0]][head[1]] == 9 {
		return 1
	}

	score := 0
	for _, dir := range directions {
		newI, newJ := head[0]+dir[0], head[1]+dir[1]
		newKey := fmt.Sprintf("%d,%d", newI, newJ)

		if newI >= 0 && newI < len(topo) && newJ >= 0 && newJ < len(topo[0]) {
			if !visited[newKey] && topo[newI][newJ] == topo[head[0]][head[1]]+1 {
				score += findMountainTrails(topo, []int{newI, newJ}, visited, unique)
			}
		}
	}

	return score
}
