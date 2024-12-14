package services  //业务逻辑层

import (
    "errors"
    "myproject/models"
    "myproject/repositories"
)

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func GetUsers() ([]models.User, error) {
    return repositories.GetUsers()
}

func Authenticate(credentials Credentials) (bool, error) {
    user, err := repositories.GetUserByUsername(credentials.Username)
    if err != nil {
        return false, err
    }
    if user == nil || user.Password != credentials.Password {
        return false, errors.New("invalid credentials")
    }
    // 返回一个模拟的 True 令牌
    return true, nil
}

func CreateUser(user models.User) error {
    // 检查用户数据是否有效
    if user.Username == "" || user.Password == "" {
        return errors.New("invalid user data")
    }
    // 调用 repositories.CreateUser 函数创建用户
    return repositories.CreateUser(user)
}