package adapters

import (
	"github.com/gin-gonic/gin"
)

func ConfigureRouter(eng *gin.Engine,
	ph *AssetsBalancerHandler) {
	v1 := eng.Group("v1")
	v1.POST("assetsGroup", ph.HandleCreateAssetsGroup)
	v1.POST("assetsGroup/asset", ph.HandleCreateAsset)
	v1.PUT("assetsGroup/contributionTotal", ph.HandleUpdateContributionTotal)
	v1.PUT("assetsGroup/asset", ph.HandleUpdateAsset)
	v1.DELETE("assetsGroup/asset", ph.HandleDeleteAsset)
	v1.DELETE("assetsGroup", ph.HandleDeleteAssetsGroup)
	v1.GET("assetsGroup", ph.HandleGetAssetsGroups)
	v1.GET("assetsGroup/:id", ph.HandleGetAssetsGroup)
}
