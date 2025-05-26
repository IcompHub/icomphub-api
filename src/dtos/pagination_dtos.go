package dtos

type PaginationDTO[T any] struct {
	TotalItems uint64 `json:"total_items" example:"20"`
	TotalPages uint64 `json:"total_pages" example:"2"`
	PageNumber uint64 `json:"page_number" example:"1"`
	PageSize   uint64 `json:"page_size" example:"10"`
	Items      []T    `json:"items"`
}

type PaginationRequestDTO struct {
	PageNumber uint64 `form:"pageNumber" json:"pageNumber" default:"1" binding:"required,min=1"`
	PageSize   uint64 `form:"pageSize" json:"pageSize" default:"10" binding:"required,min=1"`
}
