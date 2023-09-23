package domain

type Response struct {
	Status  string `json:"status,omitempty"      example:"ok"`
	Error   string `json:"error,omitempty"       example:"error message"`
	Token   string `json:"accessToken,omitempty" example:"token string"`
	Message string `json:"message"`
}
