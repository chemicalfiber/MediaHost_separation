package mediaApi

import (
	"MediaHost_separation/back_end/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/*
与媒体文件相关的API
*/

// 获取指定用户上传的视频
func GetVideosByUserId(c *gin.Context) {

}

// 获取指定用户上传的图片
func GetImagesByUserId(c *gin.Context) {

}

// 接收上传文件
func UploadHandler(c *gin.Context) {
	// 需要提供token才允许上传文件
	header := c.GetHeader("x-token")
	userTokenString := c.PostForm("x-token")
	if header == "" && userTokenString == "" {
		utils.BadRequest400(c, "未传递token，上传被拒绝", nil)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		utils.BadRequest400(c, err.Error(), nil)
	}

	// 获取前端传过来的文件MIME_TYPE，在后续操作该文件的字段
	contentType := file.Header["Content-Type"]
	fmt.Println(contentType)
	var fileType string
	if strings.Contains(contentType[0], "image") {
		fileType = "image"
	} else if strings.Contains(contentType[0], "video") {
		fileType = "video"
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "本图床系统仅接受图片/视频类型的文件",
		})
		return
	}
	fmt.Println(fileType)

}

// TODO：提供文件下载（尚未被其他路由指定）
func DownloadHandler(c *gin.Context) {

}
