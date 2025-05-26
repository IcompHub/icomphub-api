package controllers

import (
	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/handlers"
	"icomphub-api/services"

	"github.com/gin-gonic/gin"
)

type TechnologyController struct {
	service services.TechnologyService
}

func NewTechnologyController(service services.TechnologyService) *TechnologyController {
	return &TechnologyController{service: service}
}

// @Summary      List all technologies
// @Tags         technologies
// @Produce      json
// @Param        TechnologyRequestDTO  query dtos.TechnologyRequestDTO  true "TechnologyRequestDTO"
// @Success      200  {object}  dtos.Response[dtos.PaginationDTO[dtos.TechnologyDTO]]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /technologies [get]
func (controller *TechnologyController) GetAll(context *gin.Context) {
	var req dtos.TechnologyRequestDTO

	err := context.ShouldBindQuery(&req)
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
		return
	}

	var technologies []dtos.TechnologyDTO

	technologies, code, err := controller.service.GetAll(&req)
	if err != nil {
		handlers.InternalServerError(context, code, err)
		return
	}

	count, code, err := controller.service.CountAll(&req)
	if err != nil {
		handlers.InternalServerError(context, code, err)
		return
	}

	handlers.Ok(context, codes.GetAllTechnologies, "got all technologies with success", handlers.Paginate(technologies, count, req.PageNumber, req.PageSize))
}

// @Summary      Find a technology
// @Tags         technologies
// @Produce      json
// @Param        id   path  uint64  true "Technology ID"
// @Success      200  {object}  dtos.Response[dtos.TechnologyDTO]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /technologies/{id} [get]
func (controller *TechnologyController) Find(context *gin.Context) {
	id, err := handlers.ValidateId(context)
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
		return
	}

	var technology *dtos.TechnologyDTO

	technology, code, err := controller.service.Find(id)

	if code == codes.ErrorFindingTechnology {
		handlers.NotFound(context, code, err)
		return
	}

	if err != nil {
		handlers.InternalServerError(context, code, err)
		return
	}

	handlers.Ok(context, codes.FindTechnology, "technology found with success", technology)
}

// @Summary      Create a technology
// @Tags         technologies
// @Accept       json
// @Produce      json
// @Param        TechnologyCreateRequestDTO  query dtos.TechnologyCreateRequestDTO  true "TechnologyCreateRequestDTO"
// @Success      200  {object}  dtos.Response[dtos.TechnologyDTO]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /technologies [post]
func (controller *TechnologyController) Create(context *gin.Context) {
	var req dtos.TechnologyCreateRequestDTO

	err := context.ShouldBindQuery(&req)
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
		return
	}

	var technology *dtos.TechnologyDTO

	technology, code, err := controller.service.Create(&req)
	if err != nil {
		handlers.InternalServerError(context, code, err)
		return
	}

	handlers.Ok(context, codes.CreateTechnology, "created technology with success", technology)
}

// @Summary      Update a technology
// @Tags         technologies
// @Accept       json
// @Produce      json
// @Param        id   path  uint64  true "Technology ID"
// @Param        TechnologyUpdateRequestDTO  query dtos.TechnologyUpdateRequestDTO  true "TechnologyUpdateRequestDTO"
// @Success      200  {object}  dtos.Response[dtos.TechnologyDTO]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /technologies/{id} [put]
func (controller *TechnologyController) Update(context *gin.Context) {
	id, err := handlers.ValidateId(context)
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
		return
	}

	var req dtos.TechnologyUpdateRequestDTO

	err = context.ShouldBindQuery(&req)
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
		return
	}

	var technology *dtos.TechnologyDTO

	technology, code, err := controller.service.Update(id, &req)

	if code == codes.ErrorFindingTechnology {
		handlers.NotFound(context, code, err)
		return
	}

	if err != nil {
		handlers.InternalServerError(context, code, err)
		return
	}

	handlers.Ok(context, codes.FindTechnology, "technology upadated with success", technology)
}

// @Summary      Delete a technology
// @Tags         technologies
// @Produce      json
// @Param        id   path  uint64  true "Technology ID"
// @Success      200  {object}  dtos.Response[any]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /technologies/{id} [delete]
func (controller *TechnologyController) Delete(context *gin.Context) {
	id, err := handlers.ValidateId(context)
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
		return
	}

	code, err := controller.service.Delete(id)

	if code == codes.ErrorFindingTechnology {
		handlers.NotFound(context, code, err)
		return
	}

	if err != nil {
		handlers.InternalServerError(context, code, err)
		return
	}

	handlers.Ok(context, codes.FindTechnology, "technology deleted with success", nil)
}
