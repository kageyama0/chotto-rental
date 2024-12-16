package router

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	application_handler "github.com/kageyama0/chotto-rental/internal/handler/application"
	auth_handler "github.com/kageyama0/chotto-rental/internal/handler/auth"
	case_handler "github.com/kageyama0/chotto-rental/internal/handler/case"
	matching_handler "github.com/kageyama0/chotto-rental/internal/handler/matching"
	review_handler "github.com/kageyama0/chotto-rental/internal/handler/review"
	user_handler "github.com/kageyama0/chotto-rental/internal/handler/user"
	"github.com/kageyama0/chotto-rental/pkg/middleware"
	"github.com/kageyama0/chotto-rental/pkg/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// dev用の設定
	r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"}, // フロントエンドのURL
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
	}))

	authService := service.NewAuthService(os.Getenv("JWT_SECRET"))
	applicationHandler := application_handler.NewApplicationHandler(db)
	authHandler := auth_handler.NewAuthHandler(db, authService)
	caseHandler := case_handler.NewCaseHandler(db)
	matchingHandler := matching_handler.NewMatchingHandler(db)
	reviewHandler := review_handler.NewReviewHandler(db)
	userHandler := user_handler.NewUserHandler(db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	{
		// 認証不要のエンドポイント
		api.POST("/auth/signup", authHandler.Signup)
		api.POST("/auth/login", authHandler.Login)

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
				cases.GET("/:case_id", caseHandler.Get)
				cases.PUT("/:case_id", caseHandler.Update)
				cases.DELETE("/:case_id", caseHandler.Delete)

				applications := cases.Group("/applications")
				{
					applications.POST("", applicationHandler.Create)
					applications.GET("", applicationHandler.List)
					applications.PUT("/:application_id/status", applicationHandler.UpdateStatus)
				}

				matchings := cases.Group("/matchings")
				{
					matchings.POST("", matchingHandler.Create)
					matchings.POST("/:matching_id/confirm-arrival", matchingHandler.ConfirmArrival)
					reviews := matchings.Group("/review")
					{
						reviews.POST("", reviewHandler.Create)
						reviews.GET("", reviewHandler.List)
					}
				}
			}
		}
	}

	return r
}
