package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	//_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

var DB *sql.DB

// Initialize データベース接続　初期化
func Initialize(config Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	// 接続プールの設定
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	// 接続確認
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	DB = db
	log.Println("Database connection established")
	return nil
}

// Close データベース接続を閉じる
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
