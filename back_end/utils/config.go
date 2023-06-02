package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var Config AppConfig

/*
配置文件结构体
*/
type AppConfig struct {
	Port            string        `json:"port"`              // 运行在哪个端口上
	Address         string        `json:"address"`           // 监听什么地址
	SelfDomain      string        `json:"self_domain"`       // 自身域名配置，用于返回的关键信息
	FileStoragePath string        `json:"file_storage_path"` // 将上传的文件放在本地的什么位置？
	MongodbConfig   mongodbConfig `json:"mongodb_config"`    // mongoDB配置项
}

// mongoDB配置结构体
type mongodbConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

/*
读取配置文件
*/
func LoadConfig() {
	//var config Config           // 创建Config结构体的实例化对象
	workingDir, _ := os.Getwd() // 得到当前工作目录
	file, err := ioutil.ReadFile(workingDir + string(os.PathSeparator) + "config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &Config)
	if err != nil {
		log.Fatal(err)
	}
	//return config
}
