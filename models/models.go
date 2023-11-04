package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName   string `gorm:"not null;unique; size:25"`
	Email      string `gorm:"not null;unique;" validate:"email"`
	Password   string `gorm:"not null;size:10"`
	Phone      int    `gorm:"not null;unique"`
	FirstName  string `gorm:"not null;size:100"`
	LastName   string `gorm:"not null;size:100"`
	Gender     bool
	BirthPlace string    `gorm:"not null"`
	BirthDate  time.Time `gorm:"not null"`
	Address    string
	Posts      []Post `gorm:"foreignKey:UserID"`
}

type Categories struct {
	gorm.Model
	Name string `gorm:"not null;size:100"`
}

type Post struct {
	gorm.Model
	Title       string `gorm:"not null;size:255"`
	Img         string
	Description string
	UserID      uint       `gorm:"not null;"`
	CategoryID  uint       `gorm:"not null;"`
	Category    Categories `gorm:"foreignKey:CategoryID"`
	User        User       `gorm:"foreignKey:UserID"`
}
