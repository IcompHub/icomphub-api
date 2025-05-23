package controllers

import (
	"net/http"

	"icomphub-api/models"
	"icomphub-api/services"

	"github.com/gin-gonic/gin"
)

type TechnologyController struct {
	services services.TechnologyService
}

func NewTechnologyController(s services.TechnologyService) TechnologyController {
	return TechnologyController{services: s}
}

func (tc *TechnologyController) CreateTechnology(c *gin.Context) {
	var tech models.Technology

	if err := c.ShouldBindJSON(&tech); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := tc.services.CreateTechnology(&tech); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create technology"})
		return
	}

	c.JSON(http.StatusCreated, tech)
}
