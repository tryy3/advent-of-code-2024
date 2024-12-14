package tests

import (
	"strings"
	"testing"

	"github.com/tryy3/advent-of-code/day-7-part-1/calculation"
)

type testReadRow struct {
	operators []calculation.Operations
	values    []int
	result    int
}

func TestReadFile(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []testReadRow
	}{
		{
			name: "Read File Test",
			input: `190: 10 19
3267: 81 40 27`,
			expected: []testReadRow{
				{
					operators: []calculation.Operations{calculation.Addition},
					values:    []int{10, 19},
					result:    190,
				},
				{
					operators: []calculation.Operations{calculation.Addition, calculation.Addition},
					values:    []int{81, 40, 27},
					result:    3267,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			calculator, err := calculation.LoadCalculatorFromReader(reader)
			if err != nil {
				t.Errorf("error loading calculator: %v", err)
			}

			if len(calculator.Calculation) != len(tc.expected) {
				t.Errorf("expected %v rows, got %v", len(tc.expected), len(calculator.Calculation))
			}

			for i, testRow := range tc.expected {
				calculatedRow := calculator.Calculation[i]

				if len(calculatedRow.Operators) != len(testRow.operators) {
					t.Errorf("expected %v operators, got %v", len(testRow.operators), len(calculatedRow.Operators))
				}
				if len(calculatedRow.Values) != len(testRow.values) {
					t.Errorf("expected %v values, got %v", len(testRow.values), len(calculatedRow.Values))
				}
				if calculatedRow.Result != testRow.result {
					t.Errorf("expected %v, got %v", testRow.result, calculatedRow.Result)
				}

				for j, testOperator := range testRow.operators {
					if calculatedRow.Operators[j].Operation != testOperator {
						t.Errorf("expected operator %v, got %v", testOperator, calculatedRow.Operators[j].Operation)
					}
				}

				for j, testValue := range testRow.values {
					if calculatedRow.Values[j].GetValue() != testValue {
						t.Errorf("expected value %v, got %v", testValue, calculatedRow.Values[j].GetValue())
					}
				}
			}
		})
	}
}
