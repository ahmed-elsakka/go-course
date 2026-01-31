package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "admin:1234@/main"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println("Cant connect to db: ", err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Not connected to db: ", err)
	}

	fmt.Println("Connected to the database successfully")

}
