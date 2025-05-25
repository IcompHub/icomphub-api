package dto

type TechnologyDTO struct {
	Slug string `json:"slug" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type TechnologyGetDTO struct {
	ID   uint64 `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}
