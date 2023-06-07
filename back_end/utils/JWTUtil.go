package utils

import (
	"MediaHost_separation/back_end/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 指定加密JWT token的"盐"
var SECRET = "aetvhilebhsuihbuvszhbvuabguasvbhuilhervnulgahlvubu"

// 自定义Claims结构体
type MyCustomClaims struct {
	// 自定义字段
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Usertype string `json:"usertype"`
	UserId   string `json:"user_id"`
	// 实现了jwt Claims接口的结构体
	jwt.StandardClaims
}

// 生成token
func GenerateToken(user models.User) (string, error) {
	mySigningKey := []byte(SECRET) // 指定加密的"盐"

	// Create the Claims
	claims := MyCustomClaims{
		user.Username,
		user.Nickname,
		user.Usertype,
		user.ID,
		jwt.StandardClaims{
			//NotBefore: time.Now().UnixNano(),
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间，为unix时间，往后7天
			Issuer:    "MediaHostBackEnd",             // 当前token签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 第一个参数是加密算法，第二个参数是payload
	tokenString, err := token.SignedString(mySigningKey)       // 使用"盐"进行签名，这里需要提前将"盐"转换为byte数组

	//fmt.Printf("%v\n %v\n", tokenString,err)

	return tokenString, err
}

// 解析token
func DecodeToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	//	fmt.Println(token)	// 输出解析后的token结构体实例
	//	fmt.Println(token.Claims)	// 输出解析后的claims结构体实例
	// fmt.Println(token.Claims.(*MyCustomClaims))	// 类型断言
	//	fmt.Println(token.Claims.(*MyCustomClaims).Username)	// 取出我自定义MyCustomClaims中定义的字段
	return token, err
}
