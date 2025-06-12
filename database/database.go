package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database = func() (db *gorm.DB) {
	errorVars := godotenv.Load()
	if errorVars != nil {
		panic("Error loading .env file")
		return
	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connected to the database successfully")
		return db
	}
}
