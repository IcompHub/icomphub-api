package controllers

import (
	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/handlers"
	"icomphub-api/services"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
	service services.MemberService
}

func NewMemberController(service services.MemberService) *MemberController {
	return &MemberController{service}
}

// @Summary      Find a member
// @Tags         members
// @Produce      json
// @Param        id   path  uint64  true "Member ID"
// @Success      200  {object}  dtos.Response[dtos.MemberDTO]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /members/{id} [get]
func (controller *MemberController) Find(ctx *gin.Context) {
	id, err := handlers.ValidateId(ctx)
	if err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	member, code, err := controller.service.GetByID(id)
	if code == codes.ErrorFindingMember {
		handlers.NotFound(ctx, code, err)
		return
	}
	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, code, "Member found with success", member)
}

// @Summary      Create a member
// @Tags         members
// @Accept       json
// @Produce      json
// @Param        MemberCreateRequestDTO  body dtos.MemberCreateRequestDTO  true "MemberCreateRequestDTO"
// @Success      200  {object}  dtos.Response[dtos.MemberDTO]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /members [post]
func (controller *MemberController) Create(ctx *gin.Context) {
	var req dtos.MemberCreateRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	member, code, err := controller.service.Create(&req)
	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, code, "Created member with success", member)
}

// @Summary      Update a member
// @Tags         members
// @Accept       json
// @Produce      json
// @Param        id   path  uint64  true "Member ID"
// @Param        MemberUpdateRequestDTO  body dtos.MemberUpdateRequestDTO  true "MemberUpdateRequestDTO"
// @Success      200  {object}  dtos.Response[dtos.MemberDTO]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /members/{id} [put]
func (controller *MemberController) Update(ctx *gin.Context) {
	id, err := handlers.ValidateId(ctx)
	if err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	var req dtos.MemberUpdateRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	member, code, err := controller.service.Update(id, &req)
	if code == codes.ErrorFindingMember {
		handlers.NotFound(ctx, code, err)
		return
	}
	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, code, "Updated member with success", member)
}

// @Summary      Delete a member
// @Tags         members
// @Produce      json
// @Param        id   path  uint64  true "Member ID"
// @Success      200  {object}  dtos.Response[any]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /members/{id} [delete]
func (controller *MemberController) Delete(ctx *gin.Context) {
	id, err := handlers.ValidateId(ctx)
	if err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	code, err := controller.service.Delete(id)
	if code == codes.ErrorFindingMember {
		handlers.NotFound(ctx, code, err)
		return
	}
	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, code, "Deleted member with success", nil)
}
