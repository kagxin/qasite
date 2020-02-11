package api

import (
	"qasite/model"
	"qasite/utils"

	"github.com/gin-gonic/gin"
)

// UserValidation for create user
type UserValidation struct {
	Username string `form:"username" binding:"required" validate:"max=10,min=1"`
	Password string `form:"password" binding:"required" validate:"max=10,min=1"`
}

// RegisterUser  asdf
func RegisterUser(c *gin.Context) {
	var user UserValidation
	if err := c.ShouldBindJSON(&user); err == nil {
		// EUser = model.User{}
		// model.DB.Where().First(&)
		userD := model.User{Username: user.Username, Password: utils.Md5(user.Password)}
		if DB := model.DB.Create(&userD); DB.Error != nil {
			c.JSON(200, gin.H{"message": DB.Error.Error()})
		} else {
			c.JSON(200, gin.H{"message": "ok"})
		}
	} else {
		c.JSON(200, gin.H{"message": err.Error()})
	}

}
