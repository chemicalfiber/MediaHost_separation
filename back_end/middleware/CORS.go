package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 允许跨域
func CORS(c *gin.Context) {
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	//c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
	//
	//if c.Request.Method == "OPTIONS" {
	//	c.String(http.StatusOK, "")
	//}
	//
	//// 调用下个中间件
	//c.Next()
	method := c.Request.Method

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	c.Next()
}
