package routes

import (
	"github.com/fajarhidayad/waow-article/internal/handlers"
	"github.com/fajarhidayad/waow-article/internal/repositories"
	"github.com/fajarhidayad/waow-article/internal/services"
	"github.com/fajarhidayad/waow-article/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	adminService := services.NewAdminService(userRepo)
	adminHandler := handlers.NewAdminHandler(adminService)

	admin := router.Group("/admin", middleware.HasAccessToken(), middleware.HasRoleAdmin())

	users := admin.Group("users")
	{
		users.POST("/", adminHandler.CreateUser)
		users.GET("/", adminHandler.FindAllUsers)
		users.GET("/:id", adminHandler.FindUserByID)
		users.PUT("/:id", adminHandler.UpdateUser)
		users.DELETE("/:id", adminHandler.DeleteUser)
	}
}
