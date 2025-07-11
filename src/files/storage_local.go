package files

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"icomphub-api/codes"
)

type LocalFileUploadService struct {
	BasePath string // e.g., "./static/uploads"
}

func NewLocalFileUploadService(basePath string) *LocalFileUploadService {
	return &LocalFileUploadService{BasePath: basePath}
}

func (s *LocalFileUploadService) SaveFile(fileHeader *multipart.FileHeader, config UploadConfig) (string, codes.Code, error) {
	// Validate file type
	if len(config.AllowedTypes) > 0 && !contains(config.AllowedTypes, fileHeader.Header.Get("Content-Type")) {
		return "", codes.FileInvalidType, fmt.Errorf("invalid file type: %s", fileHeader.Header.Get("Content-Type"))
	}

	// Open source file
	src, err := fileHeader.Open()
	if err != nil {
		return "", codes.FileCouldNotOpen, err
	}
	defer src.Close()

	// Get extension from original filename (optional)
	ext := filepath.Ext(fileHeader.Filename)

	// Generate secure random string
	randomBytes := make([]byte, 16)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return "", codes.UnknowError, fmt.Errorf("failed to generate random filename: %w", err)
	}
	randomString := hex.EncodeToString(randomBytes)

	// Final filename: <timestamp>_<random>.ext
	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), randomString, ext)

	// Build target path
	relPath := filepath.Join(config.TargetFolder, filename)
	absPath := filepath.Join(s.BasePath, relPath)

	// Ensure dir exists
	err = os.MkdirAll(filepath.Dir(absPath), os.ModePerm)
	if err != nil {
		return "", codes.FileCouldNotFindTargetDir, err
	}

	// Save
	dst, err := os.Create(absPath)
	if err != nil {
		return "", codes.FileCouldNotSave, err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return "", codes.FileCouldNotSave, err
	}

	return filename, codes.FileSaved, nil // relative path for DB
}
