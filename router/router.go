package router

import (
	"qasite/api"

	"github.com/gin-gonic/gin"
)

// Router Init
func Router(srv *api.Service) *gin.Engine {
	r := gin.Default()
	r.GET("ping", srv.ping())
	return r
}
