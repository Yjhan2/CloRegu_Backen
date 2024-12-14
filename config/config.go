package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // 使用 os 包获取环境变量
    dbDSN := os.Getenv("DB_DSN")
    if dbDSN == "" {
        log.Fatalf("DB_DSN is not set in the environment variables")
    }
    log.Printf("Database DSN: %s\n", dbDSN)
}