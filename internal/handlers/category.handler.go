package handlers

import (
	"github.com/fajarhidayad/waow-article/internal/dtos"
	"github.com/fajarhidayad/waow-article/internal/services"
	"github.com/fajarhidayad/waow-article/pkg/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryHandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

type categoryHandlerImpl struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService services.CategoryService) CategoryHandler {
	return &categoryHandlerImpl{
		categoryService: categoryService,
	}
}

func (h *categoryHandlerImpl) Create(ctx *gin.Context) {
	var category dtos.CreateCategoryDto
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, &common.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	res, err := h.categoryService.CreateCategory(&category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusCreated, res)
}

func (h *categoryHandlerImpl) Update(ctx *gin.Context) {
	catId := ctx.Param("id")

	var category dtos.UpdateCategoryDto
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, &common.ErrorResponse{
			Error: err.Error(),
		})
	}

	res, err := h.categoryService.UpdateCategory(catId, &category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *categoryHandlerImpl) Delete(ctx *gin.Context) {
	catId := ctx.Param("id")

	res, err := h.categoryService.DeleteCategory(catId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *categoryHandlerImpl) FindAll(ctx *gin.Context) {
	res, err := h.categoryService.FindAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *categoryHandlerImpl) FindByID(ctx *gin.Context) {
	catId := ctx.Param("id")

	res, err := h.categoryService.FindCategoryByID(catId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
