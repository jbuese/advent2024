package day1

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	fmt.Println("--- Day 4: Ceres Search ---")
	fmt.Printf("C1: %d\n", c1("../data/4.txt"))
	fmt.Printf("C2: %d\n", c2("../data/4.txt"))
}

func c1(filename string) (sum int) {
	matrix := readWordSearch(filename)

	rows, cols := len(matrix), len(matrix[0])

	// horizontal
	for row := 0; row < rows; row++ {
		for cell := 3; cell < cols; cell++ {
			candidate := string(matrix[row][cell-3 : cell+1])
			if checkXMAS(candidate) {
				sum++
			}
		}
	}

	// vert
	for col := 0; col < cols; col++ {
		for row := 3; row < rows; row++ {
			candidate := string([]rune{
				matrix[row-3][col],
				matrix[row-2][col],
				matrix[row-1][col],
				matrix[row][col],
			})
			if checkXMAS(candidate) {
				sum++
			}
		}
	}

	// diagonal right
	for row := 0; row < rows-3; row++ {
		for col := 0; col < cols-3; col++ {
			candidate := string([]rune{
				matrix[row][col],
				matrix[row+1][col+1],
				matrix[row+2][col+2],
				matrix[row+3][col+3],
			})
			if checkXMAS(candidate) {
				sum++
			}
		}
	}

	// diagonal left
	for row := 0; row < rows-3; row++ {
		for col := 3; col < cols; col++ {
			candidate := string([]rune{
				matrix[row][col],
				matrix[row+1][col-1],
				matrix[row+2][col-2],
				matrix[row+3][col-3],
			})
			if checkXMAS(candidate) {
				sum++
			}
		}
	}

	return
}

func c2(filename string) (count int) {
	matrix := readWordSearch(filename)

	rows, cols := len(matrix), len(matrix[0])
	for row := 0; row < rows-2; row++ {
		for cell := 0; cell < cols-2; cell++ {
			candidate1 := string([]rune{
				matrix[row][cell],
				matrix[row+1][cell+1],
				matrix[row+2][cell+2],
			})
			candidate2 := string([]rune{
				matrix[row][cell+2],
				matrix[row+1][cell+1],
				matrix[row+2][cell],
			})
			if checkMAS(candidate1) && checkMAS(candidate2) {
				count++
			}
		}
	}

	return
}

func readWordSearch(filename string) (matrix [][]rune) {
	bytes, _ := os.ReadFile(filename)
	lines := strings.Split(strings.ToUpper(string(bytes)), "\n")
	matrix = make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	return
}

func checkXMAS(word string) bool {
	return word == "XMAS" || word == "SAMX"
}

func checkMAS(word string) bool {
	return word == "MAS" || word == "SAM"
}
