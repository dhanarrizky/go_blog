package database

import (
	"log"

	"github.com/dhanarrizky/go-blog/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/go_blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed toconnect database")
		panic(err.Error())
	}

	return db
}

func GetDB() {
	db := ConDB()
	User := models.User{}
	Post := models.Post{}
	Categories := models.Categories{}
	db.AutoMigrate(&User, &Post, &Categories)
	log.Println("connected to database successfully...")
}
