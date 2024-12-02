package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("--- Day 2: Red-Nosed Reports ---")
	fmt.Printf("C1: %d\n", c1("../data/2.txt"))
	fmt.Printf("C2: %d\n", c2("../data/2.txt"))
}

func c1(filename string) (count int) {
	reports := readReports(filename)

	for _, report := range reports {
		valid := true
		diff := report[0] - report[1]
		if diff < -3 || diff > 3 {
			continue
		}
		for lvl := 1; lvl < len(report)-1; lvl++ {
			newDiff := report[lvl] - report[lvl+1]
			if newDiff >= -3 && newDiff <= 3 && newDiff*diff > 0 {
				diff = newDiff
			} else {
				valid = false
				break
			}
		}
		if valid {
			count++
		}
	}
	return
}

func c2(filename string) (count int) {
	return
}

func readReports(filename string) (reports [][]int) {
	bytes, _ := os.ReadFile(filename)
	for _, line := range strings.Split(string(bytes), "\n") {
		if p := strings.Split(line, " "); len(p) > 1 {
			var report []int
			for _, v := range p {
				report = append(report, atoi(v))
			}
			reports = append(reports, report)
		}
	}
	return
}

func atoi(s string) int { n, _ := strconv.Atoi(s); return n }
