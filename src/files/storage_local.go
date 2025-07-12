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
	defer func() {
		if cerr := src.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "error closing src file: %v\n", cerr)
		}
	}()

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
	defer func() {
		if cerr := dst.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "error closing dst file: %v\n", cerr)
		}
	}()

	_, err = io.Copy(dst, src)
	if err != nil {
		return "", codes.FileCouldNotSave, err
	}

	return filename, codes.FileSaved, nil // relative path for DB
}

func (s *LocalFileUploadService) DeleteFile(fileName string, config UploadConfig) (codes.Code, error) {
	// Prevent accidental deletes outside allowed scope
	if fileName == "" {
		return codes.InvalidParams, fmt.Errorf("file name cannot be empty")
	}

	// Construct the full absolute path
	absPath := filepath.Join(s.BasePath, config.TargetFolder, fileName)

	// Check if file exists before attempting to delete
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return codes.FileNotFound, fmt.Errorf("file not found: %s", absPath)
	}

	// Attempt to delete the file
	err := os.Remove(absPath)
	if err != nil {
		return codes.FileCouldNotDelete, fmt.Errorf("failed to delete file: %w", err)
	}

	return codes.FileDeleted, nil
}

func (s *LocalFileUploadService) GetFile(fileName string, config UploadConfig) (string, codes.Code, error) {
	if fileName == "" {
		return "", codes.InvalidParams, fmt.Errorf("file name is required")
	}

	relPath := filepath.Join(config.TargetFolder, fileName)
	absPath := filepath.Join(s.BasePath, relPath)

	// Check if file exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return "", codes.FileNotFound, fmt.Errorf("file not found: %s", fileName)
	}

	return absPath, codes.FileFound, nil
}
