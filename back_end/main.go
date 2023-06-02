package main

import (
	"MediaHost_separation/back_end/middleware"
	"MediaHost_separation/back_end/routers"
	"MediaHost_separation/back_end/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载项目配置文件
	utils.LoadConfig()
	// 连接mongoDB
	utils.ConnectMongoDB()
	// 创建gin主路由
	engine := gin.Default()

	// CSRF中间件，允许跨域
	engine.Use(middleware.CORS)

	// 路由组注册
	routers.MediaRouteInit(engine)
	routers.UserRouteInit(engine)
	routers.ApiRouteInit(engine)

	//	_ = engine.Run(":8823")
	_ = engine.Run(utils.Config.Address + ":" + utils.Config.Port)

}
