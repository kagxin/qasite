package middleware

import (
	"qasite/api"
	"qasite/errno"
	"qasite/model"
	"qasite/utils/response"

	"github.com/gin-gonic/gin"
)

func BasicTokenAuth(service *api.Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.JSON(c, errno.TokenNotFound, gin.H{})
			return
		}
		token := model.Token{}
		if err := service.Mysql.DB.Where("token=?", tokenStr).First(&token).Error; err != nil {
			response.JSON(c, errno.TokenNotFound, gin.H{})
			return
		}
		user := model.User{}
		if err := service.Mysql.DB.Where("id=?", token.UserID).First(&user).Error; err != nil {
			response.JSON(c, errno.TokenNotFound, gin.H{})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
