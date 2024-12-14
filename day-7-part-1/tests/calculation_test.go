package tests

import (
	"testing"

	"github.com/tryy3/advent-of-code/day-7-part-1/calculation"
)

type testCalculationValues struct {
	operators []calculation.Operations
	values    []int
	result    int
	expected  bool
}

func TestCalculationAttempt(t *testing.T) {
	testCases := []struct {
		name string
		test testCalculationValues
	}{
		{
			name: "Test Successful Addition calculation",
			test: testCalculationValues{
				operators: []calculation.Operations{calculation.Addition},
				values:    []int{10, 10},
				result:    20,
				expected:  true,
			},
		},
		{
			name: "Test Successful Multiplication calculation",
			test: testCalculationValues{
				operators: []calculation.Operations{calculation.Multiplication},
				values:    []int{10, 10},
				result:    100,
				expected:  true,
			},
		},
		{
			name: "Test Failure Addition calculation",
			test: testCalculationValues{
				operators: []calculation.Operations{calculation.Addition},
				values:    []int{10, 10},
				result:    -1,
				expected:  false,
			},
		},
		{
			name: "Test Failure Multiplication calculation",
			test: testCalculationValues{
				operators: []calculation.Operations{calculation.Multiplication},
				values:    []int{10, 10},
				result:    -1,
				expected:  false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup the calculation
			calc := calculation.NewCalculation()
			for _, operator := range tc.test.operators {
				calc.Operators = append(calc.Operators, calculation.NewOperator(operator))
			}

			for _, value := range tc.test.values {
				calc.Values = append(calc.Values, calculation.NewValue(value))
			}

			calc.Result = tc.test.result

			// Attempt the calculation
			if calc.AttemptCalculate() != tc.test.expected {
				t.Errorf("Expected calculation to be %t, got %t", tc.test.expected, calc.AttemptCalculate())
			}

			// Sanity check the last value
			if tc.test.expected {
				if calc.Values[len(calc.Values)-1].GetValue() != tc.test.result {
					t.Errorf("Expected last value to be %d, got %d", tc.test.result, calc.Values[len(calc.Values)-1].GetValue())
				}
			}
		})
	}
}
