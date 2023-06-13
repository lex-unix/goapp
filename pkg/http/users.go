package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/lexunix/goapp/pkg/domain"
)

type UserHandler struct {
	UserService domain.UserService
}

func (h *UserHandler) GetById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get user id",
		})
		return
	}
	user, err := h.UserService.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h UserHandler) Login(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to parse JSON",
		})
	}

	existingUser, err := h.UserService.User(user.Username)

	if err != nil {
		if err != pgx.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "User not found",
			})

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error finding user",
			})
		}
		return
	}

	if user.Password != existingUser.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Passwords don't match",
		})
		return

	}

	session := sessions.Default(c)
	session.Set("id", existingUser.ID)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "signed in",
	})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to parse JSON",
		})
		return
	}
	if err := h.UserService.Create(&user); err != nil {
		fmt.Println(fmt.Errorf("Error saving user: %v", err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save user",
		})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get user id",
		})
		return
	}
	if err := h.UserService.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to delete user",
		})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
