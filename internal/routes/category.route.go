package routes

import (
	"github.com/fajarhidayad/waow-article/internal/handlers"
	"github.com/fajarhidayad/waow-article/internal/repositories"
	"github.com/fajarhidayad/waow-article/internal/services"
	"github.com/fajarhidayad/waow-article/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryRoutes(r *gin.RouterGroup, db *gorm.DB) {
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	route := r.Group("/categories", middleware.HasAccessToken(), middleware.HasRoleUser())

	route.POST("/", categoryHandler.Create)
	route.GET("/", categoryHandler.FindAll)
	route.GET("/:id", categoryHandler.FindByID)
	route.PUT("/:id", categoryHandler.Update)
	route.DELETE("/:id", categoryHandler.Delete)
}
