package rest

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var (
	ErrInternal           = ErrorResponse{Code: "0.1", Message: "Unknown error"}
	ErrInvalidQueryParam  = ErrorResponse{Code: "0.2", Message: "Invalid query parameter value"}
	ErrInvalidPathParam   = ErrorResponse{Code: "0.3", Message: "Invalid path parameter value"}
	ErrInvalidRequestBody = ErrorResponse{Code: "0.4", Message: "Invalid request body"}
)
