package http

import "github.com/gin-gonic/gin"

func PostRoutes(r *gin.Engine, h *PostHandler) {
	posts := r.Group("/posts")
	posts.Use(Authentication())
	posts.GET("", h.GetAll)
	posts.POST("", h.Create)
	posts.GET("/:id", h.GetOne)
	posts.DELETE("/:id", h.Delete)
	posts.GET("/user", h.UserPosts)
}

func UserRoutes(r *gin.Engine, h *UserHandler) {
	users := r.Group("/users")
	users.GET("/:id", h.GetById)
	users.POST("", h.CreateUser)
	users.POST("/login", h.Login)
}
