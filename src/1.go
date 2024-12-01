package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("--- Day 1: Historian Hysteria ---")
	fmt.Printf("C1: %d\n", c1("../data/1.txt"))
	fmt.Printf("C2: %d\n", c2("../data/1.txt"))
}

func c1(filename string) (sum int) {
	nums1, nums2 := splitScan(filename)
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i := range nums1 {
		sum += abs(nums2[i] - nums1[i])
	}
	return
}

func c2(filename string) (similarity int) {
	nums1, nums2 := splitScan(filename)
	for i := range nums1 {
		similarity += nums1[i] * countOccurrences(nums2, nums1[i])
	}
	return
}

func splitScan(filename string) ([]int, []int) {
	bytes, _ := os.ReadFile(filename)
	var n1, n2 []int
	for _, line := range strings.Split(string(bytes), "\n") {
		if p := strings.Split(line, "   "); len(p) > 1 {
			n1, n2 = append(n1, atoi(p[0])), append(n2, atoi(p[1]))
		}
	}
	return n1, n2
}

func atoi(s string) int { n, _ := strconv.Atoi(s); return n }

func countOccurrences(list []int, value int) (count int) {
	for _, v := range list {
		if v == value {
			count++
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
