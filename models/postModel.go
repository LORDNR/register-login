package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"not null; default:null"`
	Body  string `gorm:"not null; default:null"`
}
