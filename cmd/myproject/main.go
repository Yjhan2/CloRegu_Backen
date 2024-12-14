package main

import (
    "log"
    "myproject/config"
    "myproject/routes"
    "net/http"
    "github.com/rs/cors"
    "myproject/repositories"
    "os"
)

func main() {
    // 加载配置
    config.LoadConfig()

     // 初始化数据库
     dbDSN := os.Getenv("DB_DSN")
    if err := repositories.InitDB(dbDSN); err != nil {
        log.Fatalf("Could not connect to the database: %s\n", err)
    }

    // 设置路由
    router := routes.SetupRouter()

     // 处理 CORS
     c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:8080"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders:   []string{"Content-Type"},
        AllowCredentials: true,
    })

    r := c.Handler(router)

    // 启动服务器
    log.Println("Starting server on :8888")
    if err := http.ListenAndServe(":8888", r); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}