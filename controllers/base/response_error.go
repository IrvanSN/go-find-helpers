package base

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Status:  false,
		Message: message,
	}
}
