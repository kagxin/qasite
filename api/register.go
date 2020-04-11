package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"qasite/errno"
	"qasite/model"
	"qasite/utils"
	"qasite/utils/response"
)

func (s *Service) Register(c *gin.Context) {
	var regUser model.LoginReq
	if err := c.ShouldBind(&regUser); err != nil {
		log.Info(err)
		response.JSON(c, errno.RequestParamErr, nil)
		return
	}
	user := model.User{}
	if !s.DB.Where("username=?", regUser.Username).First(&user).RecordNotFound() {
		response.JSON(c, errno.UsernameExisted, nil)
		return
	}
	pwd := utils.Md5(regUser.Password)

	if err := s.DB.Where(model.User{Username: regUser.Username}).Assign(model.User{Password: pwd}).FirstOrCreate(&user).Error; err != nil {
		response.JSON(c, errno.ServerError, nil)
		return
	}
	response.JSON(c, errno.Success, model.UserRsp{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
	return

}
