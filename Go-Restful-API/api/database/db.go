package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
