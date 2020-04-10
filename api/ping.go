package api

import (
	"qasite/errno"
	"qasite/utils/response"

	"github.com/gin-gonic/gin"
)

func (s *Service) ping(c *gin.Context) {
	response.JSON(c, errno.Success, gin.H{"hello": "world"})
	return
}