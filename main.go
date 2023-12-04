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
	dbURL := "postgres://postgres:Rahasia123!@localhost:5432/pg_many_to_many"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	db.AutoMigrate(&post.Post{})

	fmt.Println("Connection to db success")

	postRepository := post.NewRepository(db)
	// just for test repository
	posts, _ := postRepository.FindAll()

	fmt.Println("debug")
	fmt.Println(len(posts))

	for _, post := range posts {
		fmt.Println(post.TITLE)
	}
	userService := post.NewService(postRepository)

	postHandler := handler.NewPostHandler(userService)
	router := gin.Default()
	api := router.Group("api")

	api.POST("/posts", postHandler.CreatePost)
	api.GET("/posts", postHandler.FindAll)
	api.GET("/posts/:id", postHandler.FindById)
	// router.GET("handler", handler)
	router.Run()
}
