package dto

type SuccessResponse struct {
	Message string `json:"message"`
}

// ErrorResponse represents an error JSON response structure.
type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
