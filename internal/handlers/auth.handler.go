package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fajarhidayad/waow-article/internal/dtos"
	"github.com/fajarhidayad/waow-article/internal/services"
	"github.com/fajarhidayad/waow-article/pkg/auth"
	"github.com/fajarhidayad/waow-article/pkg/common"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Register(*gin.Context)
	Login(*gin.Context)
	Refresh(*gin.Context)
}

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user and returns user ID
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dtos.RegisterRequest true "User Registration"
// @Success 201 {object} common.Response
// @Router /auth/register [post]
func (handler *authHandler) Register(ctx *gin.Context) {
	var user dtos.RegisterRequest
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

// Login godoc
// @Summary Sign in user
// @Description Login a user, returns accessToken and RefreshToken
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dtos.LoginRequest true "Login User"
// @Success 201 {object} auth.TokenResponse
// @Router /auth/login [post]
func (handler *authHandler) Login(ctx *gin.Context) {
	var user dtos.LoginRequest
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

	setAuthCookie(ctx, response)

	ctx.JSON(http.StatusOK, response)
}

func (handler *authHandler) Refresh(ctx *gin.Context) {
	sub := ctx.GetString("sub")
	user, err := handler.authService.GetUser(sub)
	if err != nil {
		ctx.JSON(http.StatusNotFound, common.ErrorResponse{
			Error: err.Error(),
		})
	}

	response, err := auth.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	setAuthCookie(ctx, response)

	ctx.JSON(http.StatusOK, common.Response{
		Message: "Success",
		Data:    "Token refreshed",
	})
}

func setAuthCookie(ctx *gin.Context, response *auth.TokenResponse) {
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
}
