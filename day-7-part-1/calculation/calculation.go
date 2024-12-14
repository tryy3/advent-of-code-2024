package calculation

import (
	"fmt"
	"strings"
)

func generateOperatorCombinations(operators []*Operator, numValues int) [][]*Operator {
	// fmt.Printf("Operators: %v\n", operators)
	// Calculate the number of slots for operators
	numSlots := numValues - 1

	// A slice to store all combinations
	var combinations [][]*Operator

	// Helper function for recursion
	var backtrack func(currentCombination []*Operator, depth int)
	backtrack = func(currentCombination []*Operator, depth int) {
		// If we've filled all operator slots, add the combination to results
		if depth == numSlots {
			// Make a copy of the current combination and append to results
			combinationCopy := append([]*Operator{}, currentCombination...)
			combinations = append(combinations, combinationCopy)
			return
		}

		// Try each operator in the current slot
		for _, operator := range operators {
			// Add the operator
			currentCombination = append(currentCombination, operator)
			// Recurse to the next depth
			backtrack(currentCombination, depth+1)
			// Backtrack by removing the last added operator
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}

	// Start the recursion with an empty combination
	backtrack([]*Operator{}, 0)

	return combinations
}

type CalculationRow struct {
	CalculatedOperators [][]*Operator
	Values              []*Value
	Result              int

	currentOperatorIndex int
	// lastOperationFlip    int
	isDone bool
}

func NewCalculation() *CalculationRow {
	return &CalculationRow{
		CalculatedOperators:  [][]*Operator{},
		Values:               []*Value{},
		currentOperatorIndex: 0,
		// lastOperationFlip:    -1,
	}
}

func (c *CalculationRow) AttemptCalculate() bool {
	operators := c.CalculatedOperators[c.currentOperatorIndex]
	// if c.currentOperatorIndex == 2 {
	// 	fmt.Printf("AttemptCalculate: %d\n", c.currentOperatorIndex)
	// 	fmt.Printf("Operator count: %d\n", len(c.CalculatedOperators))
	// 	fmt.Printf("Operators: %v\n", operators)
	// }

	for i := 0; i < len(c.Values)-1; i++ {
		firstValue := c.Values[i]
		secondValue := c.Values[i+1]

		secondValue.Calculate(firstValue, operators[i])
	}

	// if c.currentOperatorIndex == 2 {
	// 	spew.Dump(c.Values)
	// }

	result := c.IsCorrect()
	// fmt.Printf("Result: %t\n", result)
	if result {
		c.isDone = true
	}

	return result
}

func (c *CalculationRow) AllOperatorsIsMultiplication() bool {
	for _, operator := range c.CalculatedOperators[c.currentOperatorIndex] {
		if operator.Operation != Multiplication {
			return false
		}
	}
	return true
}

func (c *CalculationRow) IsDone() bool {
	return c.isDone
}

func (c *CalculationRow) IsCorrect() bool {
	return c.Values[len(c.Values)-1].GetValue() == c.Result
}

func (c *CalculationRow) TrySwitchNextOperator() bool {
	c.currentOperatorIndex++
	if c.currentOperatorIndex >= len(c.CalculatedOperators) {
		c.isDone = true
		return false
	}
	c.Reset()
	return true

	// if c.currentOperatorIndex == -1 {
	// 	c.currentOperatorIndex = 0
	// 	c.Operators[c.currentOperatorIndex].SwitchOperation()
	// 	return true
	// }

	// c.lastOperationFlip++
	// if c.currentOperatorIndex >= len(c.Operators) && c.lastOperationFlip >= len(c.Operators) {
	// 	c.isDone = true
	// 	return false
	// }

	// if c.lastOperationFlip == c.currentOperatorIndex {
	// 	c.lastOperationFlip++
	// }

	// if c.lastOperationFlip >= len(c.Operators) {
	// 	c.Reset()
	// 	c.currentOperatorIndex++
	// 	if c.currentOperatorIndex >= len(c.Operators) {
	// 		c.isDone = true
	// 		return false
	// 	}
	// 	c.Operators[c.currentOperatorIndex].SwitchOperation()
	// 	return true
	// }
	// c.Operators[c.lastOperationFlip].SwitchOperation()
	// return true
}

func (c *CalculationRow) ResetValues() {
	for _, value := range c.Values {
		value.Reset()
	}
}

// func (c *CalculationRow) ResetOperators() {
// 	for _, operator := range c.Operators {
// 		operator.Reset()
// 	}
// }

func (c *CalculationRow) Reset() {
	c.ResetValues()
	// c.ResetOperators()
	// c.lastOperationFlip = 0
}

func (c *CalculationRow) String() string {
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString(fmt.Sprintf("%d: ", c.Result))
	for i, value := range c.Values {
		stringBuilder.WriteString(fmt.Sprintf("%d ", value.GetOriginalValue()))
		if i < len(c.Values)-1 {
			stringBuilder.WriteString(fmt.Sprintf("%s ", c.CalculatedOperators[c.currentOperatorIndex][i].String()))
		}
	}
	stringBuilder.WriteString("= ")
	stringBuilder.WriteString(fmt.Sprintf("%d", c.Values[len(c.Values)-1].GetValue()))
	return stringBuilder.String()
}
