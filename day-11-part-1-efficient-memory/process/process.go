package process

import (
	"bufio"
	"database/sql"
	"io"
	"os"
	"strconv"
	"strings"
)

type Processor struct {
	values []int
	db     *Database
}

func NewProcessor(db *Database) *Processor {
	return &Processor{db: db}
}

func LoadProcessorFromFile(path string, sql_db *sql.DB) (*Processor, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return LoadProcessorFromReader(file, sql_db)
}

func LoadProcessorFromReader(reader io.Reader, sql_db *sql.DB) (*Processor, error) {
	processor := &Processor{}
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
			processor.values = append(processor.values, num)
		}
	}

	return processor, nil
}

func (p *Processor) GenerateValues(generation int) int {
	// var total = make(chan int)
	var total int
	// var wg sync.WaitGroup

	var generate func(value int, gen int)

	generate = func(value int, gen int) {
		// defer wg.Done()

		if gen == generation {
			if value == 0 {
				total += 1
				return
			}
			numStr := strconv.Itoa(value)
			if len(numStr)%2 == 0 {
				total += 2
				return
			}

			total += 1
		} else {
			if value == 0 {
				// wg.Add(1)
				generate(1, gen+1)
				return
			}
			numStr := strconv.Itoa(value)
			if len(numStr)%2 == 0 {
				num1, _ := strconv.Atoi(numStr[:len(numStr)/2])
				num2, _ := strconv.Atoi(numStr[len(numStr)/2:])
				// wg.Add(2)
				generate(num1, gen+1)
				generate(num2, gen+1)
				return
			}

			// wg.Add(1)
			generate(value*2024, gen+1)
		}
	}

	for _, value := range p.values {
		// wg.Add(1)
		generate(value, 1)
	}

	// go func() {
	// 	wg.Wait()
	// 	close(total)
	// }()

	// totalCount := 0
	// for val := range total {
	// 	totalCount += val
	// }

	return total
}

func (p *Processor) Blink(generation int) {
	newValues := []int{}

	for _, value := range p.values {

		// Rule #1 if the stone is equal to 0, replace with 1
		if value == 0 {
			newValues = append(newValues, 1)
			continue
		}

		// Rule #2 if the number of digits is even, split into two stones
		// remove leading 0's
		numStr := strconv.Itoa(value)
		if len(numStr)%2 == 0 {
			num1, _ := strconv.Atoi(numStr[:len(numStr)/2])
			num2, _ := strconv.Atoi(numStr[len(numStr)/2:])
			newValues = append(newValues, num1)
			newValues = append(newValues, num2)
			continue
		}

		// Last rule: multiply by 2024
		newValues = append(newValues, value*2024)
	}

	p.values = newValues
}

func (p *Processor) GetGenerationCount(generation int) int {
	return len(p.values)
}

func (p *Processor) GetValues() []int {
	return p.values
}
