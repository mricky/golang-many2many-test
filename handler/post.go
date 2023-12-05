package handler

import (
	"fmt"
	"golang-many2many-test/post"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type postHandler struct {
	postService post.Service
}

func NewPostHandler(postService post.Service) *postHandler {
	return &postHandler{postService}
}

func (h *postHandler) Delete(c *gin.Context) {
	postId, _ := strconv.Atoi(c.Param("id"))

	result := h.postService.Delete(postId)

	if result == 0 {
		c.JSON(http.StatusOK, "Gagal Delete")
	} else {
		c.JSON(http.StatusOK, "Berhasi Delete")
	}

	c.JSON(http.StatusOK, "Berhasi Delete")
}
func (h *postHandler) Update(c *gin.Context) {
	postId, _ := strconv.Atoi(c.Param("id"))
	var input post.PostInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	post, err := h.postService.Update(postId, input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	c.JSON(http.StatusOK, post)
}
func (h *postHandler) FindById(c *gin.Context) {
	postId, _ := strconv.Atoi(c.Param("id"))
	posts, err := h.postService.FindById(postId)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	c.JSON(http.StatusOK, posts)

}
func (h *postHandler) FindAll(c *gin.Context) {
	posts, err := h.postService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	c.JSON(http.StatusOK, posts)
}
func (h *postHandler) CreatePost(c *gin.Context) {
	var input post.PostInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	fmt.Println("Request Input")
	fmt.Println(input)
	post, err := h.postService.CreatePost(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	c.JSON(http.StatusOK, post)
}
