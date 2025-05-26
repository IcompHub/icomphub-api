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
)
