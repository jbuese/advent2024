package day7

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("--- Day 7: Bridge Repair ---")
	fmt.Printf("C1: %d\n", c1("../data/7.txt"))
	fmt.Printf("C2: %d\n", c2("../data/7.txt"))
}

func c1(filename string) (sum int) {
	calibrations := readCalibrations(filename)
	for _, calibration := range calibrations {
		if findSolution(calibration, calibration.Operands[0], "") {
			sum += calibration.TestValue
		}
	}
	return
}

func c2(filename string) (sum int) {
	calibrations := readCalibrations(filename)
	for _, calibration := range calibrations {
		if findSolutionConcat(calibration, calibration.Operands[0], "") {
			sum += calibration.TestValue
		}
	}
	return
}

type Calibration struct {
	TestValue int
	Operands  []int
	Operators []string
}

// probably the messiest backtracking algo you have ever seen
func findSolution(calibration Calibration, currentResult int, failedOperator string) bool {
	if len(calibration.Operators) == 0 && failedOperator == "*" {
		return false
	}

	if len(calibration.Operators)+1 == len(calibration.Operands) || failedOperator == "*" {
		if currentResult == calibration.TestValue && failedOperator == "" {
			return true
		}
		failedOperator := calibration.Operators[len(calibration.Operators)-1]
		if failedOperator == "+" {
			currentResult -= calibration.Operands[len(calibration.Operators)]
		} else {
			currentResult /= calibration.Operands[len(calibration.Operators)]
		}
		calibration.Operators = calibration.Operators[:len(calibration.Operators)-1]
		return findSolution(calibration, currentResult, failedOperator)
	} else if failedOperator == "+" {
		currentResult *= calibration.Operands[len(calibration.Operators)+1]
		calibration.Operators = append(calibration.Operators, "*")
		return findSolution(calibration, currentResult, "")
	} else if failedOperator == "" {
		currentResult += calibration.Operands[len(calibration.Operators)+1]
		calibration.Operators = append(calibration.Operators, "+")
		return findSolution(calibration, currentResult, "")
	}

	return false
}

// could've incorporated into the 1st func but no time
func findSolutionConcat(calibration Calibration, currentResult int, failedOperator string) bool {
	if len(calibration.Operators) == 0 && failedOperator == "*" {
		return false
	}

	if len(calibration.Operators)+1 == len(calibration.Operands) || failedOperator == "*" {
		if currentResult == calibration.TestValue && failedOperator == "" {
			return true
		}
		failedOperator := calibration.Operators[len(calibration.Operators)-1]
		if failedOperator == "+" {
			currentResult -= calibration.Operands[len(calibration.Operators)]
		} else if failedOperator == "*" {
			currentResult /= calibration.Operands[len(calibration.Operators)]
		} else if failedOperator == "||" {
			substractLen := len(strconv.Itoa(calibration.Operands[len(calibration.Operators)]))
			currentResultStr := strconv.Itoa(currentResult)
			currentResult = atoi(currentResultStr[:len(currentResultStr)-substractLen])
		}
		calibration.Operators = calibration.Operators[:len(calibration.Operators)-1]
		return findSolutionConcat(calibration, currentResult, failedOperator)
	} else if failedOperator == "+" {
		currentResult = atoi(fmt.Sprintf("%d%d", currentResult, calibration.Operands[len(calibration.Operators)+1]))
		calibration.Operators = append(calibration.Operators, "||")
		return findSolutionConcat(calibration, currentResult, "")
	} else if failedOperator == "||" {
		currentResult *= calibration.Operands[len(calibration.Operators)+1]
		calibration.Operators = append(calibration.Operators, "*")
		return findSolutionConcat(calibration, currentResult, "")
	} else if failedOperator == "" {
		currentResult += calibration.Operands[len(calibration.Operators)+1]
		calibration.Operators = append(calibration.Operators, "+")
		return findSolutionConcat(calibration, currentResult, "")
	}

	return false
}

func readCalibrations(filename string) (calibrations []Calibration) {
	bytes, _ := os.ReadFile(filename)
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		split := strings.Split(line, ": ")
		testValue, _ := strconv.Atoi(split[0])
		operands := make([]int, 0)
		for _, operand := range strings.Split(split[1], " ") {
			operands = append(operands, atoi(operand))
		}
		calibrations = append(calibrations, Calibration{testValue, operands, nil})
	}

	return
}

func atoi(s string) int { n, _ := strconv.Atoi(s); return n }
