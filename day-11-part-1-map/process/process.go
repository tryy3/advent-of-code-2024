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
	values map[int]int
}

func NewProcessor() *Processor {
	return &Processor{}
}

func LoadProcessorFromFile(path string) (*Processor, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return LoadProcessorFromReader(file)
}

func LoadProcessorFromReader(reader io.Reader) (*Processor, error) {
	processor := &Processor{
		values: make(map[int]int),
	}
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")

		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			processor.values[num] = 1
		}
	}

	return processor, nil
}

func (p *Processor) GenerateValues(generation int) int {
	for i := 0; i < generation; i++ {
		p.Blink(i)
	}

	total := 0
	for _, value := range p.values {
		total += value
	}

	fmt.Printf("Total: %d\n", len(p.values))

	return total
}

func (p *Processor) Blink(generation int) {
	newMap := make(map[int]int)
	for key, value := range p.values {
		if key == 0 {
			v, ok := newMap[1]
			if !ok {
				newMap[1] = 0
			}
			newMap[1] = v + value
			continue
		}

		numStr := strconv.Itoa(key)
		if len(numStr)%2 == 0 {
			num1, _ := strconv.Atoi(numStr[:len(numStr)/2])
			num2, _ := strconv.Atoi(numStr[len(numStr)/2:])

			v1, ok := newMap[num1]
			if !ok {
				newMap[num1] = 0
			}
			newMap[num1] = v1 + value

			v2, ok := newMap[num2]
			if !ok {
				newMap[num2] = 0
			}
			newMap[num2] = v2 + value
			continue
		}

		newValue := key * 2024
		v, ok := newMap[newValue]
		if !ok {
			newMap[newValue] = 0
		}
		newMap[newValue] = v + value
	}

	p.values = newMap
}

func (p *Processor) GetGenerationCount(generation int) int {
	return len(p.values)
}
