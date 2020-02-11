package main

import (
	"fmt"
	"qasite/api"
	model "qasite/model"

	"github.com/gin-gonic/gin"
)

func init() {
	model.DB.AutoMigrate(&model.User{})
	fmt.Println("main init.")
}
func main() {

	defer model.DB.Close()

	r := gin.Default()
	r.POST("/register", api.RegisterUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
