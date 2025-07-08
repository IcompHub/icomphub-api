package controllers

import (
	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/handlers"
	"icomphub-api/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service}
}

// @Summary      List all users
// @Security BearerAuth
// @Tags         users
// @Produce      json
// @Param        UserRequestDTO  query dtos.UserRequestDTO  true  "UserRequestDTO"
// @Success      200  {object}  dtos.Response[dtos.PaginationDTO[dtos.UserDTO]]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /users [get]
func (controller *UserController) GetAll(context *gin.Context) {
	var req dtos.UserRequestDTO

	err := context.ShouldBindQuery(&req)
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
		return
	}

	var users []dtos.UserDTO

	users, code, err := controller.service.GetAll(&req)
	if err != nil {
		handlers.InternalServerError(context, code, err)
		return
	}

	count, code, err := controller.service.CountAll(&req)
	if err != nil {
		handlers.InternalServerError(context, code, err)
		return
	}

	handlers.Ok(context, code, "Got all users with success", handlers.Paginate(users, count, req.PageNumber, req.PageSize))
}

// @Summary      Find an user
// @Tags         users
// @Produce      json
// @Param        id   path  uint64  true "User ID"
// @Success      200  {object}  dtos.Response[dtos.UserDTO]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /users/{id} [get]
func (controller *UserController) Find(context *gin.Context) {
	id, err := handlers.ValidateId(context)
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
		return
	}

	var user *dtos.UserDTO

	user, code, err := controller.service.Find(id)

	if code == codes.ErrorFindingUser {
		handlers.NotFound(context, code, err)
		return
	}

	if err != nil {
		handlers.InternalServerError(context, code, err)
		return
	}

	handlers.Ok(context, codes.FindUser, "user found with success", user)
}

// @Summary      Create an user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        UserCreateRequestDTO  body dtos.UserCreateRequestDTO  true "UserCreateRequestDTO"
// @Success      200  {object}  dtos.Response[dtos.UserDTO]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /users [post]
func (controller *UserController) Create(context *gin.Context) {
	var req dtos.UserCreateRequestDTO

	err := context.ShouldBindBodyWithJSON(&req)
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
		return
	}

	var user *dtos.UserDTO

	user, code, err := controller.service.Create(&req)
	if err != nil {
		handlers.InternalServerError(context, code, err)
		return
	}

	handlers.Ok(context, codes.CreateUser, "created user with success", user)
}

// @Summary      Update an user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path  uint64  true "User ID"
// @Param        UserUpdateRequestDTO  body dtos.UserUpdateRequestDTO  true "UserUpdateRequestDTO"
// @Success      200  {object}  dtos.Response[dtos.UserDTO]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /users/{id} [put]
func (controller *UserController) Update(context *gin.Context) {
	id, err := handlers.ValidateId(context)
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
		return
	}

	var req dtos.UserUpdateRequestDTO

	err = context.ShouldBindBodyWithJSON(&req)
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
		return
	}

	var user *dtos.UserDTO

	user, code, err := controller.service.Update(id, &req)

	if code == codes.ErrorFindingUser {
		handlers.NotFound(context, code, err)
		return
	}

	if err != nil {
		handlers.InternalServerError(context, code, err)
		return
	}

	handlers.Ok(context, codes.UpdateUser, "user upadated with success", user)
}

// @Summary      Delete an user
// @Tags         users
// @Produce      json
// @Param        id   path  uint64  true "User ID"
// @Success      200  {object}  dtos.Response[any]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /users/{id} [delete]
func (controller *UserController) Delete(context *gin.Context) {
	id, err := handlers.ValidateId(context)
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
		return
	}

	code, err := controller.service.Delete(id)

	if code == codes.ErrorFindingUser {
		handlers.NotFound(context, code, err)
		return
	}

	if err != nil {
		handlers.InternalServerError(context, code, err)
		return
	}

	handlers.Ok(context, codes.DeleteUser, "user deleted with success", nil)
}
