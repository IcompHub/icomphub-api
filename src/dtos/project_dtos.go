package dtos

type MemberRoleDTO struct {
	Id   uint64 `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type MemberUserDTO struct {
	Id       uint64 `json:"id"`
	Slug     string `json:"slug"`
	Nickname string `json:"nickname"`
	FullName string `json:"full_name"`
}

type ProjectMemberDTO struct {
	Id       uint64        `json:"id"`
	Nickname string        `json:"nickname"`
	Status   string        `json:"status"`
	User     MemberUserDTO `json:"user"`
	Role     MemberRoleDTO `json:"role"`
}

type ProjectDetailDTO struct {
	Id           uint64             `json:"id"`
	Slug         string             `json:"slug"`
	Name         string             `json:"name"`
	Status       string             `json:"status"`
	Data         map[string]any     `json:"data"`
	ClassGroupId uint64             `json:"class_group_id"`
	Technologies []TechnologyDTO    `json:"technologies"`
	Members      []ProjectMemberDTO `json:"members"`
}

type ProjectDTO struct {
	Id           uint64          `json:"id"`
	Slug         string          `json:"slug"`
	Name         string          `json:"name"`
	Status       string          `json:"status"`
	Data         map[string]any  `json:"data"` // JSON structure with title, resume, etc.
	ClassGroupId uint64          `json:"class_group_id"`
	Technologies []TechnologyDTO `json:"technologies"`
}

type ProjectRequestDTO struct {
	PaginationRequestDTO
	Search string `form:"search" json:"search" example:"My Project"`
}

type ProjectCreateRequestDTO struct {
	Slug         string         `form:"slug" json:"slug" binding:"required,min=2"`
	Name         string         `form:"name" json:"name" binding:"required,min=2"`
	Data         map[string]any `form:"data" json:"data" binding:"required"`
	ClassGroupId uint64         `form:"class_group_id" json:"class_group_id" binding:"required"`
}

type ProjectUpdateRequestDTO struct {
	Slug         *string         `form:"slug" json:"slug"`
	Name         *string         `form:"name" json:"name"`
	Data         *map[string]any `form:"data" json:"data"`
	ClassGroupId *uint64         `form:"class_group_id" json:"class_group_id"`
}
