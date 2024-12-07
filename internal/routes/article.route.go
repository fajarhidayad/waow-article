package routes

import (
	"github.com/fajarhidayad/waow-article/internal/handlers"
	"github.com/fajarhidayad/waow-article/internal/repositories"
	"github.com/fajarhidayad/waow-article/internal/services"
	"github.com/fajarhidayad/waow-article/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ArticleRoutes(r *gin.RouterGroup, db *gorm.DB) {
	articleRepo := repositories.NewArticleRepository(db)
	userRepo := repositories.NewUserRepository(db)
	articleService := services.NewArticleService(articleRepo, userRepo)
	articleHandler := handlers.NewArticleHandler(articleService)

	route := r.Group("/articles", middleware.HasAccessToken(), middleware.HasRoleUser())

	route.POST("/", articleHandler.Create)
	route.GET("/", articleHandler.FindAll)
	route.GET("/:id", articleHandler.FindById)
	route.PUT("/:id", articleHandler.Update)
	route.DELETE("/:id", articleHandler.Delete)
}
