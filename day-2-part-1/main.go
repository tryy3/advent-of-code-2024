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
		var firstValue = report[0]
		var lastValue = report[1]
		isIncreasing := firstValue < lastValue
		isSafe := isLevelSafe(isIncreasing, report)
		if isSafe {
			fmt.Printf("%#v - Safe\n", report)
			countSafe++
		} else {
			fmt.Printf("%#v - Unsafe\n", report)
		}
	}
	fmt.Println(countSafe)
}

func isLevelSafe(isIncreasing bool, levels []int) bool {
	prevLevel := -1
	for _, level := range levels {
		if prevLevel == -1 {
			prevLevel = level
			continue
		}

		if isIncreasing {
			if prevLevel > level {
				return false
			}
			diff := level - prevLevel
			if diff < 1 || diff > 3 {
				return false
			}
		} else {
			if prevLevel < level {
				return false
			}
			diff := prevLevel - level
			if diff < 1 || diff > 3 {
				return false
			}
		}
		prevLevel = level
	}
	return true
}
