package routers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vgbhj/SKAT/docs"
	"github.com/vgbhj/SKAT/middleware/jwt"
	"github.com/vgbhj/SKAT/routers/api"
	v1 "github.com/vgbhj/SKAT/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	distPath := filepath.Join("..", "frontend", "dist")
	r.StaticFS("/assets", http.Dir(filepath.Join(distPath, "assets")))
	r.StaticFile("/", filepath.Join(distPath, "index.html"))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.POST("/signup", api.CreateUser)
	r.POST("/login", api.Login)

	apiv1 := r.Group("/api/v1")

	apiv1.GET("/materials", v1.GetMaterials)
	apiv1.GET("/materials/:id", v1.GetMaterial)
	apiv1.GET("/materials/:id/download", v1.DownloadMaterial)

	apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/materials", v1.AddMaterial)
	}

	r.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(distPath, "index.html"))
	})
	return r
}
