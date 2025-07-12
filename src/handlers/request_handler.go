package handlers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ValidateId(c *gin.Context) (uint64, error) {
	idParam := c.Param("id")

	if idParam == "" {
		return 0, errors.New("id not defined")
	}

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return id, errors.New("id invalid")
	}

	return id, nil
}
