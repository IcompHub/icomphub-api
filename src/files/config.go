package files

const (
	DefaultUploadPath = "./static/uploads"
	MaxFileSizeMB     = 5
)

var AllowedImageTypes = []string{"image/jpeg", "image/png", "image/webp"}

// Upload targets — use these to organize subfolders
const (
	UserFolder       = "users"
	ProjectFolder    = "projects"
	TechnologyFolder = "technologies"
	ClassGroupFolder = "classgroups"
)
