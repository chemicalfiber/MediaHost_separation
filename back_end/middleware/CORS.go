package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 允许跨域
func CORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

	if c.Request.Method == "OPTIONS" {
		c.String(http.StatusOK, "")
	}

	// 调用下个中间件
	c.Next()
}
