package handlers

import (
	"net/http"

	"icomphub-api/codes"
	"icomphub-api/dtos"

	"github.com/gin-gonic/gin"
)

func errorHanlder(c *gin.Context, status int, code codes.Code, message string) {
	c.JSON(status, dtos.Response[any]{
		Success: false,
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func successHanlder(c *gin.Context, status int, code codes.Code, message string, data any) {
	c.JSON(status, dtos.Response[any]{
		Success: true,
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func abortHanlder(c *gin.Context, status int, code codes.Code, message string) {
	c.AbortWithStatusJSON(status, dtos.Response[any]{
		Success: false,
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

func Ok(c *gin.Context, code codes.Code, message string, data any) {
	successHanlder(c, http.StatusOK, code, message, data)
}

func BadRequest(c *gin.Context, code codes.Code, err error) {
	errorHanlder(c, http.StatusBadRequest, code, err.Error())
}

func InternalServerError(c *gin.Context, code codes.Code, err error) {
	errorHanlder(c, http.StatusInternalServerError, code, err.Error())
}

func NotFound(c *gin.Context, code codes.Code, err error) {
	errorHanlder(c, http.StatusNotFound, code, err.Error())
}

func Forbidden(c *gin.Context, code codes.Code, err error) {
	errorHanlder(c, http.StatusForbidden, code, err.Error())
}

func Unauthorized(c *gin.Context, code codes.Code, err error) {
	errorHanlder(c, http.StatusUnauthorized, code, err.Error())
}

func AbortForbidden(c *gin.Context, code codes.Code, err error) {
	abortHanlder(c, http.StatusForbidden, code, err.Error())
}

func AbortUnauthorized(c *gin.Context, code codes.Code, err error) {
	abortHanlder(c, http.StatusUnauthorized, code, err.Error())
}
