package api

import (
	"github.com/gin-gonic/gin"
	"qasite/errno"
	"qasite/model"
	"qasite/utils/response"
)

func (s *Service) CreateCommit(c *gin.Context) {
	questionId := c.Param("QID")

	if s.DB.Where("id=?", questionId).First(&model.Question{}).RecordNotFound() {
		response.JSON(c, errno.NotFound, nil)
		return
	}

}
