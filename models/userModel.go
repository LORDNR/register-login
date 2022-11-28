package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email      string `gorm:"uniqueIndex; not null ; default:null"`
	Password   string `gorm:"not null; default:null"`
	First_Name string `gorm:"not null; default:null"`
	Last_Name  string `gorm:"not null; default:null"`
}
