package api

import (
	"qasite/errno"
	"qasite/model"
	"qasite/utils"
	"qasite/utils/response"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Service) Login(c *gin.Context) {
	loginReq := model.LoginReq{}
	if err := c.ShouldBind(&loginReq); err != nil {
		response.JSON(c, errno.RequestParamErr, gin.H{})
		return
	}
	user := model.User{}
	if err := s.Mysql.DB.Where("username=?", loginReq.Username).First(&user).Error; err != nil {
		response.JSON(c, errno.UsernameNotFound, nil)
		return
	}
	if utils.Md5(loginReq.Password) != user.Password {
		response.JSON(c, errno.PasswordError, nil)
		return
	}
	tokenStr := utils.GenToken(user.ID)
	expiresIn := time.Now().Add(time.Hour * 24 * 7)
	token := model.Token{}
	s.Mysql.DB.Where(model.Token{UserID: user.ID}).
		Assign(model.Token{Token: tokenStr, ExpiresIn: expiresIn}).
		FirstOrCreate(&token)
	response.JSON(c, errno.Success, token)
	return
}

/*
	if err := s.DB.Create(&token).Error; err != nil {
		response.JSON(c, errno.UsernameExisted, nil)
		return
	}
*/
