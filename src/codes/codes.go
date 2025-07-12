package codes

type (
	Code string
)

const (
	UnknowError   Code = "unknow_error"
	InvalidParams Code = "invalid_params"

	GetAllUsers   Code = "get_all_users"
	CountAllUsers Code = "count_all_users"
	FindUser      Code = "find_user"
	UpdateUser    Code = "update_user"
	CreateUser    Code = "create_user"
	DeleteUser    Code = "delete_user"

	ErrorGettingAllUsers Code = "error_getting_users"
	ErrorCoutingAllUsers Code = "error_counting_users"
	ErrorFindingUser     Code = "error_finding_user"
	ErrorUpdatingUser    Code = "error_updating_user"
	ErrorCreatingUser    Code = "error_creating_user"
	ErrorDeletingUser    Code = "error_deleting_user"

	GetAllTechnologies   Code = "get_all_technologies"
	CountAllTechnologies Code = "count_all_technologies"
	FindTechnology       Code = "find_technology"
	UpdateTechnology     Code = "update_technology"
	CreateTechnology     Code = "create_technology"
	DeleteTechnology     Code = "delete_technology"

	ErrorGettingAllTechnologies Code = "error_getting_technologies"
	ErrorCoutingAllTechnologies Code = "error_counting_technologies"
	ErrorFindingTechnology      Code = "error_finding_technology"
	ErrorUpdatingTechnology     Code = "error_updating_technology"
	ErrorCreatingTechnology     Code = "error_creating_technology"
	ErrorDeletingTechnology     Code = "error_deleting_technology"

	GetAllProjects   Code = "get_all_projects"
	CountAllProjects Code = "count_all_projects"
	FindProject      Code = "find_project"
	CreateProject    Code = "create_project"
	UpdateProject    Code = "update_project"
	DeleteProject    Code = "delete_project"

	ErrorGettingAllProjects  Code = "error_getting_projects"
	ErrorCountingAllProjects Code = "error_counting_projects"
	ErrorFindingProject      Code = "error_finding_project"
	ErrorCreatingProject     Code = "error_creating_project"
	ErrorUpdatingProject     Code = "error_updating_project"
	ErrorDeletingProject     Code = "error_deleting_project"

	GetAllClassGroups   Code = "get_all_class_groups"
	CountAllClassGroups Code = "count_all_class_groups"
	FindClassGroup      Code = "find_class_group"
	CreateClassGroup    Code = "create_class_group"
	UpdateClassGroup    Code = "update_class_group"
	DeleteClassGroup    Code = "delete_class_group"

	ErrorGettingAllClassGroups  Code = "error_getting_class_groups"
	ErrorCountingAllClassGroups Code = "error_counting_class_groups"
	ErrorFindingClassGroup      Code = "error_finding_class_group"
	ErrorCreatingClassGroup     Code = "error_creating_class_group"
	ErrorUpdatingClassGroup     Code = "error_updating_class_group"
	ErrorDeletingClassGroup     Code = "error_deleting_class_group"

	LoginFailed  Code = "login_failed"
	LoginSuccess Code = "login_success"

	AuthMissingToken Code = "auth_missing_token"
	AuthExpiredToken Code = "auth_expired_token"
	AuthInvalidToken Code = "auth_invalid_token"
	AuthValidToken   Code = "auth_valid_token"

	RoleAccessDenied Code = "role_access_denied"
	RoleInsufficient Code = "role_insufficient"

	FileInvalidType           Code = "file_invalid_type"
	FileCouldNotOpen          Code = "file_could_not_open"
	FileCouldNotFindTargetDir Code = "file_could_not_find_target_dir"
	FileCouldNotSave          Code = "file_could_not_save"
	FileSaved                 Code = "file_saved"
	FileDeleted               Code = "file_deleted"
	FileFound                 Code = "file_found"
	FileNotFound              Code = "file_not_found"
	FileCouldNotDelete        Code = "file_could_not_delete"
)
