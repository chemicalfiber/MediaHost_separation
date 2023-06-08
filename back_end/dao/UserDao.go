package dao

import (
	"MediaHost_separation/back_end/models"
	"MediaHost_separation/back_end/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 将用户信息存储到数据库
func SaveUser(user models.User) interface{} {
	collection := utils.MongoDB.Collection("user")

	result, err := collection.InsertOne(context.Background(), bson.M{
		"username": user.Username,
		"nickname": user.Nickname,
		//"uuid":     user.UUID,
		"password": user.Password,
	})

	if err != nil {
		return nil
	}

	// 将用户ID返回，在调用处判断是否为nil即可得知是否插入成功
	return result.InsertedID
}

// 更新用户
func UpdateUser(user models.User) interface{} {
	collection := utils.MongoDB.Collection("user")

	// 转换string类型的ID为ObjectID
	hex, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return nil
	}
	filter := bson.M{"_id": hex}
	update := bson.M{"$set": bson.M{
		"nickname": user.Nickname,
		"password": user.Password,
	}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil
	}
	// 将影响的行数返回，在调用处判断是否为nil即可得知是否修改成功
	return result.ModifiedCount
}

// 使用用户名和密码查询用户
func GetUserByUsernameAndPassword(username string, password string) models.User {
	collection := utils.MongoDB.Collection("user")

	var user models.User

	// 从数据库查询
	result := collection.FindOne(context.Background(), bson.M{
		"username": username,
		"password": password,
	})
	err := result.Decode(&user)
	if err != nil {
		// 将用户结构体返回，在调用处判断用户ID是否为nil即可得知是否插入成功
		return models.User{}
	}

	return user
}

// 根据用户名获取用户
func GetUserByUsername(username string) models.User {
	collection := utils.MongoDB.Collection("user")

	var user models.User

	result := collection.FindOne(context.Background(), bson.M{"username": username})
	err := result.Decode(&user)
	if err == nil { // decode不发生错误，则说明数据库中有对应用户使用了这个用户名
		return user
	}
	return models.User{}
}
