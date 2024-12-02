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
)

func main() {
	r := gin.Default()
	db := database.ConnectDB()

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
	}

	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	err := r.Run(PORT)
	if err != nil {
		panic(err)
	}
}
