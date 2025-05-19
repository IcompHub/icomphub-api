package controllers

import (
	"icomphub-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService}
}

// GetAllUsers godoc
// @Summary      List all users
// @Description  Get all users from the database
// @Tags         users
// @Produce      json
// @Success      200  {array}  model.User
// @Failure      500  {object}  gin.H
// @Router       /users [get]
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve users",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}