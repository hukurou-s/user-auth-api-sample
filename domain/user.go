package domain

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;unique;not null" json:"name"`
	Email    string `gorm:"size:255;unique;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`
}

type LoginParams struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
