package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

// User 结构体表示用户数据
type User struct {
    ID   int
    Name string
}

func main() {
    // 数据库连接字符串
    dsn := "root:123@tcp(127.0.0.1:3306)/test"
    
    // 连接数据库
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("dsn:%s invalid, err:%v\n", dsn, err)
    }
    defer db.Close()

    // 尝试连接数据库
    err = db.Ping()
    if err != nil {
        log.Fatalf("open %s failed, err:%v\n", dsn, err)
    }
    fmt.Println("连接数据库成功~")

    // 查询用户数据
    rows, err := db.Query("SELECT id, name FROM users")
    if err != nil {
        log.Fatalf("query failed, err:%v\n", err)
    }
    defer rows.Close()

    // 处理查询结果
    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name)
        if err != nil {
            log.Fatalf("scan failed, err:%v\n", err)
        }
        users = append(users, user)
    }

    // 检查是否有错误
    if err = rows.Err(); err != nil {
        log.Fatalf("rows error, err:%v\n", err)
    }

    // 打印用户数据
    if len(users) == 0 {
        fmt.Println("没有查询到用户数据")
    } else {
        for _, user := range users {
            fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Name)
        }
    }
}