package controllers  // 处理HTTTP请求控制器

import (
    "encoding/json"
    "myproject/models"
    "myproject/services"
    "net/http"
    "log"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
    users, err := services.GetUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(users)
}

// 登录处理函数
func Login(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request for Login")
    var credentials services.Credentials
    err := json.NewDecoder(r.Body).Decode(&credentials)
    if err != nil {
        log.Println("Error decoding user:", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // 打印解码出的用户信息
    log.Printf("Decoded user: %+v\n", credentials)

    token, err := services.Authenticate(credentials)
    if err != nil {
        log.Println("Error creating user:", err)
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(map[string]interface{}{"success": token})
}

// 注册处理函数
func Register(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request for Register")
    var user models.User
    //解析用户数据
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        log.Println("Error decoding user:", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // 打印解码出的用户信息
    log.Printf("Decoded user: %+v\n", user)

    err = services.CreateUser(user)
    if err != nil {
        if err.Error() == "用户已存在" {
            log.Println("User already exists:", user.Username)
            json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "message": "注册失败，邮箱或用户名已存在"})
            return
        }
        log.Println("Error creating user:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Println("User registered successfully:", user.Username)
    json.NewEncoder(w).Encode(map[string]interface{}{"success": true})
}