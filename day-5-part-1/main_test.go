package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCalculate(t *testing.T) {
	printer := &Printer{
		Rules: []Rule{
			{First: 10, Second: 20},
			{First: 20, Second: 30},
		},
		Paper: [][]int{
			{10, 20, 30},
			{30, 20, 10},
			{10, 5, 30},
		},
	}

	testCases := []struct {
		name     string
		expected int
	}{
		{
			name:     "Test order is correct",
			expected: 25,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := printer.Calculate()
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestOrder(t *testing.T) {
	printer := &Printer{
		Rules: []Rule{
			{First: 10, Second: 20},
			{First: 20, Second: 30},
		},
		Paper: [][]int{
			{10, 20, 30},
			{30, 20, 10},
		},
	}

	testCases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "Test order is correct",
			input:    printer.Paper[0],
			expected: true,
		},
		{
			name:     "Test order is incorrect",
			input:    printer.Paper[1],
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := printer.IsInOrder(tc.input)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestInput(t *testing.T) {
	content := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	t.Run("Read AOC Test Input", func(t *testing.T) {
		printer := NewPrinter()
		err := printer.ReadFile(strings.NewReader(content))
		if err != nil {
			t.Errorf("error reading file: %v", err)
		}

		result := printer.Calculate()
		if result != 143 {
			t.Errorf("Calculate(): expected %v, got %v", 143, result)
		}

		result = printer.CalculateInCorrectOrder()
		if result != 123 {
			t.Errorf("CalculateInCorrectOrder(): unexpected %v, got %v", 123, result)
		}
	})
}

func TestReordering(t *testing.T) {
	content := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	printer := NewPrinter()
	err := printer.ReadFile(strings.NewReader(content))
	if err != nil {
		t.Errorf("error reading file: %v", err)
	}

	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "First reordering",
			input:    printer.Paper[3],
			expected: []int{97, 75, 47, 61, 53},
		},
		{
			name:     "Second reordering",
			input:    printer.Paper[4],
			expected: []int{61, 29, 13},
		},
		{
			name:     "Third reordering",
			input:    printer.Paper[5],
			expected: []int{97, 75, 47, 29, 13},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := printer.Reorder(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected *Printer
	}{
		{
			name: "Read File Test",
			input: `75|15
32|11

13,32
53,13,25`,
			expected: &Printer{
				Paper: [][]int{
					{13, 32},
					{53, 13, 25},
				},
				Rules: []Rule{
					{First: 75, Second: 15},
					{First: 32, Second: 11},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			printer := NewPrinter()
			err := printer.ReadFile(strings.NewReader(tc.input))
			if err != nil {
				t.Errorf("error reading file: %v", err)
			}
			if !reflect.DeepEqual(printer, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, printer)
			}
		})
	}
}
