package process

import (
	"database/sql"
	"log"
)

type Database struct {
	sql_db *sql.DB
}

func NewDatabase(sql_db *sql.DB) *Database {
	return &Database{sql_db: sql_db}
}

func (p *Database) InsertGeneration(generation int, value int) {

	insertSQL := `INSERT INTO day_11 (value, generation) VALUES (?, ?)`
	_, err := p.sql_db.Exec(insertSQL, value, generation)
	if err != nil {
		log.Fatalf("Error inserting data: %v", err)
	}
}

func (p *Database) ClearData() {
	deleteSQL := `DELETE FROM day_11`
	_, err := p.sql_db.Exec(deleteSQL)
	if err != nil {
		log.Fatalf("Error clearing table: %v", err)
	}
}

func (p *Database) GetGeneration(generation int, ids *[]int) {
	querySQL := `SELECT id FROM day_11 WHERE generation = ?`
	rows, err := p.sql_db.Query(querySQL, generation)
	if err != nil {
		log.Fatalf("Error querying data: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		*ids = append(*ids, id)
	}
}

type GenerationValue struct {
	ID         int `json:"id"`
	Generation int `json:"generation"`
	Value      int `json:"value"`
}

func (p *Database) GetGenerationValues(generation int, id int) GenerationValue {
	generationValues := []GenerationValue{}
	querySQL := `SELECT id, generation, value FROM day_11 WHERE generation = ? AND id = ?`
	rows, err := p.sql_db.Query(querySQL, generation, id)
	if err != nil {
		log.Fatalf("Error querying data: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var generation int
		var value int
		err = rows.Scan(&id, &generation, &value)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		generationValues = append(generationValues, GenerationValue{ID: id, Generation: generation, Value: value})
	}
	return generationValues[0]
}

func (p *Database) GetGenerationCount(generation int) int {
	querySQL := `SELECT COUNT(*) FROM day_11 WHERE generation = ?`
	rows, err := p.sql_db.Query(querySQL, generation)
	if err != nil {
		log.Fatalf("Error querying data: %v", err)
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			log.Println(err)
		}
	}
	return int(count)
}
