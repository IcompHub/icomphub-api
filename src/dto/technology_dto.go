package dto

type TechnologyDTO struct {
	Slug string `json:"slug" binding: "required"`
	Name string `json:"name" binding: "required"`
}
