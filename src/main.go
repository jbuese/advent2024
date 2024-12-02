package main

import (
	day1 "advent2024/src/d1"
	day2 "advent2024/src/d2"
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
	default:
		fmt.Printf("Day %d not implemented yet\n", day)
	}
}
