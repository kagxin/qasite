package middleware

import (
	"github.com/gin-gonic/gin"
	"qasite/api"
	"qasite/errno"
	"qasite/model"
	"qasite/utils/response"
)

func BasicTokenAuth(service *api.Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.JSON(c, errno.TokenNotFound, gin.H{})
			return
		}
		token := model.Token{}
		if err := service.DB.Where("token=?", tokenStr).First(&token).Error; err != nil {
			response.JSON(c, errno.TokenNotFound, gin.H{})
			return
		}
		user := model.User{}
		if err := service.DB.Where("id=?", token.UserID).First(&user).Error; err != nil {
			response.JSON(c, errno.TokenNotFound, gin.H{})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
