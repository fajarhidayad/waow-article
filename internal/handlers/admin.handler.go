package handlers

import (
	"github.com/fajarhidayad/waow-article/internal/dtos"
	"github.com/fajarhidayad/waow-article/internal/services"
	"github.com/fajarhidayad/waow-article/pkg/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminHandler interface {
	CreateUser(ctx *gin.Context)
	FindAllUsers(ctx *gin.Context)
	FindUserByID(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type adminHandlerImpl struct {
	userService services.AdminService
}

func NewAdminHandler(userService services.AdminService) AdminHandler {
	return &adminHandlerImpl{
		userService: userService,
	}
}

func (h *adminHandlerImpl) CreateUser(ctx *gin.Context) {
	var user dtos.CreateUserDTO
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	response, errResponse := h.userService.CreateUser(&user)
	if errResponse != nil {
		ctx.JSON(http.StatusConflict, errResponse)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *adminHandlerImpl) FindAllUsers(ctx *gin.Context) {
	res, errResponse := h.userService.FindUsers()
	if errResponse != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *adminHandlerImpl) FindUserByID(ctx *gin.Context) {
	userID := ctx.Param("id")

	response, errResponse := h.userService.FindUserByID(userID)
	if errResponse != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *adminHandlerImpl) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	var user dtos.UpdateUserDTO
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse{
			Error: err.Error(),
		})
	}

	res, err := h.userService.UpdateUser(userID, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *adminHandlerImpl) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	res, errResponse := h.userService.DeleteUser(userID)
	if errResponse != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
