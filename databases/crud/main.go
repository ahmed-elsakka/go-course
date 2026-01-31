package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	ID    int
	Name  string
	Email string
}

func main() {
	dsn := "admin:1234@/main"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error connecting to db: ", err)
	}
	defer db.Close()

	insertQuery := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := db.Exec(insertQuery, "Alice", "alice@example.com")
	if err != nil {
		fmt.Println("Error querying db: ", err)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("Inserted user with ID %d\n", id)

	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		fmt.Println("Error querying db: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var s Student
		err := rows.Scan(&s.ID, &s.Name, &s.Email)
		if err != nil {
			fmt.Println("Error querying db: ", err)
		}
		fmt.Printf("%d: %s - %s\n", s.ID, s.Name, s.Email)
	}

	updateQuery := "UPDATE users SET email=? WHERE id=?"
	_, err = db.Exec(updateQuery, "alice@newdomain.com", id)
	if err != nil {
		fmt.Println("Error querying db: ", err)
	}
	fmt.Println("Updated email successfully")

	deleteQuery := "DELETE FROM users WHERE id=?"
	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		fmt.Println("Error querying db: ", err)
	}
	fmt.Println("Deleted student successfully")
}
