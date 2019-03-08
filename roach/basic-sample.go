package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://roach@localhost:26257/bank?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	if _, err := db.Exec("INSERT INTO accounts (id, balance) VALUES (4,200),(5,300)"); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, balance FROM accounts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("Initial balances: ")
	for rows.Next() {
		var id, balance int
		if err := rows.Scan(&id, &balance); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d -- %d\n", id, balance)
	}
}
