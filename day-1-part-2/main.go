package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var listA = []int{}
var listB = []int{}

func main() {
	readFile()
	calcuateSimilarity()
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

func calcuateSimilarity() {
	similarity := 0
	for _, aValue := range listA {
		count := 0
		for _, bValue := range listB {
			if aValue == bValue {
				count++
			}
		}
		fmt.Printf("%#v - %#v\n", aValue, count)
		similarity += aValue * count
	}

	fmt.Println(similarity)
}
