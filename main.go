package main

import (
	"log"
	"qasite/api"
	"qasite/middleware"
	model "qasite/model"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func init() {
	model.DB.SingularTable(true)
	model.DB.AutoMigrate(&model.User{}, &model.Question{}, &model.Comment{})
}

func main() {

	defer model.DB.Close()
	model.DB.LogMode(true)
	r := gin.Default()
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:             middleware.KEY,
		Timeout:         time.Hour * 24 * 10,
		MaxRefresh:      time.Hour,
		IdentityKey:     middleware.IdentityKey,
		PayloadFunc:     middleware.PayloadFunc,
		IdentityHandler: middleware.IdentityHandlerFunc,
		Authenticator:   middleware.AuthenticatorFunc,
		Authorizator:    middleware.AuthorizatorFunc,
		Unauthorized:    middleware.UnauthorizedFunc,
		TokenLookup:     "header: Authorization",

		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", api.RegisterUser)

	auth := r.Group("/")

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.GET("/user", api.UserInfo)
		auth.POST("/question", api.CreateQuestion)
		auth.GET("/question/:QID", api.ShowQuestion)
		auth.GET("/question", api.ShowQuestion)
		auth.POST("/question/:QID/commit", api.CreateComment)
		auth.GET("/question/:QID/commit", api.GetComment)

	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
