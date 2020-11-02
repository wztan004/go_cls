package main

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"fmt"
)

func main() {
	connStr := "user=postgres dbname=socratica password=password sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT choice_text FROM polls_choice")

	if err != nil {
		fmt.Println("E",err)
	}

	defer rows.Close()


	for rows.Next() {
		var choice_text string
		err := rows.Scan(&choice_text)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("%s\n", choice_text)
	}
}