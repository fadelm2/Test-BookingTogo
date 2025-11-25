package helper

import "net/http"

type AppError struct {
	StatusCode int         `json:"-"`
	Message    string      `json:"message"`
	Details    interface{} `json:"details,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewBadRequest(msg string, details interface{}) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
		Details:    details,
	}
}

func NewNotFound(msg string) *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Message:    msg,
	}
}

func NewInternal(msg string) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    msg,
	}
}
