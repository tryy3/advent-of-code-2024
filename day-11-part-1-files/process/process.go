package process

import (
	"bufio"
	"database/sql"
	"io"
	"log"
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
	database := NewDatabase(sql_db)
	database.ClearData()
	processor := &Processor{db: database}
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	generation := 0

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")

		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			processor.values = append(processor.values, num)

			processor.db.InsertGeneration(generation, num)
		}
	}

	return processor, nil
}

func (p *Processor) Blink(generation int) {
	file, err := p.db.OpenFile(generation)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Process each line
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.Atoi(line)
		if err != nil {
			log.Println(err)
			return
		}

		// Rule #1 if the stone is equal to 0, replace with 1
		if value == 0 {
			p.db.InsertGeneration(generation+1, 1)
			continue
		}

		// Rule #2 if the number of digits is even, split into two stones
		// remove leading 0's
		numStr := strconv.Itoa(value)
		if len(numStr)%2 == 0 {
			num1, _ := strconv.Atoi(numStr[:len(numStr)/2])
			num2, _ := strconv.Atoi(numStr[len(numStr)/2:])
			p.db.InsertGeneration(generation+1, num1)
			p.db.InsertGeneration(generation+1, num2)
			continue
		}

		// Last rule: multiply by 2024
		p.db.InsertGeneration(generation+1, value*2024)
	}
}

func (p *Processor) GetGenerationCount(generation int) int {
	return p.db.GetGenerationCount(generation)
}

func (p *Processor) GetValues() []int {
	return p.values
}
