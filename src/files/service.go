package files

import (
	"mime/multipart"

	"icomphub-api/codes"
)

type UploadConfig struct {
	TargetFolder string   // e.g., "users", "projects"
	AllowedTypes []string // e.g., ["image/jpeg", "image/png"]
	MaxSizeMB    int64    // e.g., 5
}

type FileUploadService interface {
	SaveFile(fileHeader *multipart.FileHeader, config UploadConfig) (string, codes.Code, error)
}
