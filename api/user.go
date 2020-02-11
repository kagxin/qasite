package api

import (
	"net/http"
	"qasite/middleware"
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
	var userVal UserValidation
	if err := c.ShouldBindJSON(&userVal); err == nil {
		var user model.User
		if model.DB.Where("username=?", userVal.Username).First(&user); user.Username == userVal.Username {
			c.JSON(http.StatusOK, gin.H{"message": "existed"})
			return
		}

		userD := model.User{Username: userVal.Username, Password: utils.Md5(userVal.Password)}
		if DB := model.DB.Create(&userD); DB.Error != nil {
			c.JSON(http.StatusOK, gin.H{"message": DB.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "ok"})

	} else {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
	}

}

// UserInfo asdf
func UserInfo(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	user, _ := c.Get(middleware.IdentityKey)
	c.JSON(200, gin.H{
		"id":       user.(*model.User).ID,
		"username": user.(*model.User).Username,
	})

}
