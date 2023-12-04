package handler

import (
	"golang-many2many-test/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

type postHandler struct {
	postService post.Service
}

func NewPostHandler(postService post.Service) *postHandler {
	return &postHandler{postService}
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
