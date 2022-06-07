package services

import (
	"final-project/app/model"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabaseTest() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("testsqlite.db"), &gorm.Config{})
	if nil == err {
		fmt.Println("Database error")
	}
	db.AutoMigrate(&model.Car{})
	DB = db
	return db
}
