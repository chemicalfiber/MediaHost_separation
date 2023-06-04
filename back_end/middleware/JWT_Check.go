package middleware

import (
	"MediaHost_separation/back_end/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckJWT(c *gin.Context) {
	tokenString := c.GetHeader("x-token")
	token, err := utils.DecodeToken(tokenString)
	if err != nil {
		utils.BadRequest400(c, "token失效！", err)
		c.Abort()
		return
	}
	fmt.Println(token.Claims.(*utils.MyCustomClaims))
	// 调用下个中间件
	c.Next()
}
