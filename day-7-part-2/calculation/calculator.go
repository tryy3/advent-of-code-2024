package calculation

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Calculator struct {
	Calculation []*CalculationRow
}

func NewCalculator() *Calculator {
	return &Calculator{
		Calculation: []*CalculationRow{},
	}
}

func LoadCalculatorFromFile(path string) (*Calculator, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	return LoadCalculatorFromReader(file)
}

func LoadCalculatorFromReader(reader io.Reader) (*Calculator, error) {
	lineRegex := regexp.MustCompile(`^(\d+):\s*((?:\d+\s*)+)$`)

	calculator := NewCalculator()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		calculation := NewCalculation()
		line := scanner.Text()

		matches := lineRegex.FindStringSubmatch(line)
		if matches == nil {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		result, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing result: %v", err)
		}

		calculation.Result = result

		values := strings.Split(matches[2], " ")
		for _, value := range values {
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("error parsing value: %v", err)
			}
			calculation.Values = append(calculation.Values, NewValue(valueInt))
		}

		// generateOperations := len(values) - 1
		// operators := []*Operator{}
		// for i := 0; i < generateOperations; i++ {
		// 	operator := NewOperator(Addition)
		// 	operators = append(operators, operator)
		// }
		operators := []*Operator{NewOperator(Addition), NewOperator(Multiplication), NewOperator(Concatenation)}
		calculation.CalculatedOperators = generateOperatorCombinations(operators, len(values))
		calculator.Calculation = append(calculator.Calculation, calculation)
	}

	return calculator, nil
}

func (c *Calculator) Calculate() int {
	allCalculated := false

	for !allCalculated {
		for _, calculation := range c.Calculation {
			if !calculation.IsDone() {
				ok := calculation.AttemptCalculate()
				// fmt.Println(calculation.String())
				if !ok {
					calculation.TrySwitchNextOperator()
					if !calculation.IsCorrect() {
						calculation.ResetValues()
					}
				}
			}
		}

		allCalculated = true
		for _, calculation := range c.Calculation {
			if !calculation.IsDone() {
				allCalculated = false
				break
			}
		}
	}

	total := 0
	for _, calculation := range c.Calculation {
		if calculation.IsCorrect() {
			total += calculation.Result
		}
	}
	return total
}
