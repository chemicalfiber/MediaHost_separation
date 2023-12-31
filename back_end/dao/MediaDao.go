package dao

import (
	"MediaHost_separation/back_end/models"
	"MediaHost_separation/back_end/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 将文件信息存储到数据库,TODO:重写这部分的逻辑，将文件存储到mongoDB
func SaveFileInfo(media models.Media) (interface{}, error) {
	collection := utils.MongoDB.Collection("media")
	// 向数据库插入数据
	insertResult, err := collection.InsertOne(context.Background(), bson.M{"_id": media.Id,
		"name":        media.Name,
		"upload_date": media.UploadDate,
		"uploader_id": media.UploaderId,
		"type":        media.Type,
		//"local_disk_path": media.LocalDiskPath,
		"grid_fs_key": media.GridFSKey,
	})

	// 当插入出现问题时，返回nil，在调用者处判断返回值是否为nil即可知晓插入是否成功
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, err
}

// 使用文件ID从数据库中查询
func GetFileInfo(id string) (models.Media, error) {
	collection := utils.MongoDB.Collection("media")

	var media models.Media

	result := collection.FindOne(context.Background(), bson.M{"_id": id})
	err := result.Decode(&media)
	// 当查询和转换出现问题时，返回空内容的结构体，在调用者处判断该结构体的任意一个字段是否为"零值"即可知晓是否查询成功
	if err != nil {
		return models.Media{}, err
	}

	return media, nil
}

// 通过用户ID查询该用户的指定类型的媒体文件信息
func GetFilesByUserId(userId string, fileType string) ([]models.Media, error) {
	var medias []models.Media

	collection := utils.MongoDB.Collection("media")
	result, err := collection.Find(context.Background(), bson.M{
		"uploader_id": userId,
		"type":        fileType,
	})
	if err != nil {
		return medias, err
	}
	for result.Next(context.Background()) {
		var media models.Media
		err := result.Decode(&media)
		if err != nil {
			continue
		}
		// 敏感信息保护
		media.GridFSKey = primitive.ObjectID{}
		media.UploaderId = ""
		media.Link = utils.Config.SelfDomain + ":" + utils.Config.Port + "/f/" + media.Id
		medias = append(medias, media)
	}

	return medias, nil
}

// 根据ID删除文件
func DeleteMediaById(id string) (int64, error) {
	collection := utils.MongoDB.Collection("media")

	// 先根据访问id查询到GridFS中对应的ID
	info, err := GetFileInfo(id)
	if err != nil {
		return 0, err
	}

	// 创建数据库连接和存储桶
	fsBucket, err := gridfs.NewBucket(
		utils.MongoDB,
		options.GridFSBucket().SetName("fs"),
	)

	if err != nil {
		return 0, err
	}

	// 删除GridFS中的文件
	err = fsBucket.Delete(info.GridFSKey)
	if err != nil {
		return 0, err
	}

	// 删除文件记录
	result, err := collection.DeleteOne(context.Background(), bson.M{
		"_id": id,
	})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, err
}
