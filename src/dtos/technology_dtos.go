package dtos

type TechnologyDTO struct {
	Id   uint64 `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type TechnologyRequestDTO struct {
	PaginationRequestDTO
	Search string `form:"search" json:"search" example:"Nodejs"`
}

type TechnologyCreateRequestDTO struct {
	Slug string `form:"slug" json:"slug" binding:"required,min=2"`
	Name string `form:"name" json:"name" binding:"required,min=2"`
}

type TechnologyUpdateRequestDTO struct {
	Slug string `form:"slug" json:"slug"`
	Name string `form:"name" json:"name"`
}
