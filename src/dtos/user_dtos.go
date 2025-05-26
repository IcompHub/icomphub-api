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

type UserCreateRequestDTO struct {
	Slug               string  `form:"slug" json:"slug" binding:"required,min=2"`
	Nickname           string  `form:"nickname" json:"nickname" binding:"required,min=2"`
	FullName           string  `form:"full_name" json:"full_name" binding:"required,min=2"`
	PersonalEmail      string  `form:"personal_email" json:"personal_email" binding:"required,email,min=5"`
	InstitutionalEmail *string `form:"institutional_email" json:"institutional_email"`
	Registration       *string `form:"registration" json:"registration"`
}

type UserUpdateRequestDTO struct {
	Slug               string `form:"slug" json:"slug"`
	Nickname           string `form:"nickname" json:"nickname"`
	FullName           string `form:"full_name" json:"full_name"`
	PersonalEmail      string `form:"personal_email" json:"personal_email"`
	InstitutionalEmail string `form:"institutional_email" json:"institutional_email"`
	Registration       string `form:"registration" json:"registration"`
}
