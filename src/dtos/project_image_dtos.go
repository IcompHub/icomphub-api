package dtos

import "mime/multipart"

type ProjectImageDTO struct {
	Id      uint64 `json:"id"`
	ImageID string `json:"image_id"`
}

type ProjectImageRequestDTO struct {
	PaginationRequestDTO
}

type ProjectImageCreateRequestDTO struct {
	ProjectID uint64                `form:"project_id" json:"project_id" binding:"required"`
	Image     *multipart.FileHeader `form:"image" json:"image" binding:"required"`
}

type ProjectImageUpdateRequestDTO struct {
	Image *multipart.FileHeader `form:"image" json:"image" binding:"required"`
}
