package routes

import (
	"github.com/fajarhidayad/waow-article/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ArticleRoutes(r *gin.RouterGroup, db *gorm.DB) {

	route := r.Group("/articles", middleware.HasAccessToken(), middleware.HasRoleUser())

	route.POST("/", func(context *gin.Context) {

	})
	route.GET("/", func(context *gin.Context) {

	})
	route.GET("/:id", func(context *gin.Context) {

	})
	route.PUT("/:id", func(context *gin.Context) {

	})
	route.DELETE("/:id", func(context *gin.Context) {

	})
}
