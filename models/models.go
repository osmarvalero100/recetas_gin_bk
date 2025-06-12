package models

import (
	"backend/database"
	"time"
)

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Slug string `gorm:"type:varchar(100);not null;unique" json:"slug"`
}

type categories []Category

type Recipe struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CategoryID  uint      `gorm:"not null" json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	slug        string    `gorm:"type:varchar(100);not null;unique" json:"slug"`
	Time        string    `gorm:"type:varchar(100);not null" json:"time"`
	Photo       string    `gorm:"type:varchar(255);not null" json:"photo"`
	Description string    `gorm:"type:text;not null" json:"description"`
	DateAdd     time.Time `json:"date_add"`
}

type Contact struct {
	ID      uint      `gorm:"primaryKey" json:"id"`
	Name    string    `gorm:"type:varchar(100);not null" json:"name"`
	Email   string    `gorm:"type:varchar(100);not null" json:"email"`
	Phone   string    `gorm:"type:varchar(15);not null" json:"phone"`
	Message string    `gorm:"type:text;not null" json:"message"`
	DateAdd time.Time `json:"date_add"`
}

func Migrations() {
	db := database.Database()
	db.AutoMigrate(&Category{}, &Recipe{}, &Contact{})
}
