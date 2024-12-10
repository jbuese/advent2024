package main

import (
	day1 "advent2024/src/d1"
	day10 "advent2024/src/d10"
	day2 "advent2024/src/d2"
	day3 "advent2024/src/d3"
	day4 "advent2024/src/d4"
	day5 "advent2024/src/d5"
	day6 "advent2024/src/d6"
	day7 "advent2024/src/d7"
	day9 "advent2024/src/d9"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify which day to run (1-25)")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please provide a valid day number")
		return
	}

	switch day {
	case 1:
		day1.Run()
	case 2:
		day2.Run()
	case 3:
		day3.Run()
	case 4:
		day4.Run()
	case 5:
		day5.Run()
	case 6:
		day6.Run()
	case 7:
		day7.Run()
	case 9:
		day9.Run()
	case 10:
		day10.Run()
	default:
		fmt.Printf("Day %d not implemented yet\n", day)
	}
}
