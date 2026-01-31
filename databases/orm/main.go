package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

func main() {
	dsn := "root:root@/main"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&User{})

	db.Create(&User{Name: "Bob", Email: "bob@example.com"})

	var user User
	db.Where("name = ?", "Bob").First(&user)
	fmt.Println(user)

	db.Model(&user).Update("Email", "bob@newdomain.com")

	//db.Delete(&student)
}
