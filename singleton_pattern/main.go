package main

import (
	"fmt"

	"github.com/MrWebUzb/facade/singleton_pattern/db"
)

func main() {
	conn := db.GetInstance()

	query := "SELECT * FROM users WHERE username='john' AND password='doe';"

	if err := conn.Query(query).Exec(); err != nil {
		fmt.Printf("Error when execute query: %v", err)
	}
}
