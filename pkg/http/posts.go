package http

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lexunix/goapp/pkg/domain"
)

type PostHandler struct {
	PostService domain.PostService
}

func (h *PostHandler) GetAll(c *gin.Context) {
	posts, err := h.PostService.All()
	fmt.Println(c.Cookie("cook-my-sess"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) GetOne(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing int")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Param should be integer",
		})
		return
	}
	session := sessions.Default(c)
	session.Get("id")
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

	if err := h.PostService.Create(&post); err != nil {
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
	if _, err = h.PostService.Get(id); err != nil {
		fmt.Println(fmt.Errorf("Post not found: %v", err))
		c.JSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}
	if err = h.PostService.Delete(id); err != nil {
		fmt.Println(fmt.Errorf("Error deleting post: %v", err))
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "post deleted"})
}

func (h *PostHandler) UserPosts(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("id")

	userID, ok := v.(int64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "session is empty"})
		return
	}

	posts, err := h.PostService.UserPosts(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot parse session key to int"})
		return
	}
	c.JSON(http.StatusOK, posts)
}
