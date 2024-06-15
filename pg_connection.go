package main

import (
	"log"

	"github.com/Harsh-apk/notesPostgres/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=harsh password=root dbname=users port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)

	}

	//run this whenever you edit struct or
	//create new package which you can run standalone to make new tables
	//as per requirements
	err = db.AutoMigrate(&types.User{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
