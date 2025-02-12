package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"unique;not null" json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type Product struct {
    gorm.Model
    Name        string  `json:"name" gorm:"not null"`
    Description string  `json:"description"`
    Price       float64 `json:"price" gorm:"not null"`
    Stock       int     `json:"stock" gorm:"not null"`
    UserID      uint    `json:"user_id" gorm:"not null"`
    User        User    `json:"user" gorm:"foreignKey:UserID"`
}
