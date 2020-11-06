package db

import "fmt"

// Database ...
type Database struct {
	query string
}

var instance *Database

// GetInstance ...
func GetInstance() *Database {
	if instance == nil {
		instance = &Database{}
	}

	return instance
}

// Query ...
func (d *Database) Query(q string) *Database {
	fmt.Println("Query parsed")
	d.query = q

	return d
}

// Exec ...
func (d *Database) Exec() error {
	fmt.Printf("Executed query is: %v\n", d.query)
	return nil
}
