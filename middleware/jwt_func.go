package middleware

import (
	"fmt"
	"qasite/model"
	"qasite/utils"
	"reflect"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type UserValidation struct {
	Username string `form:"username" binding:"required" validate:"max=10,min=1"`
	Password string `form:"password" binding:"required" validate:"max=10,min=1"`
}

// IdentityKey asf
var IdentityKey string = "Username"

// KEY sercet key
var KEY []byte = []byte("asdfjakldsfjlskjflkd")

// UnauthorizedFunc asdfjsdasd
func UnauthorizedFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"code": code, "message": message})
}

// AuthorizatorFunc asdfasdf
func AuthorizatorFunc(data interface{}, c *gin.Context) bool {
	var a int = 10
	fmt.Println(a)
	return true
}

// AuthenticatorFunc asdf
func AuthenticatorFunc(c *gin.Context) (interface{}, error) {
	var loginVals UserValidation
	var user model.User
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	model.DB.Where("username=?", loginVals.Username).First(&user)

	if utils.Md5(loginVals.Password) == user.Password {
		return &user, nil
	}
	return nil, jwt.ErrFailedAuthentication
}

// PayloadFunc asdf
func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*model.User); ok {
		fmt.Println(reflect.TypeOf(v.ID))
		return jwt.MapClaims{
			IdentityKey: v.Username,
		}
	}
	return jwt.MapClaims{}
}

// IdentityHandlerFunc asdf
func IdentityHandlerFunc(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	var user model.User
	username := claims[IdentityKey]
	model.DB.Where("username=?", username).First(&user)
	return &user
}
