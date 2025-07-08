package controllers

import (
	"icomphub-api/codes"
	"icomphub-api/dtos"
	"icomphub-api/handlers"
	"icomphub-api/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{service}
}

// @Summary Login user
// @Tags    auth
// @Accept  json
// @Produce json
// @Param   login body dtos.LoginRequestDTO true "Login credentials"
// @Success 200 {object} dtos.Response[string]
// @Failure 401 {object} dtos.Response[any]
// @Router  /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var req dtos.LoginRequestDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handlers.BadRequest(ctx, codes.InvalidParams, err)
		return
	}

	token, code, err := c.service.Login(&req)
	if err != nil {
		handlers.Unauthorized(ctx, code, err)
		return
	}

	handlers.Ok(ctx, codes.LoginSuccess, "login successful", token)
}
