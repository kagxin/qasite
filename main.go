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
	model.DB.AutoMigrate(&model.User{})
}

func main() {

	defer model.DB.Close()

	r := gin.Default()
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:             middleware.KEY,
		Timeout:         time.Hour,
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
	// Refresh time can be longer than token timeout

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.GET("/user", api.UserInfo)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
