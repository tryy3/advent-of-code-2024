package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/tryy3/advent-of-code/day-11-part-1-files/process"
	_ "modernc.org/sqlite"
)

func main() {
	file := flag.String("file", "input.txt", "input file")
	input := flag.String("input", "", "input")
	generations := flag.Int("generations", 2, "number of generations to blink")
	flag.Parse()

	var processor *process.Processor
	var err error

	// Open a database connection
	sql_db, err := sql.Open("sqlite", "./example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer sql_db.Close()

	// Create a table
	createTableSQL := `CREATE TABLE IF NOT EXISTS day_11 (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			value INTEGER,
			generation INTEGER
		);`
	_, err = sql_db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	fmt.Println("Table created successfully")

	if *input != "" {
		processor, err = process.LoadProcessorFromReader(strings.NewReader(*input), sql_db)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		processor, err = process.LoadProcessorFromFile(*file, sql_db)
		if err != nil {
			log.Println(err)
			return
		}
	}

	for i := 0; i < *generations; i++ {
		currentTime := time.Now()
		fmt.Printf("Blinking generation %d\n", i)
		processor.Blink(i)
		elapsedTime := time.Since(currentTime)
		fmt.Printf("Blinking generation %d took %s\n", i, elapsedTime)
	}
	spew.Dump(processor.GetGenerationCount(*generations))
}
