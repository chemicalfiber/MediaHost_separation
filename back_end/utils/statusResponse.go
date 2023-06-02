package utils

/*自定义响应封装*/
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok200(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    v,
	})
}

func InternalError500(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
	})
}

func NotFound404(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": err.Error(),
	})
}
func BadRequest400(c *gin.Context, errMessage string) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": errMessage,
	})
}
