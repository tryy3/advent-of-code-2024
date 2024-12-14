package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Rule struct {
	First  int
	Second int
}

type Printer struct {
	Rules []Rule
	Paper [][]int
}

func main() {
	start := time.Now()
	startProcessing()
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}

func startProcessing() {
	f, err := openFile()
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	printer := NewPrinter()
	err = printer.ReadFile(f)
	if err != nil {
		log.Println("Error reading file:", err)
		return
	}

	fmt.Println(printer.Calculate())
	fmt.Println(printer.CalculateInCorrectOrder())
}

func openFile() (*os.File, error) {
	// Open the file
	// file, err := os.Open("test_input.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	return file, nil
}

func NewPrinter() *Printer {
	return &Printer{
		Paper: [][]int{},
		Rules: []Rule{},
	}
}

func (p *Printer) ReadFile(reader io.Reader) error {
	ruleRegex := regexp.MustCompile(`(\d+)\|(\d+)`)
	pageRegex := regexp.MustCompile(`(\d+),?`)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		rules := ruleRegex.FindStringSubmatch(line)

		if len(rules) == 3 {
			firstRule, err := strconv.Atoi(rules[1])
			if err != nil {
				return fmt.Errorf("error parsing first rule: %v", err)
			}
			secondRule, err := strconv.Atoi(rules[2])
			if err != nil {
				return fmt.Errorf("error parsing second rule: %v", err)
			}
			p.Rules = append(p.Rules, Rule{First: firstRule, Second: secondRule})
		} else {
			pages := pageRegex.FindAllStringSubmatch(line, -1)
			parsedPages := []int{}
			for _, page := range pages {
				parsedPage, err := strconv.Atoi(page[1])
				if err != nil {
					return fmt.Errorf("error parsing page: %v", err)
				}
				parsedPages = append(parsedPages, parsedPage)
			}
			p.Paper = append(p.Paper, parsedPages)
		}
	}
	return scanner.Err()
}

func (p *Printer) IsInOrder(pages []int) bool {
	for _, rule := range p.Rules {
		firstMatchedIndex := -1
		secondMatchedIndex := -1
		for i, page := range pages {
			if page == rule.First {
				firstMatchedIndex = i
			}
			if page == rule.Second {
				secondMatchedIndex = i
			}
		}

		if firstMatchedIndex == -1 || secondMatchedIndex == -1 {
			continue
		}

		if firstMatchedIndex > secondMatchedIndex {
			return false
		}
	}
	return true
}

func (p *Printer) Reorder(pages []int) []int {
	isInOrder := false
	for {
		if isInOrder {
			break
		}

		for _, rule := range p.Rules {
			firstMatchedIndex := -1
			secondMatchedIndex := -1
			for i, page := range pages {
				if page == rule.First {
					firstMatchedIndex = i
				}
				if page == rule.Second {
					secondMatchedIndex = i
				}
			}

			if firstMatchedIndex == -1 || secondMatchedIndex == -1 {
				continue
			}

			if firstMatchedIndex > secondMatchedIndex {
				pages[firstMatchedIndex], pages[secondMatchedIndex] = pages[secondMatchedIndex], pages[firstMatchedIndex]
			}
		}

		isInOrder = p.IsInOrder(pages)
	}
	return pages
}

func (p *Printer) Calculate() int {
	total := 0
	for _, page := range p.Paper {
		if p.IsInOrder(page) {
			total += page[len(page)/2]
		}
	}
	return total
}

func (p *Printer) CalculateInCorrectOrder() int {
	total := 0
	for _, page := range p.Paper {
		if !p.IsInOrder(page) {
			correctionOrderPage := p.Reorder(page)
			total += correctionOrderPage[len(correctionOrderPage)/2]
			// total += page[len(page)/2]
		}
	}
	return total
}
