package userApi

import (
	"MediaHost_separation/back_end/utils"
	"github.com/gin-gonic/gin"
)

/*
与用户相关的API
*/

// 注册
func Register(c *gin.Context) {
	username := c.PostForm("username")
	utils.Ok200(c, username)
}

// 登录
func Login(c *gin.Context) {

}

// 修改
func Update(c *gin.Context) {

}

// 登出
func Logout(c *gin.Context) {

}

// 检验用户名是否重复
func CheckUserName(c *gin.Context) {

}

// TODO：用户首页（尚未被其他路由指定）
func HomePage(c *gin.Context) {

}

// TODO：根据UUID查询用户（尚未被其他路由指定）
func FindById(c *gin.Context) {

}
