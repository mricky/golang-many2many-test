package main

import (
	"fmt"
	"golang-many2many-test/handler"
	"golang-many2many-test/post"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbURL := "postgres://postgres:Rahasia123!@localhost:5432/db_many_to_many"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("Connection to db success")

	postRepository := post.NewRepository(db)

	userService := post.NewService(postRepository)

	postHandler := handler.NewPostHandler(userService)
	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/posts", postHandler.CreatePost)
	// router.GET("handler", handler)
	router.Run()
}
