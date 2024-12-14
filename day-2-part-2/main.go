package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var listReports = [][]int{}

func main() {
	readFile()
	checkSafeLevel()
}

func readFile() {
	// Open the file
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var listLevels = []int{}
		text := scanner.Text()
		splitText := strings.Split(text, " ")
		for _, val := range splitText {
			parsedValue, err := strconv.Atoi(val)
			if err != nil {
				log.Fatalf("Error parsing string to int: %v", err)
			}
			listLevels = append(listLevels, parsedValue)
		}
		listReports = append(listReports, listLevels)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func checkSafeLevel() {
	countSafe := 0
	for _, report := range listReports {
		if checkLevel(report) {
			countSafe++
		}
	}
	fmt.Println(countSafe)
	// checkLevel(listReports[7])
}

// RemoveIndex creates a copy of the slice and removes the specified index.
func RemoveIndex(slice []int, index int) ([]int, error) {
	// Validate index
	if index < 0 || index >= len(slice) {
		return nil, fmt.Errorf("index out of bounds: %d", index)
	}

	// Create a new slice and remove the specified index
	newSlice := append([]int{}, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)

	return newSlice, nil
}

func checkLevel(report []int) bool {
	reportCopy := make([]int, len(report))
	copy(reportCopy, report)
	isIncreasing := isIncreasingLevel(reportCopy)
	isSafe, index, rules := isLevelSafe(isIncreasing, reportCopy)
	// Check once more by removing the false index
	if !isSafe {
		reportCopy, err := RemoveIndex(report, index)
		if err != nil {
			log.Fatalf("error removing index: %v", err)
		}
		isSafe, _, _ = isLevelSafe(isIncreasing, reportCopy)
	}
	if !isSafe {
		reportCopy, err := RemoveIndex(report, 0)
		if err != nil {
			log.Fatalf("error removing index: %v", err)
		}
		isIncreasing := isIncreasingLevel(reportCopy)
		isSafe, _, _ = isLevelSafe(isIncreasing, reportCopy)
	}
	if !isSafe {
		reportCopy, err := RemoveIndex(report, 1)
		if err != nil {
			log.Fatalf("error removing index: %v", err)
		}
		isIncreasing := isIncreasingLevel(reportCopy)
		isSafe, _, _ = isLevelSafe(isIncreasing, reportCopy)
	}

	if isSafe {
		// fmt.Printf("%#v - Safe\n", report)
		return true
	} else {
		fmt.Printf("%#v - Unsafe - %#v\n", report, rules)
	}
	return false
}

func isIncreasingLevel(levels []int) bool {
	var firstValue = levels[0]
	var lastValue = levels[1]
	return firstValue < lastValue
}

func isLevelSafe(isIncreasing bool, levels []int) (bool, int, []string) {
	prevLevel := -1
	rules := []string{}
	index := -1

	for i, level := range levels {
		if prevLevel == -1 {
			prevLevel = level
			continue
		}

		isFailed, rule := checkRuleFails(isIncreasing, prevLevel, level)
		if isFailed {
			if index == -1 {
				index = i
			}
			rules = append(rules, rule)
		}
		prevLevel = level
	}
	return (index) == -1, index, rules
}

func checkRuleFails(isIncreasing bool, val1, val2 int) (bool, string) {
	if isIncreasing {
		if val1 > val2 {
			return true, "IsIncrease - fail"
		}
		diff := val2 - val1
		if diff < 1 {
			return true, "Low number"
		}
		if diff > 3 {
			return true, "High number"
		}
	} else {
		if val1 < val2 {
			return true, "IsDecrease - fail"
		}
		diff := val1 - val2
		if diff < 1 {
			return true, "Low number"
		}
		if diff > 3 {
			return true, "High number"
		}
	}
	return false, ""
}
