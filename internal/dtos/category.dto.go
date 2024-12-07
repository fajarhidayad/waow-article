package dtos

type CreateCategoryDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateCategoryDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
