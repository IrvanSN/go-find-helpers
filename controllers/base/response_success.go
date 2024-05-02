package base

type SuccessResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewSuccessResponse(message string, data any) *SuccessResponse {
	return &SuccessResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}
