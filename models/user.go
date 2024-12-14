package models  // 定义数据模型

type User struct {
    ID int `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    Email string `json:"email"`
}