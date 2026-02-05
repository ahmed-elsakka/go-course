package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := "admin:1234@/main"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
