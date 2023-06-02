package utils

import "github.com/google/uuid"

// 为新用户生产UUID，或者用户请求重新生成UUID时调用的函数
func GenerateUUID() string {
	return uuid.New().String()
}
