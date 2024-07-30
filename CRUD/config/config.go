package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=crud sslmode=disable password=sourav")
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	DB.LogMode(true)
}
