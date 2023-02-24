package adapters

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigureRouter(eng *gin.Engine,
	ph *AssetsBalancerHandler) {
	eng.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	v1 := eng.Group("v1")
	v1.POST("assetsGroup", ph.HandleCreateAssetsGroup)
	v1.POST("assetsGroup/asset", ph.HandleCreateAsset)
	v1.PUT("assetsGroup/contributionTotal", ph.HandleUpdateAssetsGroup)
	v1.PUT("assetsGroup/asset", ph.HandleUpdateAsset)
	v1.DELETE("assetsGroup/asset", ph.HandleDeleteAsset)
	v1.DELETE("assetsGroup", ph.HandleDeleteAssetsGroup)
	v1.GET("assetsGroup", ph.HandleGetAssetsGroups)
	v1.GET("assetsGroup/:id", ph.HandleGetAssetsGroup)
}
