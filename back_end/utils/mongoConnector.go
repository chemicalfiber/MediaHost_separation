package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client *mongo.Client    // 数据库连接对象
var MongoDB *mongo.Database // 数据库对象，应该直接在连接数据库的地方直接指定要操作的数据库

func ConnectMongoDB() {
	// 建立MongoDB连接
	clientOptions := options.Client().ApplyURI("mongodb://" + Config.MongodbConfig.Host + ":" + Config.MongodbConfig.Port)
	connect, err := mongo.Connect(context.Background(), clientOptions)

	//client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return
	}
	MongoDB = connect.Database(Config.MongodbConfig.Database) // 指定要操作的数据库
	Client = connect
}
