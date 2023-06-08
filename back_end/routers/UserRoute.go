package routers

import (
	"MediaHost_separation/back_end/api/mediaApi"
	"MediaHost_separation/back_end/api/userApi"
	"github.com/gin-gonic/gin"
)

func UserRouteInit(engine *gin.Engine) {
	userGroup := engine.Group("/user")
	//userGroup.GET("/", handlers.UserMainPage) // 当访问「/user」下的子目录匹配不到路由时，就回落到这个路由
	{
		// 登录相关
		userGroup.POST("/login", userApi.Login)

		// 注册相关
		userGroup.POST("/register", userApi.Register)

		// 修改用户
		userGroup.POST("/update", userApi.Update)

		// "我的媒体"
		userMediaGroup := userGroup.Group("/media")
		userMediaGroup.GET("/", mediaApi.GetImagesByUserId) // 当访问「/user/media」下的子目录匹配不到路由时，就回落到这个路由
		{
			userMediaGroup.GET("/img", mediaApi.GetImagesByUserId)
			userMediaGroup.GET("/video", mediaApi.GetVideosByUserId)
		}
	}

}
