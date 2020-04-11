package api

import (
	"qasite/errno"
	"qasite/utils/response"

	"github.com/gin-gonic/gin"
)

func (s *Service) Ping(c *gin.Context) {
	response.JSON(c, errno.Success, gin.H{"message": "pong"})
	return
}
