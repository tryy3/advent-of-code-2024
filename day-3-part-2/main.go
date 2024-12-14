package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Define parser states
const (
	Start = iota
	Operator
	OpenParen
	Value
	Comma
	CloseParen
)

type Operation struct {
	Operator string
	Values   []int
}

type StateMachine struct {
	LastRune         rune
	State            int
	CurrentString    string
	CurrentValue     strings.Builder
	CurrentOperation *Operation
	ParsedOperations []*Operation
	Process          bool
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		State:            Start,
		CurrentString:    "",
		CurrentValue:     strings.Builder{},
		CurrentOperation: NewOperation(),
		Process:          true,
	}
}

func NewOperation() *Operation {
	return &Operation{
		Operator: "",
		Values:   []int{},
	}
}

func main() {
	file, err := openFile()
	if err != nil {
		log.Fatalf("error executing openFile: %v", err)
	}
	defer file.Close()

	machine := NewStateMachine()
	err = machine.Start(file)
	if err != nil {
		log.Fatalf("error executing State Machine: %v", err)
	}

	calc := machine.Calculate()
	fmt.Println(calc)
}

func openFile() (*os.File, error) {
	// Open the file
	// file, err := os.Open("test_input.txt")
	file, err := os.Open("file.txt")
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	return file, nil
}

func (s *StateMachine) Start(r io.Reader) error {
	// Create a buffered reader
	reader := bufio.NewReader(r)

	for {
		// Read a single UTF-8 encoded rune (character)
		r, _, err := reader.ReadRune()
		if err != nil {
			// Break at EOF
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("error reading file: %v", err)
		}

		err = s.NextRune(r)
		if err != nil {
			return fmt.Errorf("error parsing next rune: %v", err)
		}
	}
	return nil
}

func (s *StateMachine) Calculate() int {
	total := 0
	for _, operation := range s.ParsedOperations {
		if strings.HasPrefix(operation.Operator, "mul") {
			total += operation.Values[0] * operation.Values[1]
		}
	}
	return total
}

func (s *StateMachine) NextRune(r rune) error {
	s.LastRune = r

	switch s.State {
	case Start:
		if unicode.IsLetter(r) || r == '\'' {
			s.CurrentString += string(r)
			s.State = Operator
		} else {
			s.ResetState()
		}

	case Operator:
		if unicode.IsLetter(r) || r == '\'' {
			s.TryAcceptingNextRune(r)
		} else if r == '(' {
			s.State = OpenParen
			s.CurrentString += string(r)
			s.CurrentOperation.Operator = s.CurrentString
		} else if r == ')' {
			s.CurrentString += string(r)
			s.State = CloseParen
		} else {
			s.ResetState()
		}

	case OpenParen, Comma:
		if unicode.IsDigit(r) {
			s.CurrentValue.WriteRune(r)
			s.State = Value
		} else if r == ')' {
			s.CurrentString += string(r)
			s.State = CloseParen
		} else {
			s.ResetState()
		}

	case Value:
		if unicode.IsDigit(r) {
			s.CurrentValue.WriteRune(r)
		} else if r == ',' {
			s.TryParseValue()
		} else if r == ')' {
			s.TryParseValue()
			s.State = CloseParen
		} else {
			s.ResetState()
		}
	}

	if s.State == CloseParen && s.IsValidOperations(s.CurrentString) {
		if s.CurrentString == "do()" {
			s.Process = true
		} else {
			s.Process = false
		}
	}

	s.CheckState()
	return nil
}

func (s *StateMachine) IsValidOperator(operator, comparison string) bool {
	if len(operator) > len(comparison) {
		return false
	}
	return strings.HasPrefix(comparison, operator)
}

func (s *StateMachine) IsValidOperations(operation string) bool {
	methods := []string{"don't()", "do()", "mul"}
	for _, method := range methods {
		if s.IsValidOperator(operation, method) {
			return true
		}
	}
	return false
}

func (s *StateMachine) TryAcceptingNextRune(r rune) {
	m := s.CurrentString + string(r)
	if s.IsValidOperations(m) {
		s.CurrentString = m
	} else {
		s.CurrentString = string(r)
	}
}

func (s *StateMachine) TryParseValue() {
	// Convert and store the current value
	val, err := strconv.Atoi(s.CurrentValue.String())
	if err != nil {
		log.Printf("invalid number: %v", err)
		s.ResetState()
		return
	}
	s.CurrentOperation.Values = append(s.CurrentOperation.Values, val)
	s.CurrentValue.Reset()
	s.State = Comma
}

func (s *StateMachine) CheckOperation() {
	if s.Process && len(s.CurrentOperation.Values) == 2 && strings.HasPrefix(s.CurrentOperation.Operator, "mul") {
		s.ParsedOperations = append(s.ParsedOperations, s.CurrentOperation)
	}
}

func (s *StateMachine) CheckState() {
	switch s.State {
	case Start, Operator:
		if !s.IsValidOperations(s.CurrentString) {
			s.ResetState()
		}
	case CloseParen:
		s.CheckOperation()
		s.ResetState()
	}
}

func (s *StateMachine) ResetState() {
	s.CurrentString = string(s.LastRune)
	s.CurrentValue.Reset()
	s.State = Start
	s.CurrentOperation = NewOperation()
}
