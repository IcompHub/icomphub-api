package controllers

import (
	"net/http"

	"icomphub-api/dto"

	"icomphub-api/services"

	"github.com/gin-gonic/gin"
)

type TechnologyController struct {
	services services.TechnologyService
}

func NewTechnologyController(s services.TechnologyService) TechnologyController {
	return TechnologyController{services: s}
}

// CreateTechnology godoc
// @Summary      Cria uma nova tecnologia
// @Description  Adiciona uma nova tecnologia ao sistema com base nos dados fornecidos.
// @Tags         Technologies
// @Accept       json
// @Produce      json
// @Param        technology  body      dto.TechnologyDTO  true  "Dados da Tecnologia para criar"
// @Success      201           {object}  models.Technology  "Tecnologia criada com sucesso"
// @Failure      400         {object}  models.ErrorResponse   "Erro: Entrada inválida" // <--- Alterado aqui
// @Failure      500         {object}  models.ErrorResponse   "Erro: Falha ao criar tecnologia" // <--- Alterado aqui
// @Router       /technologies [post]
func (tc *TechnologyController) CreateTechnology(c *gin.Context) {
	var techDTO dto.TechnologyDTO

	if err := c.ShouldBindJSON(&techDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := tc.services.CreateTechnology(&techDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create technology", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, techDTO)
}
