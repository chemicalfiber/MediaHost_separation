package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// "媒体文件"结构体
type Media struct {
	Id         string             `json:"_id" bson:"_id"`                 // 文件的唯一id，62进制数字
	Name       string             `json:"name" bson:"name"`               // 文件名
	UploadDate time.Time          `json:"upload_date" bson:"upload_date"` // 上传日期
	UploaderId string             `json:"uploader_id" bson:"uploader_id"` // 上传者ID
	Type       string             `json:"type" bson:"type"`               // 媒体类型，图片还是视频？ 这里是MIME type的参考 http://blog.haoji.me/mime-type.html
	GridFSKey  primitive.ObjectID `json:"grid_fs_key" bson:"grid_fs_key"` // 存储在GridFS中的ObjectId
	//LocalDiskPath string    `json:"local_disk_path" bson:"local_disk_path"` // 存储在本地磁盘的路径
}
