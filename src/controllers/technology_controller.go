package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"icomphub-api/dto"
	"icomphub-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// DeleteTechnology godoc
// @Summary      Deleta uma tecnologia existente
// @Description  Remove uma tecnologia do sistema com base no ID fornecido.
// @Tags         Technologies
// @Accept       json
// @Produce      json
// @Param        id   path      uint64  true  "ID da Tecnologia a ser deletada"
// @Success      204  {object}  nil     "Tecnologia deletada com sucesso (sem conteúdo)"
// @Failure      400  {object}  models.ErrorResponse "Erro: ID inválido"
// @Failure      404  {object}  models.ErrorResponse "Erro: Tecnologia não encontrada"
// @Failure      500  {object}  models.ErrorResponse "Erro: Falha ao deletar tecnologia"
// @Router       /technologies/{id} [delete]
func (tc *TechnologyController) DeleteTechnology(c *gin.Context) {
	idParam := c.Param("id")
	technologyID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido", "details": "O ID fornecido não é um número válido."})
		return
	}

	err = tc.services.DeleteTechnology(technologyID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tecnologia não encontrada", "details": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao deletar tecnologia", "details": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// UpdateTechnology godoc
// @Summary      Atualiza uma tecnologia existente
// @Description  Atualiza os dados de uma tecnologia com base no ID e nos dados fornecidos.
// @Tags         Technologies
// @Accept       json
// @Produce      json
// @Param        technology  body      dto.TechnologyGetDTO  true  "Dados da Tecnologia para atualizar"
// @Success      204         {object}  nil                   "Tecnologia atualizada com sucesso (sem conteúdo)"
// @Failure      400         {object}  models.ErrorResponse  "Erro: Entrada inválida ou ID inválido"
// @Failure      404         {object}  models.ErrorResponse  "Erro: Tecnologia não encontrada"
// @Failure      500         {object}  models.ErrorResponse  "Erro: Falha ao atualizar tecnologia"
// @Router       /technologies [put]
func (tc *TechnologyController) UpdateTechnology(c *gin.Context) {
	var techGetDTO dto.TechnologyGetDTO

	if err := c.ShouldBindJSON(&techGetDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := tc.services.UpdateTechnology(&techGetDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update technology", "details": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetTechnologies godoc
// @Summary      List all technologies
// @Description  Get all technologies from the database
// @Tags         Technologies
// @Produce      json
// @Success      200  {array}  dto.TechnologyGetDTO
// @Failure      500  {object}  models.ErrorResponse
// @Router       /technologies [get]
func (uc *TechnologyController) GetTechnology(c *gin.Context) {
	tecnologies, err := uc.services.GetTechnology()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve technologies",
		})
		return
	}

	c.JSON(http.StatusOK, tecnologies)
}
