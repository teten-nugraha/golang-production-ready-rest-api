package utils

type SuccessResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func NewSuccessResponse(data interface{}, message string) SuccessResponse {
	return SuccessResponse{
		Status:  "success",
		Data:    data,
		Message: message,
	}
}

func NewErrorResponse(message string, err error) ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: message,
		Error:   err.Error(),
	}
}
