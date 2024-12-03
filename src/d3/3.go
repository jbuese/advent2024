package day2

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("--- Day 3: Mull It Over ---")
	fmt.Printf("C1: %d\n", c1("../data/3.txt"))
	fmt.Printf("C2: %d\n", c2("../data/3.txt"))
}

func c1(filename string) (mult int) {
	memory := readMemory(filename)
	instr, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)

	for _, match := range instr.FindAllString(memory, -1) {
		parts := strings.Split(match, ",")
		mult += atoi(parts[0][4:]) * atoi(parts[1][:len(parts[1])-1])
	}

	return
}

func c2(filename string) (mult int) {
	memory := readMemory(filename)
	mulInstr, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	doInstr, _ := regexp.Compile(`do\(\)`)
	dontInstr, _ := regexp.Compile(`don't\(\)`)

	mulIx := mulInstr.FindAllStringIndex(memory, -1)
	doIx := doInstr.FindAllStringIndex(memory, -1)
	dontIx := dontInstr.FindAllStringIndex(memory, -1)

	for _, mulMatch := range mulIx {
		mulStart := mulMatch[0]

		lastDo := -1
		for _, do := range doIx {
			if do[0] < mulStart {
				lastDo = do[0]
			} else {
				break
			}
		}

		lastDont := -1
		for _, dont := range dontIx {
			if dont[0] < mulStart {
				lastDont = dont[0]
			} else {
				break
			}
		}

		if lastDo > lastDont || (lastDo == -1 && lastDont == -1) {
			parts := strings.Split(memory[mulMatch[0]:mulMatch[1]], ",")
			mult += atoi(parts[0][4:]) * atoi(parts[1][:len(parts[1])-1])
		}
	}

	return
}

func readMemory(filename string) string {
	bytes, _ := os.ReadFile(filename)
	return strings.TrimSpace(string(bytes))
}

func atoi(s string) int { n, _ := strconv.Atoi(s); return n }
