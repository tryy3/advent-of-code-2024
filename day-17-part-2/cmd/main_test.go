package main

import (
	"strings"
	"testing"

	"github.com/tryy3/advent-of-code/day-8-part-1/process"
)

func Test_run_adv(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "simple case",
			input: `Register A: 729
Register B: 0
Register C: 0

Program: 0,2`,
			expected: 182,
		},
		{
			name: "simple case",
			input: `Register A: 729
Register B: 0
Register C: 0

Program: 0,3`,
			expected: 91,
		},
		{
			name: "read from register B",
			input: `Register A: 729
Register B: 7
Register C: 0

Program: 0,5`,
			expected: 5,
		},
		{
			name: "read from register B then C",
			input: `Register A: 729
Register B: 1
Register C: 2

Program: 0,5,0,6`,
			expected: 91,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			processor := process.NewProcessor()
			processor.LoadFromReader(strings.NewReader(tc.input))
			for {
				ok := processor.ReadInstruction()
				if !ok {
					break
				}
			}

			if processor.GetRegisterA() != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, processor.GetRegisterA())
			}
		})
	}
}

func Test_run_bxl(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "simple case",
			input: `Register A: 729
Register B: 5
Register C: 0

Program: 1,3`,
			expected: 6,
		},
		{
			name: "multiple reads",
			input: `Register A: 729
Register B: 63
Register C: 0

Program: 1,7324,1,32,1,124`, // 7331, 7299, 7423
			expected: 7423,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			processor := process.NewProcessor()
			processor.LoadFromReader(strings.NewReader(tc.input))
			for {
				ok := processor.ReadInstruction()
				if !ok {
					break
				}
			}

			if processor.GetRegisterB() != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, processor.GetRegisterB())
			}
		})
	}
}

func Test_run_bst(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "simple case",
			input: `Register A: 729
Register B: 0
Register C: 0

Program: 2,2`,
			expected: 2,
		},
		{
			name: "register B",
			input: `Register A: 729
Register B: 44
Register C: 0

Program: 2,5`,
			expected: 4,
		},
		{
			name: "multiple reads",
			input: `Register A: 6
Register B: 0
Register C: 327

Program: 2,4,2,6`,
			expected: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			processor := process.NewProcessor()
			processor.LoadFromReader(strings.NewReader(tc.input))
			for {
				ok := processor.ReadInstruction()
				if !ok {
					break
				}
			}

			if processor.GetRegisterB() != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, processor.GetRegisterB())
			}
		})
	}
}

func Test_run_jnz(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "simple case",
			input: `Register A: 729
Register B: 0
Register C: 0

Program: 3,5`,
			expected: 5,
		},
		{
			name: "jump",
			input: `Register A: 729
Register B: 0
Register C: 0

Program: 3,4,3,7,3,2`,
			expected: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			processor := process.NewProcessor()
			processor.LoadFromReader(strings.NewReader(tc.input))

			for {
				ok := processor.ReadInstruction()
				if !ok {
					break
				}
			}

			if processor.GetInstructionPointer() != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, processor.GetInstructionPointer())
			}
		})
	}
}
