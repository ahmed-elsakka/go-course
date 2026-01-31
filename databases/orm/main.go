package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Email string
}

func main() {
	dsn := "admin:1234@/main"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to db: ", err)
	}

	newUser := User{Name: "Alice", Email: "alice@email.com"}
	db.Create(&newUser)

	var user User
	db.Where("name = ?", "Alice").First(&user)
	fmt.Println(user)

	db.Model(&user).Update("Email", "alice@newemail.com")

	db.Delete(&user)

}
