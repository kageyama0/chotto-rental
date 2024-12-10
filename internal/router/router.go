package router

import (
	"os"

	"github.com/gin-gonic/gin"
	application_handler "github.com/kageyama0/chotto-rental/internal/handler/application"
	auth_handler "github.com/kageyama0/chotto-rental/internal/handler/auth"
	case_handler "github.com/kageyama0/chotto-rental/internal/handler/case"
	matching_handler "github.com/kageyama0/chotto-rental/internal/handler/matching"
	review_handler "github.com/kageyama0/chotto-rental/internal/handler/review"
	user_handler "github.com/kageyama0/chotto-rental/internal/handler/user"
	"github.com/kageyama0/chotto-rental/pkg/auth"
	"github.com/kageyama0/chotto-rental/pkg/middleware"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	authService := auth.NewAuthService(os.Getenv("JWT_SECRET"))
	applicationHandler := application_handler.NewApplicationHandler(db)
	authHandler := auth_handler.NewAuthHandler(db, authService)
	caseHandler := case_handler.NewCaseHandler(db)
	matchingHandler := matching_handler.NewMatchingHandler(db)
	reviewHandler := review_handler.NewReviewHandler(db)
	userHandler := user_handler.NewUserHandler(db)

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
