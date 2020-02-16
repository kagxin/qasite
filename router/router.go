package router

import (
	"qasite/api"
	"qasite/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/login", middleware.AuthMiddleware.LoginHandler)
	r.POST("/register", api.RegisterUser)

	auth := r.Group("/")

	auth.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", middleware.AuthMiddleware.RefreshHandler)
		auth.GET("/user", api.UserInfo)
		auth.POST("/question", api.CreateQuestion)
		auth.GET("/question/:QID", api.ShowQuestion)
		auth.GET("/question", api.ShowQuestion)
		auth.POST("/question/:QID/commit", api.CreateComment)
		auth.GET("/question/:QID/commit", api.GetComment)

	}
	return r
}
