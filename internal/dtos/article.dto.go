package dtos

type CreateArticleDto struct {
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	CategoryID string `json:"category_id" binding:"required"`
}
type UpdateArticleDto struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryID string `json:"category_id"`
}
