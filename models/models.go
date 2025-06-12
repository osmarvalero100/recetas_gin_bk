package models

import "backend/database"

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Slug string `gorm:"type:varchar(100);not null;unique" json:"slug"`
}

type categories []Category

func Migrations() {
	db := database.Database()
	db.AutoMigrate(&Category{})
}
