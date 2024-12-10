package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	// "github.com/kageyama0/chotto-rental/internal/model"
	"github.com/kageyama0/chotto-rental/internal/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

func initDB() *gorm.DB {
	// TODO: 設定値の読み込みは、configフォルダ内でやる
	// dsn := os.Getenv("DB_URL")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := "disable"
	timezone := "Asia/Tokyo"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// ログを出したい時だけでいいので、コメントアウト
	// db.Config.Logger = logger.Default.LogMode(logger.Info)

	// マイグレーションは、必要な時だけでいいので、一旦コメントアウト
	// model.Migrate(db)

	return db
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	db := initDB()
	r := router.SetupRouter(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
