package models

import (
	"time"

	_ "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName     string    `gorm:"not null;unique; size:25" json:"user_name" validate:"required" `
	Email        string    `gorm:"not null;unique;" json:"email" validate:"email,required" `
	Phone        int       `gorm:"not null;unique;" json:"phone" validate:"required"`
	Password     string    `gorm:"not null;" json:"password" validate:"required"`
	FirstName    string    `gorm:"not null;type:varchar(100);" json:"rirst_name" validate:"required"`
	LastName     string    `gorm:"not null;type:varchar(100);" json:"last_name"`
	Gender       bool      `json:"gender"`
	Token        *string   `gorm:"not null;" json:"token"`
	RefreshToken *string   `gorm:"not null;" json:"refresh_token"`
	BirthPlace   string    `gorm:"not null;" json:"birth_place" validate:"required"`
	Birth        string    `gorm:"not null;" json:"birthdate" validate:"required"`
	BirthDate    time.Time `gorm:"not null;" json:"birth_date"`
	Address      string    `json:"address"`
	Posts        []Post    `gorm:"foreignKey:UserID;" json:"posts"`
	Role         string    `gorm:"not null;size:8;" json:"role" validate:"eq=ADMIN|eq=USER"`
}

type Categories struct {
	gorm.Model
	Name string `gorm:"not null;type:varchar(100)" json:"name" validate:"required"`
}

// type ImgFile struct {
// 	ImgFile *multipart.FileHeader
// }

type Post struct {
	gorm.Model
	Title string `gorm:"not null;type:varchar(255);" json:"title" validate:"required"`
	// Img   multipart.FileHeader `json:"img"`
	Img         string     `json:"img"`
	Description string     `json:"description" validate:"required"`
	UserID      uint       `gorm:"not null;" json:"user_id"`
	CategoryID  uint       `gorm:"not null;" json:"category_id"`
	Category    Categories `gorm:"foreignKey:CategoryID;" json:"category"`
	User        User       `gorm:"foreignKey:UserID;" json:"user"`
	// Images      ImgFile
}

// almost all of post is non
