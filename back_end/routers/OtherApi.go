package routers

import (
	"MediaHost_separation/back_end/api"
	"MediaHost_separation/back_end/api/userApi"
	"github.com/gin-gonic/gin"
)

func ApiRouteInit(engine *gin.Engine) {
	apiGroup := engine.Group("/api")
	{
		// 生成验证码
		apiGroup.GET("/verifyCode.png", api.GenerateVerifyCode)
		// 查询用户名是否可用
		apiGroup.GET("/checkName", userApi.CheckUserName)
	}
}
