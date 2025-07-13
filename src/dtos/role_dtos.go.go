package dtos

type RoleDTO struct {
	Id     uint64 `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type RoleRequestDTO struct {
	PaginationRequestDTO
	Search string `form:"search" json:"search" example:"My Role"`
}

type RoleCreateRequestDTO struct {
	Slug string `form:"slug" json:"slug" binding:"required,min=2"`
	Name string `form:"name" json:"name" binding:"required,min=2"`
}

type RoleUpdateRequestDTO struct {
	Slug *string `form:"slug" json:"slug"`
	Name *string `form:"name" json:"name"`
}
