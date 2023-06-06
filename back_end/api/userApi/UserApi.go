package userApi

import (
	"MediaHost_separation/back_end/dao"
	"MediaHost_separation/back_end/models"
	"MediaHost_separation/back_end/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
与用户相关的API
*/

// 注册	// TODO：返回一些信息，让前端的catch捕捉并显示错误
func Register(c *gin.Context) {
	username := c.PostForm("username")
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirmPassword")

	// 检查前端传递过来的"密码"和"确认密码"是否相等
	if confirmPassword != password {
		utils.BadRequest400(c, "两次输入的密码不一致！", gin.H{
			"username": username,
			"nickname": nickname,
		})
		return
	}

	// 检查用户名是否重复
	getUserByUsername := dao.GetUserByUsername(username)
	if getUserByUsername.ID != "" {
		// 用户名重复的情况，返回错误信息
		utils.BadRequest400(c, "用户名被占用！", gin.H{
			"username":            username,
			"nickname":            nickname,
			"usernameIsBeingUsed": true, // 用户名被占用了！
		})
		return
	}

	user := models.User{
		Username: username,
		Nickname: nickname,
		Password: password,
		//UUID:     utils.GenerateUUID(),
	}

	result := dao.SaveUser(user)
	if result != nil {
		// 注册通过！回到登录页面，并让用户重新登录
		//c.Redirect(http.StatusFound,"/login")
		//c.HTML(http.StatusOK, "loginAndRegister.html", gin.H{
		//	"registerSuccess": true,
		//})
		utils.Ok200(c, nil)
	}
}

// 登录
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 检查参数是否被正确传递
	if username == "" || password == "" {
		utils.BadRequest400(c, "用户名或密码错误", gin.H{
			"isLoginError": true,
		})
		return
	}

	// 从数据库查询用户
	user := dao.GetUserByUsernameAndPassword(username, password)
	if user.ID == "" {
		utils.BadRequest400(c, "用户名或密码错误", gin.H{
			"isLoginError": true,
		})
		return
	}

	token, err := utils.GenerateToken(user) // 生成JWT token
	if err != nil {
		utils.InternalError500(c, err)
	}

	utils.Ok200(c, gin.H{
		"token": token,
	})
}

// 修改
func Update(c *gin.Context) {
	// 确认是修改当前用户，从token中取出用户id
}

// 登出
func Logout(c *gin.Context) {

}

// 检验用户名是否重复
func CheckUserName(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"canUse": false,
		})
		return
	}

	user := dao.GetUserByUsername(username)
	if user.ID == "" { // 找不到用户，说明用户名未被占用
		c.JSON(http.StatusOK, gin.H{
			"canUse": true,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"canUse": false,
	})
}

// TODO：用户首页（尚未被其他路由指定）
func HomePage(c *gin.Context) {

}

// TODO：根据UUID查询用户（尚未被其他路由指定）
func FindById(c *gin.Context) {

}

// 检查token是否有效
func CheckToken(c *gin.Context) {
	tokenString := c.Query("token")
	_, err := utils.DecodeToken(tokenString)

	// token验证不通过
	if err != nil {
		utils.BadRequest400(c, "token失效！", nil)
		return
	}
	// token验证通过
	utils.Ok200(c, nil)
}
