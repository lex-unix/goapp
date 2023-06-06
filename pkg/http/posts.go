package http

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexunix/goapp/pkg/domain"
)

type PostHandler struct {
	PostService domain.PostService
}

func (h *PostHandler) Get(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing int")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Param should be integer",
		})
		return
	}
	post, err := h.PostService.Get(id)
	if err != nil {
		fmt.Println(fmt.Errorf("Error finding post: %v", err))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
		})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) Create(c *gin.Context) {
	var post domain.Post

	if err := c.BindJSON(&post); err != nil {
		fmt.Println(fmt.Errorf("Error binding JSON: %v", err))
		return
	}

	err := h.PostService.Create(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error creating post": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing int")
		return
	}
	_, err = h.PostService.Get(id)
	if err != nil {
		fmt.Println(fmt.Errorf("Post not found: %v", err))
		c.JSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}
	err = h.PostService.Delete(id)
	if err != nil {
		fmt.Println(fmt.Errorf("Error deleting post: %v", err))
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "post deleted"})
}
