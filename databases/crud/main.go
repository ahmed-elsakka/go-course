package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	ID    int
	Name  string
	Email string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/school"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	insertQuery := "INSERT INTO students (name, email) VALUES (?, ?)"
	result, err := db.Exec(insertQuery, "Alice", "alice@example.com")
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("Inserted student with ID %d\n", id)

	rows, err := db.Query("SELECT id, name, email FROM students")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var s Student
		err := rows.Scan(&s.ID, &s.Name, &s.Email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s - %s\n", s.ID, s.Name, s.Email)
	}

	updateQuery := "UPDATE students SET email=? WHERE id=?"
	_, err = db.Exec(updateQuery, "alice@newdomain.com", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated email successfully")

	deleteQuery := "DELETE FROM students WHERE id=?"
	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted student successfully")
}
