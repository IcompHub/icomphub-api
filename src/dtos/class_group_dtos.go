package dtos

type ClassGroupDTO struct {
	Id     uint64 `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type ClassGroupRequestDTO struct {
	PaginationRequestDTO
	Search string `form:"search" json:"search" example:"Math 101"`
}

type ClassGroupCreateRequestDTO struct {
	Slug string `form:"slug" json:"slug" binding:"required,min=2"`
	Name string `form:"name" json:"name" binding:"required,min=2"`
}

type ClassGroupUpdateRequestDTO struct {
	Slug *string `form:"slug" json:"slug"`
	Name *string `form:"name" json:"name"`
}
