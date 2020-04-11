package router

import (
	"qasite/api"
	"qasite/middleware"

	"github.com/gin-gonic/gin"
)

// Router Init
func Router(srv *api.Service) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", srv.Ping)
	r.POST("/register", srv.Register)
	r.POST("/login", srv.Login)
	authorized := r.Group("", middleware.BasicTokenAuth(srv))
	{
		authorized.GET("/user", srv.UserInfo)
		authorized.GET("/question", srv.Question)
		authorized.POST("/question", srv.CreateQuestion)
		authorized.GET("/question/:QID", srv.QuestionDetail)
		authorized.GET("/question/:QID/comment", srv.Comment)
		authorized.POST("/question/:QID/comment", srv.CreateComment)
	}

	return r
}
