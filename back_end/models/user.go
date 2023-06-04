package models

// “用户”结构体
type User struct {
	ID       string `json:"_id" bson:"_id"`           // 用户唯一id
	Username string `json:"username" bson:"username"` // 用户名（登录用）
	Nickname string `json:"nickname" bson:"nickname"` // 用户昵称（显示用）
	Password string `json:"password" bson:"password"` // 密码
	Usertype string `json:"usertype" bson:"usertype"` // 用户类型（是否为管理员？）
	// UUID     string `json:"uuid"`           // 用户token（在其他地方调用本图床的上传API时需要传递），有了ObjectID后不需要uuid了
}
