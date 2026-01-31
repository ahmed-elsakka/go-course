package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	dsn := "admin:1234@/main"
	db, err := sql.Open("mysql", dsn)
	handleError(err)
	defer db.Close()

	insertQuery := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := db.Exec(insertQuery, "Alice", "alice@example.com")
	handleError(err)

	id, _ := result.LastInsertId()
	fmt.Println("Inserted a new user with id: ", id)

	selectQuery := "SELECT id, name, email FROM users"
	rows, err := db.Query(selectQuery)
	handleError(err)

	defer rows.Close()

	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name, &user.Email)
		fmt.Printf("%d: %s - %s\n", user.ID, user.Name, user.Email)
	}
	updateQuery := "UPDATE users SET email=? WHERE id=?"
	result, err = db.Exec(updateQuery, "alice@newemail.com", "1")
	handleError(err)

	deleteQuery := "DELETE FROM users WHERE id=?"
	result, err = db.Exec(deleteQuery, id)
	handleError(err)

}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
