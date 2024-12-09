package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kageyama0/chotto-rental/internal/auth"
	"github.com/kageyama0/chotto-rental/internal/handler"
	"github.com/kageyama0/chotto-rental/internal/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := os.Getenv("DEV_DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	authService := auth.NewAuthService(os.Getenv("JWT_SECRET"))
	applicationHandler := handler.NewApplicationHandler(db)
	authHandler := handler.NewAuthHandler(db, authService)
	caseHandler := handler.NewCaseHandler(db)
	matchingHandler := handler.NewMatchingHandler(db)
	reviewHandler := handler.NewReviewHandler(db)
	userHandler := handler.NewUserHandler(db)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	{
		// 認証不要のエンドポイント
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)

		// 認証が必要なエンドポイント
		auth := api.Group("", middleware.AuthMiddleware(authService))
		{
			users := auth.Group("/users")
			{
				users.GET("/:id", userHandler.Get)
				users.PUT("/:id", userHandler.Update)
				users.DELETE("/:id", userHandler.Delete)
				users.GET("/:id/reviews", userHandler.GetReviews)
			}

			cases := auth.Group("/cases")
			{
				cases.POST("", caseHandler.Create)
				cases.GET("", caseHandler.List)
				cases.GET("/:id", caseHandler.Get)
				cases.PUT("/:id", caseHandler.Update)
				cases.DELETE("/:id", caseHandler.Delete)
			}

			applications := auth.Group("/applications")
			{
				applications.POST("", applicationHandler.Create)
				applications.GET("", applicationHandler.List)
				applications.PUT("/:id/status", applicationHandler.UpdateStatus)
			}

			matchings := auth.Group("/matchings")
			{
				matchings.POST("", matchingHandler.Create)
				matchings.POST("/:id/confirm-arrival", matchingHandler.ConfirmArrival)
			}

			reviews := auth.Group("/reviews")
			{
				reviews.POST("", reviewHandler.Create)
				reviews.GET("", reviewHandler.List)
			}
		}
	}

	return r
}

func main() {
	db := initDB()
	r := setupRouter(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
