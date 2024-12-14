package repositories  // 数据访问层

import (
    "database/sql"
    "myproject/models"
    "os"
    "log"
    "errors"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(dataSourceName string) error {
    var err error
    db, err = sql.Open("mysql", dataSourceName) // 使用 "mysql" 作为驱动程序名称
    if err != nil {
        log.Println("Error opening database:", err)
        return err
    }
    return db.Ping()
}

func GetUsers() ([]models.User, error) {
    dsn := os.Getenv("DB_DSN")
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    defer db.Close()

    rows, err := db.Query("SELECT id, username FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.ID, &user.Username)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

func GetUserByUsername(username string) (*models.User, error) {
    var user models.User
    err := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}


func CreateUser(user models.User) error {
    if db == nil {
        return errors.New("database connection is not initialized")
    }

   // 检查用户名是否已存在
   existingUser, err := GetUserByUsername(user.Username)
    if err != nil {
        return err
    }
    if existingUser != nil {
        return errors.New("用户已存在")
    }

    _, err = db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", user.Username, user.Password, user.Email)
    if err != nil {
        log.Println("Error inserting user:", err)
    }

    return err
}