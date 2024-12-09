package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
    godotenv.Load()
}

func initDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// ヘルスチェック
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API グループ
	api := r.Group("/api")
	{
		// ユーザー関連
		users := api.Group("/users")
		{
			users.POST("/register", nil) // TODO
			users.POST("/login", nil)    // TODO
		}

		// 認証が必要なエンドポイント
		auth := api.Group("/")
		{
			// 案件関連
			cases := auth.Group("/cases")
			{
				cases.GET("", nil)          // 一覧取得
				cases.POST("", nil)         // 作成
				cases.GET("/:id", nil)      // 詳細取得
				cases.PUT("/:id", nil)      // 更新
				cases.DELETE("/:id", nil)   // 削除
			}

			// 応募関連
			applications := auth.Group("/applications")
			{
				applications.POST("", nil)      // 応募する
				applications.GET("", nil)       // 応募一覧取得
				applications.PUT("/:id", nil)   // 応募ステータス更新
			}

			// マッチング関連
			matchings := auth.Group("/matchings")
			{
				matchings.POST("", nil)                     // マッチング作成
				matchings.PUT("/:id/confirm-arrival", nil)  // 到着確認
			}

			// レビュー関連
			reviews := auth.Group("/reviews")
			{
				reviews.POST("", nil)      // レビュー投稿
				reviews.GET("", nil)       // レビュー一覧取得
			}
		}
	}

	return r
}

func main() {
	db := initDB()
	r := setupRouter(db)
	r.Run(":8080")
}
