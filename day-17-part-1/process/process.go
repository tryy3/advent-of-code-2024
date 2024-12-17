package process

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Processor struct {
	computer *Computer
}

func NewProcessor() *Processor {
	return &Processor{
		computer: NewComputer(),
	}
}

func (p *Processor) LoadFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return p.LoadFromReader(file)
}

func (p *Processor) LoadFromReader(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Register A:") {
			v := strings.TrimPrefix(line, "Register A:")
			v = strings.TrimSpace(v)

			num, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("failed to parse register A: %w", err)
			}

			p.computer.registerA.SetValue(num)
			continue
		}
		if strings.HasPrefix(line, "Register B:") {
			v := strings.TrimPrefix(line, "Register B:")
			v = strings.TrimSpace(v)

			num, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("failed to parse register B: %w", err)
			}

			p.computer.registerB.SetValue(num)
			continue
		}
		if strings.HasPrefix(line, "Register C:") {
			v := strings.TrimPrefix(line, "Register C:")
			v = strings.TrimSpace(v)

			num, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("failed to parse register C: %w", err)
			}

			p.computer.registerC.SetValue(num)
			continue
		}
		if strings.HasPrefix(line, "Program:") {
			v := strings.TrimPrefix(line, "Program:")
			v = strings.TrimSpace(v)

			instructions := strings.Split(v, ",")
			for _, instruction := range instructions {
				num, err := strconv.Atoi(instruction)
				if err != nil {
					return fmt.Errorf("failed to parse program instruction: %w", err)
				}
				p.computer.program.AddInstruction(num)
			}
			continue
		}
	}

	return nil
}

func (p *Processor) ReadInstruction() bool {

	return p.computer.ReadNextInstruction()
}

func (p *Processor) GetRegisterA() int {
	return p.computer.GetRegisterAValue()
}

func (p *Processor) GetRegisterB() int {
	return p.computer.GetRegisterBValue()
}

func (p *Processor) GetRegisterC() int {
	return p.computer.GetRegisterCValue()
}

func (p *Processor) GetInstructionPointer() int {
	return p.computer.GetInstructionPointer()
}

func (p *Processor) PrintOutput() {
	output := []string{}
	for _, v := range p.computer.GetOutput() {
		output = append(output, strconv.Itoa(v))
	}

	fmt.Println(strings.Join(output, ","))
}
