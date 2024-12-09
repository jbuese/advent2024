package day9

import (
	"fmt"
	"os"
)

func Run() {
	fmt.Println("--- Day 9: Disk Fragmenter ---")
	fmt.Printf("C1: %d\n", c1("../data/9.txt"))
	fmt.Printf("C2: %d\n", c2("../data/9.txt"))
}

func c1(filename string) (sum int) {
	disk := readDisk(filename)
	lastValidPos := len(disk) - 1
	for lastValidPos >= 0 && disk[lastValidPos] == -1 {
		lastValidPos--
	}

	for i := 0; i <= lastValidPos; i++ {
		if disk[i] == -1 {
			for lastValidPos > i && disk[lastValidPos] == -1 {
				lastValidPos--
			}
			if lastValidPos <= i {
				break
			}
			disk[i] = disk[lastValidPos]
			disk[lastValidPos] = -1
			lastValidPos--
		}
		if disk[i] != -1 {
			sum += i * disk[i]
		}
	}

	return sum
}

func c2(filename string) (count int) {
	return
}

func readDisk(filename string) (res []int) {
	bytes, _ := os.ReadFile(filename)
	numbers := make([]int, len(bytes))
	for i, b := range bytes {
		numbers[i] = int(b - '0')
	}

	for i := 0; i < len(numbers); i++ {
		if i%2 == 0 {
			for j := 0; j < numbers[i]; j++ {
				res = append(res, i/2)
			}
		} else {
			for j := 0; j < numbers[i]; j++ {
				res = append(res, -1)
			}
		}
	}

	return
}

func findLastFileIndex(disk []int) int {
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != -1 {
			return i
		}
	}
	return -1
}
