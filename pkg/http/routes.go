package http

import "github.com/gin-gonic/gin"

func PostRoutes(r *gin.Engine, h *PostHandler) {
	posts := r.Group("/post")
	posts.POST("", h.Create)
	posts.GET("/:id", h.Get)
	posts.DELETE("/:id", h.Delete)
}

func UserRoutes(r *gin.Engine, h *UserHandler) {
	users := r.Group("/user")
	users.GET("/:id", h.GetById)
	users.POST("", h.CreateUser)
}
