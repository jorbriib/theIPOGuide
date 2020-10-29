package application

import "fmt"

type AppError struct{
	message string
}

func NewAppError(message string) *AppError {
	return &AppError{message: message}
}

func (e *AppError) Error() string {
	return fmt.Sprintf(e.message)
}