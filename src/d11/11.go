package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("--- Day 11: Plutonian Pebbles ---")
	fmt.Printf("C1: %d\n", c1("../data/11.txt"))
	fmt.Printf("C2: %d\n", c2("../data/11.txt"))
}

func c1(filename string) int {
	stones := readStones(filename)
	for i := 1; i <= 25; i++ {
		stones = applyRules(stones)
	}

	return len(stones)
}

type CountedGroup struct {
	value int
	count int
}

func c2(filename string) (total int) {
	stones := readStones(filename)

	groups := make(map[int]int)
	for _, stone := range stones {
		groups[stone]++
	}

	initialGroups := make([]CountedGroup, 0, len(groups))
	for v, c := range groups {
		initialGroups = append(initialGroups, CountedGroup{v, c})
	}

	currentGroups := initialGroups
	for i := 1; i <= 75; i++ {
		currentGroups = applyRulesOptimized(currentGroups)
	}

	for _, g := range currentGroups {
		total += g.count
	}

	return
}

func applyRulesOptimized(groups []CountedGroup) []CountedGroup {
	newCountedGroups := make(map[int]int)

	for _, g := range groups {
		if g.value == 0 { // 0 becomes 1
			newCountedGroups[1] += g.count
			continue
		}

		digits := 1
		for temp := g.value; temp >= 10; temp /= 10 {
			digits++
		}

		if digits%2 == 0 {
			divisor := 1
			for i := 0; i < digits/2; i++ {
				divisor *= 10
			}
			first := g.value / divisor
			second := g.value % divisor
			newCountedGroups[first] += g.count
			newCountedGroups[second] += g.count
		} else {
			newCountedGroups[g.value*2024] += g.count
		}
	}

	// Convert map back to slice
	result := make([]CountedGroup, 0, len(newCountedGroups))
	for v, c := range newCountedGroups {
		result = append(result, CountedGroup{v, c})
	}
	return result
}

func applyRules(stones []int) []int {
	newStones := make([]int, 0, len(stones)*2)

	for _, n := range stones {
		if n == 0 {
			newStones = append(newStones, 1)
			continue
		}

		// Let's count digits properly first
		digits := 1
		for temp := n; temp >= 10; temp /= 10 {
			digits++
		}

		if digits%2 == 0 {
			// Calculate divisor directly for the specific number
			divisor := 1
			for i := 0; i < digits/2; i++ {
				divisor *= 10
			}

			first := n / divisor
			second := n % divisor

			newStones = append(newStones, first, second)
		} else {
			newStones = append(newStones, n*2024)
		}
	}
	return newStones
}

func readStones(filename string) (stones []int) {
	b, _ := os.ReadFile(filename)
	for _, s := range strings.Fields(string(b)) {
		n, _ := strconv.Atoi(s)
		stones = append(stones, n)
	}
	return
}
