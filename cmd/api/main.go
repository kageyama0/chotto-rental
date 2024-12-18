package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kageyama0/chotto-rental/config"
	"github.com/kageyama0/chotto-rental/internal/model"
	"github.com/kageyama0/chotto-rental/internal/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/kageyama0/chotto-rental/docs" // ここで docs を import
	"gorm.io/gorm/logger"
)

func initDB(config config.DatabaseConfig) *gorm.DB {
	host := config.Host
	user := config.User
	password := config.Password
	name := config.Name
	port := config.Port
	sslmode := config.SSLMode
	timezone := config.Timezone
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, name, port, sslmode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// ログを出したい時だけでいいので、コメントアウト
	db.Config.Logger = logger.Default.LogMode(logger.Info)

	// マイグレーションは、必要な時だけでいいので、一旦コメントアウト
	// TODO: 設定で変えれるようにする
	if err := model.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}


// @title Chotto Rental API
// @version 1.0
// @description ちょっとレンタルサービスのAPI仕様書
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description 「Bearer 」の後にJWTトークンを付与してください
// @schemes http
func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db := initDB(config.Database)
	r := router.SetupRouter(db, config)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
