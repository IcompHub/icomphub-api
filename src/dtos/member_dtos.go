package dtos

type MemberDTO struct {
	ID        uint64  `json:"id" `
	Nickname  string  `json:"nickname"`
	Status    string  `json:"status" `
	ProjectID uint64  `json:"project_id"`
	UserId    *uint64 `json:"user_id" `

	Project ProjectShortDTO `json:"project"`
	User    *UserDTO        `json:"user"`
	Roles   []RoleDTO       `json:"roles"`
}

type MemberShortDTO struct {
	ID       uint64    `json:"id" `
	Nickname string    `json:"nickname"`
	Status   string    `json:"status" `
	Roles    []RoleDTO `json:"roles"`
}

type MemberRequestDTO struct {
	PaginationRequestDTO
	Search string `form:"search" json:"search" example:"My Member"`
}

type MemberCreateRequestDTO struct {
	Nickname  string   `form:"nickname" json:"nickname" binding:"required,min=2"`
	ProjectID uint64   `form:"project_id" json:"project_id" binding:"required"`
	UserId    *uint64  `form:"user_id" json:"user_id"`
	RoleIDs   []uint64 `form:"role_ids" json:"role_ids" binding:"required"`
}

type MemberUpdateRequestDTO struct {
	Nickname  *string                `form:"nickname" json:"nickname"`
	ProjectID *uint64                `form:"project_id" json:"project_id"`
	UserId    *uint64                `form:"user_id" json:"user_id"`
	Roles     []RoleCreateRequestDTO `form:"roles" json:"roles" binding:"required"`
}
