package db

import (
	"golang-many2many-test/post"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:Rahasia123!@localhost:5432/pg_many_to_many"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&post.Post{})

	return db
}
