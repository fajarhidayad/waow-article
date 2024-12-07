package handlers

import (
	"github.com/fajarhidayad/waow-article/internal/dtos"
	"github.com/fajarhidayad/waow-article/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleHandler interface {
	Create(*gin.Context)
	FindAll(*gin.Context)
	FindById(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type articleHandler struct {
	articleService services.ArticleService
}

func NewArticleHandler(a services.ArticleService) ArticleHandler {
	return &articleHandler{
		articleService: a,
	}
}

func (h *articleHandler) Create(c *gin.Context) {
	var article dtos.CreateArticleDto
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := c.MustGet("username").(string)

	res, err := h.articleService.Create(username, &article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, res)
}

func (h *articleHandler) FindAll(c *gin.Context) {
	res, err := h.articleService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *articleHandler) FindById(c *gin.Context) {
	id := c.Param("id")
	res, err := h.articleService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, res)
}

func (h *articleHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var article dtos.UpdateArticleDto
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := c.MustGet("username").(string)

	res, err := h.articleService.Update(id, username, &article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *articleHandler) Delete(c *gin.Context) {
	articleId := c.Param("id")
	res, err := h.articleService.Delete(articleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
