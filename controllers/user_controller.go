package controllers

import (
	"icomphub-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) UserController {
	return UserController{DB: db}
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
	var users []models.User

	if err := uc.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve users",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}