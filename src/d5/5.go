package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("--- Day 5: Print Queue ---")
	fmt.Printf("C1: %d\n", c1("../data/5.txt"))
	fmt.Printf("C2: %d\n", c2("../data/5.txt"))
}

func c1(filename string) (sum int) {
	rules, sequences := readInput(filename)

	for _, seq := range sequences {
		valid := true
	seqLoop:
		for i := 0; i < len(seq); i++ {
			for _, rule := range rules {
				if seq[i] == rule[1] {
					for j := i + 1; j < len(seq); j++ {
						if seq[j] == rule[0] {
							valid = false
							break seqLoop
						}
					}
				}
			}
		}
		if valid {
			sum += seq[len(seq)/2]
		}
	}
	return
}

func c2(filename string) (sum int) {
	rules, sequences := readInput(filename)

	for _, seq := range sequences {
		fixed := false
	checkSequence:
		for i := 0; i < len(seq); i++ {
			for _, rule := range rules {
				if seq[i] == rule[1] {
					for j := i + 1; j < len(seq); j++ {
						if seq[j] == rule[0] {
							seq[i], seq[j] = seq[j], seq[i]
							fixed = true
							i = -1 //reset
							continue checkSequence
						}
					}
				}
			}
		}
		if fixed {
			sum += seq[len(seq)/2]
		}
	}
	return
}

func readInput(filename string) (rules [][]int, sequences [][]int) {
	bytes, _ := os.ReadFile(filename)
	parts := strings.SplitN(strings.ToUpper(string(bytes)), "\n\n", 2)
	ruleLines := strings.Split(parts[0], "\n")
	for _, line := range ruleLines {
		split := strings.Split(line, "|")
		rules = append(rules, []int{atoi(split[0]), atoi(split[1])})
	}

	pageLines := strings.Split(parts[1], "\n")
	for _, line := range pageLines {
		if p := strings.Split(line, ","); len(p) > 1 {
			var page []int
			for _, v := range p {
				page = append(page, atoi(v))
			}
			sequences = append(sequences, page)
		}
	}

	return
}

func atoi(s string) int { n, _ := strconv.Atoi(s); return n }
