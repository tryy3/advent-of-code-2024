package process

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Database struct {
	sql_db *sql.DB
}

func NewDatabase(sql_db *sql.DB) *Database {
	return &Database{sql_db: sql_db}
}

func (p *Database) InsertGeneration(generation int, value int) {
	f, err := os.OpenFile(fmt.Sprintf("cmd/generations/gen_%d.txt", generation), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("%d\n", value))
}

func (p *Database) ClearData() {
	files, err := os.ReadDir("cmd/generations")
	if err != nil {
		log.Println(err)
	}
	for _, file := range files {
		os.Remove(fmt.Sprintf("cmd/generations/%s", file.Name()))
	}
}

func (p *Database) OpenFile(generation int) (*os.File, error) {
	file, err := os.Open(fmt.Sprintf("cmd/generations/gen_%d.txt", generation))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (p *Database) GetGenerationCount(generation int) int {
	count := 0
	file, err := p.OpenFile(generation)
	if err != nil {
		log.Println(err)
		return 0
	}

	defer file.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Process each line
	for scanner.Scan() {
		count++
	}
	return count
}
