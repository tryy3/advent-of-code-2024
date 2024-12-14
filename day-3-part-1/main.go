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
	CurrentOperation *Operation
	ParsedOperations []*Operation
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		State:            Start,
		CurrentString:    "",
		CurrentOperation: NewOperation(),
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
		if operation.Operator == "mul" {
			total += operation.Values[0] * operation.Values[1]
		}
	}
	return total
}

func (s *StateMachine) NextRune(r rune) error {
	s.LastRune = r

	switch s.State {
	case Start:
		if unicode.IsLetter(r) {
			s.CurrentString += string(r)
			s.State = Operator
		} else {
			s.ResetState()
		}

	case Operator:
		if unicode.IsLetter(r) {
			s.TryAcceptingNextRune(r)
		} else if r == '(' {
			s.State = OpenParen
			s.CurrentOperation.Operator = s.CurrentString
			s.CurrentString = ""
		} else {
			s.ResetState()
		}

	case OpenParen, Comma:
		if unicode.IsDigit(r) {
			s.CurrentString += string(r)
			s.State = Value
		} else {
			s.ResetState()
		}

	case Value:
		if unicode.IsDigit(r) {
			s.CurrentString += string(r)
		} else if r == ',' {
			s.TryParseValue()
		} else if r == ')' {
			s.TryParseValue()
			s.State = CloseParen
		} else {
			s.ResetState()
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

func (s *StateMachine) TryAcceptingNextRune(r rune) {
	m := s.CurrentString + string(r)
	if !s.IsValidOperator(m, "mul") {
		s.CurrentString = string(r)
	} else {
		s.CurrentString = m
	}
}

func (s *StateMachine) TryParseValue() {
	// Convert and store the current value
	val, err := strconv.Atoi(s.CurrentString)
	if err != nil {
		log.Printf("invalid number: %v", err)
		s.ResetState()
		return
	}
	s.CurrentOperation.Values = append(s.CurrentOperation.Values, val)
	s.CurrentString = ""
	s.State = Comma
}

func (s *StateMachine) CheckOperation() {
	if len(s.CurrentOperation.Values) == 2 && s.CurrentOperation.Operator == "mul" {
		s.ParsedOperations = append(s.ParsedOperations, s.CurrentOperation)
	}
}

func (s *StateMachine) CheckState() {
	switch s.State {
	case Start, Operator:
		if !s.IsValidOperator(s.CurrentString, "mul") {
			s.ResetState()
		}
	case CloseParen:
		s.CheckOperation()
		s.ResetState()
	}
}

func (s *StateMachine) ResetState() {
	s.CurrentString = string(s.LastRune)
	s.State = Start
	s.CurrentOperation = NewOperation()
}
