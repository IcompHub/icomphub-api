package dtos

import "icomphub-api/codes"

type Response[T any] struct {
	Success bool       `json:"success"`
	Code    codes.Code `json:"code" example:"USER_CREATED"`
	Message string     `json:"message" example:"User created successfully"`
	Data    T          `json:"data"`
}
