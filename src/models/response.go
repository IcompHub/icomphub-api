package models

type ErrorResponse struct {
	Error   string `json: "error" example: "Mensagem de erro descritiva"`
	Details string `json: "details,omitempty" example: "Detalhes adicionais do erro, se houver"`
}
