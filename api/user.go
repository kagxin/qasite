package api

import (
	"github.com/gin-gonic/gin"
	"qasite/errno"
	"qasite/model"
	"qasite/utils/response"
)

func (s *Service) UserInfo(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(model.User)
	response.JSON(c, errno.Success, model.UserRsp{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
	return
}
