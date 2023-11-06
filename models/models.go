package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName     string    `gorm:"not null;unique; size:25" json:"user_name"`
	Email        string    `gorm:"not null;unique;" validate:"email" json:"email"`
	Phone        int       `gorm:"not null;unique" json:"phone"`
	Password     string    `gorm:"not null;size:10" json:"password"`
	FirstName    string    `gorm:"not null;size:100" json:"rirst_name"`
	LastName     string    `gorm:"not null;size:100" json:"last_name"`
	Gender       bool      `json:"gender"`
	Token        *string   `gorm:"not null" json:"token"`
	RefreshToken *string   `gorm:"not null" json:"refresh_token"`
	BirthPlace   string    `gorm:"not null" json:"birth_place"`
	BirthDate    time.Time `gorm:"not null" json:"birth_date"`
	Address      string    `json:"address"`
	Posts        []Post    `gorm:"foreignKey:UserID" json:"posts"`
}

type Categories struct {
	gorm.Model
	Name string `gorm:"not null;size:100" json:"name"`
}

type Post struct {
	gorm.Model
	Title       string     `gorm:"not null;size:255" json:"title"`
	Img         string     `json:"img"`
	Description string     `json:"description"`
	UserID      uint       `gorm:"not null;" json:"user_id"`
	CategoryID  uint       `gorm:"not null;" json:"category_id"`
	Category    Categories `gorm:"foreignKey:CategoryID" json:"category"`
	User        User       `gorm:"foreignKey:UserID" json:"user"`
}
