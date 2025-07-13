package controllers

import (
	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/handlers"
	"icomphub-api/services"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	service services.RoleService
}

func NewRoleController(service services.RoleService) *RoleController {
	return &RoleController{service}
}

// @Summary      Find a role
// @Tags         roles
// @Produce      json
// @Param        id   path  uint64  true "Role ID"
// @Success      200  {object}  dtos.Response[dtos.RoleDTO]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /roles/{id} [get]
func (controller *RoleController) Find(ctx *gin.Context) {
	id, err := handlers.ValidateId(ctx)
	if err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	role, code, err := controller.service.GetByID(id)

	if code == codes.ErrorFindingRole {
		handlers.NotFound(ctx, code, err)
		return
	}

	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, codes.FindRole, "role found with success", role)
}

// @Summary      Create a role
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        RoleCreateRequestDTO  body dtos.RoleCreateRequestDTO  true "RoleCreateRequestDTO"
// @Success      200  {object}  dtos.Response[dtos.RoleDTO]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /roles [post]
func (controller *RoleController) Create(ctx *gin.Context) {
	var req dtos.RoleCreateRequestDTO

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	role, code, err := controller.service.Create(&req)
	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, codes.CreateRole, "created role with success", role)
}

// @Summary      Update a role
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        id   path  uint64  true "Role ID"
// @Param        RoleUpdateRequestDTO  body dtos.RoleUpdateRequestDTO  true "RoleUpdateRequestDTO"
// @Success      200  {object}  dtos.Response[dtos.RoleDTO]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /roles/{id} [put]
func (controller *RoleController) Update(ctx *gin.Context) {
	id, err := handlers.ValidateId(ctx)
	if err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	var req dtos.RoleUpdateRequestDTO

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	role, code, err := controller.service.Update(id, &req)

	if code == codes.ErrorFindingRole {
		handlers.NotFound(ctx, code, err)
		return
	}

	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, codes.UpdateRole, "role updated with success", role)
}

// @Summary      Delete a role
// @Tags         roles
// @Produce      json
// @Param        id   path  uint64  true "Role ID"
// @Success      200  {object}  dtos.Response[any]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /roles/{id} [delete]
func (controller *RoleController) Delete(ctx *gin.Context) {
	id, err := handlers.ValidateId(ctx)
	if err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	code, err := controller.service.Delete(id)

	if code == codes.ErrorFindingRole {
		handlers.NotFound(ctx, code, err)
		return
	}

	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, codes.DeleteRole, "role deleted with success", nil)
}
