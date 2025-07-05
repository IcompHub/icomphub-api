package controllers

import (
	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/handlers"
	"icomphub-api/services"

	"github.com/gin-gonic/gin"
)

type ClassGroupController struct {
	service services.ClassGroupService
}

func NewClassGroupController(service services.ClassGroupService) *ClassGroupController {
	return &ClassGroupController{service}
}

// @Summary      List all class groups
// @Tags         class_groups
// @Produce      json
// @Param        ClassGroupRequestDTO  query dtos.ClassGroupRequestDTO  true  "ClassGroupRequestDTO"
// @Router       /class_groups [get]
func (controller *ClassGroupController) GetAll(ctx *gin.Context) {
	var req dtos.ClassGroupRequestDTO

	if err := ctx.ShouldBindQuery(&req); err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	classGroups, code, err := controller.service.GetAll(&req)
	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	count, code, err := controller.service.CountAll(&req)
	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, code, "Got all class groups with success", handlers.Paginate(classGroups, count, req.PageNumber, req.PageSize))
}

// @Summary      Find a class group
// @Tags         class_groups
// @Produce      json
// @Param        id   path  uint64  true "ClassGroup ID"
// @Router       /class_groups/{id} [get]
func (controller *ClassGroupController) Find(ctx *gin.Context) {
	id, err := handlers.ValidateId(ctx)
	if err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	classGroup, code, err := controller.service.Find(id)

	if code == codes.ErrorFindingClassGroup {
		handlers.NotFound(ctx, code, err)
		return
	}

	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, codes.FindClassGroup, "class group found with success", classGroup)
}

// @Summary      Create a class group
// @Tags         class_groups
// @Accept       json
// @Produce      json
// @Param        ClassGroupCreateRequestDTO  body dtos.ClassGroupCreateRequestDTO  true "ClassGroupCreateRequestDTO"
// @Router       /class_groups [post]
func (controller *ClassGroupController) Create(ctx *gin.Context) {
	var req dtos.ClassGroupCreateRequestDTO

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	classGroup, code, err := controller.service.Create(&req)
	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, codes.CreateClassGroup, "created class group with success", classGroup)
}

// @Summary      Update a class group
// @Tags         class_groups
// @Accept       json
// @Produce      json
// @Param        id   path  uint64  true "ClassGroup ID"
// @Param        ClassGroupUpdateRequestDTO  body dtos.ClassGroupUpdateRequestDTO  true "ClassGroupUpdateRequestDTO"
// @Router       /class_groups/{id} [put]
func (controller *ClassGroupController) Update(ctx *gin.Context) {
	id, err := handlers.ValidateId(ctx)
	if err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	var req dtos.ClassGroupUpdateRequestDTO

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	classGroup, code, err := controller.service.Update(id, &req)

	if code == codes.ErrorFindingClassGroup {
		handlers.NotFound(ctx, code, err)
		return
	}

	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, codes.UpdateClassGroup, "class group updated with success", classGroup)
}

// @Summary      Delete a class group
// @Tags         class_groups
// @Produce      json
// @Param        id   path  uint64  true "ClassGroup ID"
// @Router       /class_groups/{id} [delete]
func (controller *ClassGroupController) Delete(ctx *gin.Context) {
	id, err := handlers.ValidateId(ctx)
	if err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	code, err := controller.service.Delete(id)

	if code == codes.ErrorFindingClassGroup {
		handlers.NotFound(ctx, code, err)
		return
	}

	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, codes.DeleteClassGroup, "class group deleted with success", nil)
}
