package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {

}

func NewUserController() userController {
	return userController{
		
	}
}

// @Summary      Get Users
// @Description  Get all users from database
// @Tags         users
// @Accept       json
// @Produce      json
// @Router       /users/ [get]
func (p *userController) GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "helloo"})
}