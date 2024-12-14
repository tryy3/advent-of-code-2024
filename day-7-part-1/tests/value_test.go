package tests

import (
	"testing"

	"github.com/tryy3/advent-of-code/day-7-part-1/calculation"
)

type testCalculatedValue struct {
	firstValue      int
	secondValue     int
	calculatedValue int
	operator        calculation.Operations
}

func TestValueCalculate(t *testing.T) {
	testCases := []struct {
		name string
		test testCalculatedValue
	}{
		{
			name: "Test Addition value",
			test: testCalculatedValue{
				firstValue:      10,
				secondValue:     10,
				calculatedValue: 20,
				operator:        calculation.Addition,
			},
		},
		{
			name: "Test Multiplation value",
			test: testCalculatedValue{
				firstValue:      10,
				secondValue:     10,
				calculatedValue: 100,
				operator:        calculation.Multiplication,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup the test values
			firstValue := calculation.NewValue(tc.test.firstValue)
			secondValue := calculation.NewValue(tc.test.secondValue)
			operator := calculation.NewOperator(tc.test.operator)
			secondValue.Calculate(firstValue, operator)

			if firstValue.GetValue() != tc.test.firstValue {
				t.Errorf("Expected first value to be %d, got %d", tc.test.firstValue, firstValue.GetValue())
			}

			if secondValue.GetValue() != tc.test.calculatedValue {
				t.Errorf("Expected second value to be %d, got %d", tc.test.calculatedValue, secondValue.GetValue())
			}
		})
	}
}
