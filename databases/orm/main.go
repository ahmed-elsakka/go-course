package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Student{})

	db.Create(&Student{Name: "Bob", Email: "bob@example.com"})

	var student Student
	db.Where("name = ?", "Bob").First(&student)
	fmt.Println(student)

	db.Model(&student).Update("Email", "bob@newdomain.com")

	//db.Delete(&student)
}
