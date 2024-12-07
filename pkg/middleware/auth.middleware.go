package middleware

import (
	"fmt"
	"github.com/fajarhidayad/waow-article/internal/models"
	"net/http"
	"os"
	"strings"

	"github.com/fajarhidayad/waow-article/pkg/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func HasAccessToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//tokenString, err := ctx.Cookie("Authorization")
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse{
				Error: "Unauthorized",
			})
			return
		}

		if !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, common.ErrorResponse{
				Error: "Unauthorized",
			})
		}

		tokenString = strings.Split(tokenString, " ")[1]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			secret := os.Getenv("SECRET_ACCESS_TOKEN")
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &common.ErrorResponse{
				Error: "invalid or expired token",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			ctx.Set("username", claims["username"])
			ctx.Set("role", claims["role"])
		}

		ctx.Next()
	}
}

func HasRefreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("X-REFRESH-TOKEN")
		if err != nil || tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse{
				Error: "Unauthorized",
			})
			return
		}

		tokenString = strings.Split(tokenString, " ")[1]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			secret := os.Getenv("SECRET_REFRESH_TOKEN")
			return []byte(secret), nil
		})

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &common.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx.Set("sub", claims["sub"])
		}

		ctx.Next()
	}
}

func HasRoleAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exist := ctx.Get("role")
		if !exist {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse{
				Error: "Unauthorized",
			})
		}

		if role != models.ROLE_ADMIN {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse{
				Error: "Unauthorized",
			})
		}

		ctx.Next()
	}
}

func HasRoleUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exist := ctx.Get("role")
		if !exist {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse{
				Error: "Unauthorized",
			})
		}

		if role != models.ROLE_USER {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse{
				Error: "Unauthorized",
			})
		}

		ctx.Next()
	}
}
