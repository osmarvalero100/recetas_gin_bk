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

type Recipes struct {
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

func Migrations() {
	db := database.Database()
	db.AutoMigrate(&Category{}, &Recipes{})
}
