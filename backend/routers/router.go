package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vgbhj/SKAT/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// r.POST("/signup", api.CreateUser)
	r.POST("/login", api.Login)

	return r
}
