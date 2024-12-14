package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// var listA []int = []int{3, 4, 2, 1, 3, 3}
// var listB []int = []int{4, 3, 5, 3, 9, 3}

var listA = []int{}
var listB = []int{}

func main() {
	readFile()
	calculate()
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
		// Print each line
		var text = strings.Split(scanner.Text(), "   ")

		numA, err := strconv.Atoi(text[0])
		if err != nil {
			fmt.Println("Error parsing string to int:", err)
			return
		}

		numB, err := strconv.Atoi(text[1])
		if err != nil {
			fmt.Println("Error parsing string to int:", err)
			return
		}
		listA = append(listA, numA)
		listB = append(listB, numB)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func calculate() {
	// var paired = []int{}
	sort.Slice(listA, func(i, j int) bool {
		return listA[i] < listA[j]
	})
	sort.Slice(listB, func(i, j int) bool {
		return listB[i] < listB[j]
	})

	var sum = 0

	for i := 0; i < len(listA); i++ {
		var aValue = listA[i]
		var bValue = listB[i]
		var diff = 0
		if aValue > bValue {
			diff = aValue - bValue
		} else {
			diff = bValue - aValue
		}
		sum += diff
	}

	fmt.Printf("%#v\n", sum)
}
