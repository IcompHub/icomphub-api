package dtos

import "icomphub-api/enums"

type UserDTO struct {
	Id            uint64               `json:"id"`
	Slug          string               `json:"slug"`
	Nickname      string               `json:"nickname"`
	PersonalEmail string               `json:"personal_email"`
	Role          enums.SystemRoleEnum `json:"role"`
}

type UserRequestDTO struct {
	PaginationRequestDTO
	Search string `form:"search" json:"search" example:"Ana"`
}
