package data

import "fmt"

const (
	ErrorTypeFatal        = "Fatal"
	ErrorTypeError        = "Error"
	ErrorTypeValidation   = "Validation Error"
	ErrorTypeInfo         = "Info"
	ErrorTypeDebug        = "Debug"
	ErrorTypeUnauthorized = "Unauthorized"
)

type ErrorDetail struct {
	ErrorType    string `json:"errorType,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (err *ErrorDetail) Error() string {
	return fmt.Sprintf("ErrorType: %s, Error Message: %s", err.ErrorType, err.ErrorMessage)
}

type Response struct {
	Data    interface{}   `json:"data,omitempty"`
	Status  int           `json:"status,omitempty"`
	Error   []ErrorDetail `json:"errors,omitempty"`
	Message string        `json:"message,omitempty"`
}
