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
