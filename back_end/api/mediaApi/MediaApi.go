package mediaApi

import (
	"MediaHost_separation/back_end/dao"
	"MediaHost_separation/back_end/models"
	"MediaHost_separation/back_end/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"net/http"
	"strings"
	"time"
)

/*
与媒体文件相关的API
*/

// 获取指定用户上传的视频
func GetVideosByUserId(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		utils.BadRequest400(c, "未提供用户API token", nil)
		return
	}

	videos, err := dao.GetFilesByUserId(userId, "video")
	if err != nil {
		utils.BadRequest400(c, err.Error(), nil)
		return
	}

	utils.Ok200(c, videos)
}

// 获取指定用户上传的图片
func GetImagesByUserId(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		utils.BadRequest400(c, "未提供用户API token", nil)
		return
	}

	images, err := dao.GetFilesByUserId(userId, "image")
	if err != nil {
		utils.BadRequest400(c, err.Error(), nil)
		return
	}

	utils.Ok200(c, images)
}

// 获取文件详情
func GetMediaDetail(c *gin.Context) {
	fileId := c.Param("id")
	//fmt.Println(fileId)
	fileInfo, err := dao.GetFileInfo(fileId)
	if err != nil {
		utils.BadRequest400(c, "文件ID无效", nil)
		return
	}

	user := dao.GetUserById(fileInfo.UploaderId)
	if user.ID == "" {
		utils.BadRequest400(c, "找不到对应用户", nil)
		return
	}
	// 安全处理，屏蔽关键信息
	fileInfo.GridFSKey = primitive.ObjectID{}
	fileInfo.UploaderId = ""

	//fmt.Println(info)
	utils.Ok200(c, gin.H{
		"upload_user": user.Nickname,
		"info":        fileInfo,
		"file_link":   utils.Config.SelfDomain + ":" + utils.Config.Port + "/f/" + fileInfo.Id,
	})
}

// 接收上传文件
func UploadHandler(c *gin.Context) {
	// 需要提供token才允许上传文件
	userTokenString := c.GetHeader("x-token")
	if userTokenString == "" {
		userTokenString = c.PostForm("x-token")
	}
	userId := c.PostForm("userId") // 如果没有token，就以传递过来的userId作为写入数据库的依据

	if userTokenString == "" && userId == "" {
		utils.BadRequest400(c, "未传递token或用户id，上传被拒绝", nil)
		return
	}

	// 以token中的id为准，忽略请求参数中的userId
	if userTokenString != "" {
		token, err := utils.DecodeToken(userTokenString)
		if err != nil {
			utils.BadRequest400(c, err.Error(), nil)
		}
		userId = token.Claims.(*utils.MyCustomClaims).UserId
	}
	//username := token.Claims.(*utils.MyCustomClaims).Username

	file, err := c.FormFile("file")
	if err != nil {
		utils.BadRequest400(c, err.Error(), nil)
	}

	// 获取前端传过来的文件MIME_TYPE，在后续操作该文件的字段
	contentType := file.Header["Content-Type"]
	fmt.Println("Content-Type", contentType)
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
	fmt.Println("文件类型：", fileType)
	fmt.Println("上传者ID：", userId)
	//fmt.Println("上传者昵称：", username)

	fmt.Println("---------------")

	// 创建数据库连接和存储桶
	fsBucket, err := gridfs.NewBucket(
		utils.MongoDB,
		options.GridFSBucket().SetName("fs"),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误"})
		return
	}

	// 打开上传流
	uploadStream, err := fsBucket.OpenUploadStream(file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误"})
		return
	}
	defer uploadStream.Close()

	// 打开上传的文件
	sourceFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "文件打开失败"})
		return
	}
	defer sourceFile.Close()

	// 写入文件
	_, err = io.Copy(uploadStream, sourceFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "文件写入失败"})
		return
	}

	// 创建文件上传时间，用于保存到数据库中，同时修正时区问题
	parsedTime, err := time.Parse("2006-01-02 15:04:05.000", time.Now().Format("2006-01-02 15:04:05.000"))
	fmt.Println(parsedTime)
	if err != nil {
		fmt.Println(err)
	}
	// 创建文件信息结构体，用于保存到数据库中
	media := models.Media{
		Id:         utils.Encode10to62(time.Now().UnixNano()),
		Name:       file.Filename,
		UploadDate: parsedTime,
		UploaderId: userId,                                   // uploader_id"是当前登录用户（session中）的id，或参数传递的ID，以登录用户的ID优先
		Type:       fileType,                                 // 文件类型
		GridFSKey:  uploadStream.FileID.(primitive.ObjectID), // 存储在GridFS中的ObjectId
	}

	// 写入数据库
	insertId, err := dao.SaveFileInfo(media)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	utils.Ok200(c, gin.H{
		"fileLink": utils.Config.SelfDomain + ":" + utils.Config.Port + "/f/" + insertId.(string),
		//"fileInfoKey": insertId,                                          // 文件信息id
		//"fileKey":     uploadStream.FileID.(primitive.ObjectID).String(), // GridFS桶中的文件id
	})
}

// 提供文件下载（尚未被其他路由指定）
func DownloadHandler(c *gin.Context) {
	fileId := c.Param("id") // REST ful查询
	if fileId == "" {
		//c.JSON(http.StatusBadRequest, gin.H{"message": "未提供文件ID"})
		utils.BadRequest400(c, "无效的文件ID", nil)
		return
	}

	// 从数据库查询文件对应的GridFS ID
	media, err := dao.GetFileInfo(fileId)
	if err != nil {
		utils.BadRequest400(c, "无效的文件ID", nil)
		return
	}

	gridFSKey := media.GridFSKey.String()[10:34]
	fmt.Println(gridFSKey)
	objectID, err := primitive.ObjectIDFromHex(gridFSKey)
	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"message": "无效的文件ID"})
		utils.BadRequest400(c, "无效的文件ID", nil)
		return
	}

	// 创建数据库连接和存储桶
	fsBucket, err := gridfs.NewBucket(
		utils.MongoDB,
		options.GridFSBucket().SetName("fs"),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误"})
		return
	}

	// 创建GridFS下载流
	downloadStream, err := fsBucket.OpenDownloadStream(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误，文件下载失败"})
		return
	}
	defer downloadStream.Close()

	// 获取文件信息
	file := downloadStream.GetFile()
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "文件信息获取失败"})
	//	return
	//}

	// 将文件数据写入ResponseWriter
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, file.Name))
	_, err = io.Copy(c.Writer, downloadStream)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误，文件下载失败"})
		return
	}
}
