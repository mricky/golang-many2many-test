package handler

import (
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

	post, err := h.postService.CreatePost(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	c.JSON(http.StatusOK, post)
}
