package main

import (
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestSimpleMock(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected int
	}{
		{"simple data", "xmul(2,4)", 8},
		{"multiple wrong char data", "z(5muxmul(7,4)asdx", 28},
		{"multiple wrong cases", "z(5muxmul(7,4)asdmul(43,2mul(2,2)x", 32},
		{"weird error case", "mulmul(2,2)", 4},
		{"another weird error case", "mul(7,4)mul(32,1mul(2,2)", 32},
		{"test case part 2", "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", 48},
		{"do func", "do()x", 0},
		{"don't func", "don't()x", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock reader using strings.NewReader
			reader := strings.NewReader(tt.content)

			t.Run("file read successful", func(t *testing.T) {
				machine := NewStateMachine()
				err := machine.Start(reader)
				if err != nil {
					t.Errorf("error running readFile: %v", err)
				}

				result := machine.Calculate()
				if result != tt.expected {
					t.Errorf("Calculate() = %d; want %d; content %s", result, tt.expected, tt.content)
					t.Errorf("State: %s", spew.Sdump(machine))
				}
			})
		})
	}
}
