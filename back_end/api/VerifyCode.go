package api

import (
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateVerifyCode(c *gin.Context) {
	id := captcha.NewLen(5)

	fmt.Println("生成的验证码ID是：" + id)

	// TODO：移除此处的SESSION，让前端携带验证码ID
	session := sessions.Default(c)
	session.Set("verifyCodeId", id)
	err := session.Save()
	if err != nil {
		c.String(http.StatusInternalServerError, "服务器内部错误！")
		return
	}
	err = captcha.WriteImage(c.Writer, id, 120, 80)
	if err != nil {
		c.String(http.StatusInternalServerError, "服务器内部错误！")
		return
	}
}
