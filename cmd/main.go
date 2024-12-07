package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fajarhidayad/waow-article/internal/routes"
	"github.com/fajarhidayad/waow-article/pkg/database"
	"github.com/fajarhidayad/waow-article/pkg/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	docs "github.com/fajarhidayad/waow-article/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			WAOW Article API
//	@version		1.0
//	@description	An Article API for WAOW homework.

//	@contact.name	API Support
//	@contact.email	fajarsuryahidayad@gmail.com

//	@host		localhost:8000
//	@BasePath	/api

//	@securityDefinitions.basic	JWT Auth

// @securityDefinitions.apikey	JWT Bearer Auth
// @in							header
// @name						Authorization
// @description				Bearer Token
func main() {
	r := gin.Default()
	db := database.ConnectDB()

	docs.SwaggerInfo.BasePath = "/api"

	r.GET("/health", middleware.HasAccessToken(), func(ctx *gin.Context) {
		if db != nil {
			fmt.Println("Database Connected")
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "100% running",
		})
	})

	api := r.Group("/api")
	{
		routes.AuthRoutes(api, db)
		routes.AdminRoutes(api, db)
	}
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	err := r.Run(PORT)
	if err != nil {
		panic(err)
	}
}
