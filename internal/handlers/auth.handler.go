package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fajarhidayad/waow-article/internal/dtos"
	"github.com/fajarhidayad/waow-article/internal/services"
	"github.com/fajarhidayad/waow-article/pkg/common"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Register(*gin.Context)
	Login(*gin.Context)
}

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (handler *authHandler) Register(ctx *gin.Context) {
	var user dtos.RegisterDto
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, &common.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	response, err := handler.authService.Register(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (handler *authHandler) Login(ctx *gin.Context) {
	var user dtos.LoginDto
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, &common.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	response, err := handler.authService.Login(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &common.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	ctx.SetCookie("Authorization",
		fmt.Sprintf("Bearer %s", response.AccessToken),
		60*15, // 15 minutes
		"/",
		os.Getenv("DOMAIN"),
		false,
		true,
	)
	ctx.SetCookie("X-REFRESH-TOKEN",
		response.RefreshToken,
		60*60*24*7, // 7 Days
		"/",
		os.Getenv("DOMAIN"),
		false,
		true,
	)

	ctx.JSON(http.StatusOK, response)
}
