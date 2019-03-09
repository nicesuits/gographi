package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// TodoDB struct
type TodoDB struct {
	ID      int
	balance float64
}

func main() {
	db, err := sql.Open("postgres", "postgresql://roach@localhost:26257/bank?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	if _, err := db.Exec("INSERT INTO accounts (id, balance) VALUES (10,200.5),(11,300.0)"); err != nil {
		log.Fatal(err)
	}

	var todos []TodoDB
	rows, err := db.Query("SELECT id, balance FROM accounts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("Initial balances: ")
	for rows.Next() {
		var todo TodoDB
		if err := rows.Scan(&todo.ID, &todo.balance); err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
		fmt.Printf("%d -- %.2f\n", todo.ID, todo.balance)
	}
}
