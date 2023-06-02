package routers

import (
	"MediaHost_separation/back_end/api/mediaApi"
	"github.com/gin-gonic/gin"
)

func MediaRouteInit(engine *gin.Engine) {
	mediaRouter := engine.Group("/file")
	// TODO：使用一条404路由规则
	//fileRouter.GET("/",handlers.ToFileUpload)
	{
		mediaRouter.POST("/upload", mediaApi.UploadHandler)
	}
}
