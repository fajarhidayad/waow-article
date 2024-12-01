package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "100% running",
		})
	})

	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))
	err := r.Run(PORT)
	if err != nil {
		panic(err)
	}
}
