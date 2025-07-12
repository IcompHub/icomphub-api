package controllers

import (
	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/handlers"
	"icomphub-api/services"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	service services.ProjectService
}

func NewProjectController(service services.ProjectService) *ProjectController {
	return &ProjectController{service}
}

// @Summary      List all projects
// @Tags         projects
// @Produce      json
// @Param        ProjectRequestDTO  query dtos.ProjectRequestDTO  true  "ProjectRequestDTO"
// @Success      200  {object}  dtos.Response[dtos.PaginationDTO[dtos.ProjectDTO]]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /projects [get]
func (controller *ProjectController) GetAll(ctx *gin.Context) {
	var req dtos.ProjectRequestDTO

	if err := ctx.ShouldBindQuery(&req); err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	projects, code, err := controller.service.GetAll(&req)
	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	count, code, err := controller.service.CountAll(&req)
	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, code, "Got all projects with success", handlers.Paginate(projects, count, req.PageNumber, req.PageSize))
}

// @Summary      Find a project
// @Tags         projects
// @Produce      json
// @Param        id   path  uint64  true "Project ID"
// @Success      200  {object}  dtos.Response[dtos.ProjectDTO]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /projects/{id} [get]
func (controller *ProjectController) Find(ctx *gin.Context) {
	id, err := handlers.ValidateId(ctx)
	if err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	project, code, err := controller.service.Find(id)

	if code == codes.ErrorFindingProject {
		handlers.NotFound(ctx, code, err)
		return
	}

	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, codes.FindProject, "project found with success", project)
}

// @Summary      Create a project
// @Tags         projects
// @Accept       json
// @Produce      json
// @Param        ProjectCreateRequestDTO  body dtos.ProjectCreateRequestDTO  true "ProjectCreateRequestDTO"
// @Success      200  {object}  dtos.Response[dtos.ProjectDTO]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /projects [post]
func (controller *ProjectController) Create(ctx *gin.Context) {
	var req dtos.ProjectCreateRequestDTO

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	project, code, err := controller.service.Create(&req)
	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, codes.CreateProject, "created project with success", project)
}

// @Summary      Update a project
// @Tags         projects
// @Accept       json
// @Produce      json
// @Param        id   path  uint64  true "Project ID"
// @Param        ProjectUpdateRequestDTO  body dtos.ProjectUpdateRequestDTO  true "ProjectUpdateRequestDTO"
// @Success      200  {object}  dtos.Response[dtos.ProjectDTO]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /projects/{id} [put]
func (controller *ProjectController) Update(ctx *gin.Context) {
	id, err := handlers.ValidateId(ctx)
	if err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	var req dtos.ProjectUpdateRequestDTO

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	project, code, err := controller.service.Update(id, &req)

	if code == codes.ErrorFindingProject {
		handlers.NotFound(ctx, code, err)
		return
	}

	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, codes.UpdateProject, "project updated with success", project)
}

// @Summary      Delete a project
// @Tags         projects
// @Produce      json
// @Param        id   path  uint64  true "Project ID"
// @Success      200  {object}  dtos.Response[any]
// @Failure      500  {object}  dtos.Response[any]
// @Router       /projects/{id} [delete]
func (controller *ProjectController) Delete(ctx *gin.Context) {
	id, err := handlers.ValidateId(ctx)
	if err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	code, err := controller.service.Delete(id)

	if code == codes.ErrorFindingProject {
		handlers.NotFound(ctx, code, err)
		return
	}

	if err != nil {
		handlers.InternalServerError(ctx, code, err)
		return
	}

	handlers.Ok(ctx, codes.DeleteProject, "project deleted with success", nil)
}

// @Summary      Update  project thumbnail
// @Tags         projects
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param        id   path  uint64  true "Project ID"
// @Param image formData file false "Project thumbnail image"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router       /projects/thumbnail/{id} [post]
func (controller *ProjectController) UpdateThumbnail(context *gin.Context) {
	projectID, err := handlers.ValidateId(context)
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
		return
	}

	image, err := context.FormFile("image")
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
	}

	project, code, err := controller.service.UpdateThumbnail(projectID, image)
	if err != nil {
		handlers.BadRequest(context, code, err)
		return
	}

	handlers.Ok(context, code, "project thumbnail with success", project)
}

// @Summary      Delete project thumbnail
// @Tags         projects
// @Security BearerAuth
// @Param        id   path  uint64  true "Project ID"
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router       /projects/thumbnail/{id} [delete]
func (controller *ProjectController) DeleteThumbnail(context *gin.Context) {
	projectID, err := handlers.ValidateId(context)
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
		return
	}

	project, code, err := controller.service.DeleteThumbnail(projectID)
	if err != nil {
		handlers.BadRequest(context, code, err)
		return
	}

	handlers.Ok(context, code, "project thumbnail deleted with success", project)
}

// @Summary      Get project thumbnail by ID
// @Tags         projects
// @Security BearerAuth
// @Param        id   path  uint64  true "Project ID"
// @Produce json
// @Success      200 {file} file
// @Failure 400 {object} map[string]string
// @Router       /projects/thumbnail/{id} [get]
func (controller *ProjectController) GetThumbnail(context *gin.Context) {
	projectID, err := handlers.ValidateId(context)
	if err != nil {
		handlers.BadRequest(context, codes.InvalidParams, err)
		return
	}

	fullPath, code, err := controller.service.GetThumbnailFullPath(projectID)
	if err != nil {
		handlers.BadRequest(context, code, err)
		return
	}

	context.File(fullPath)
}
