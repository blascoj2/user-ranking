package domain

import (
	"fmt"
	"runtime"
)

const (
	ServerErrorTag         string = "server-error"
	InvalidRequestErrorTag string = "invalid-request-error"
	UserErrorTag           string = "user-error"

	NotFoundCode        string = "resource-not-found"
	UnexpectedErrorCode string = "unexpected-error"
	BadRequestCode      string = "bad-request"
)

type ErrInterface interface {
	Code() string
	Tag() string
	Message() string
}

type DomainError struct {
	message string
	code    string
	tag     string
	stack   []uintptr
}

func (e DomainError) Error() string {
	return fmt.Sprintf("code: %s message: %s", e.code, e.message)
}

func (e DomainError) Code() string {
	return e.code
}

func (e DomainError) Tag() string {
	return e.tag
}

func (e DomainError) Message() string {
	return e.message
}
func (e *DomainError) Unwrap() error {
	return e
}

func (e DomainError) StackTrace() []uintptr {
	f := make([]uintptr, len(e.stack))
	for i := 0; i < len(f); i++ {
		f[i] = (e.stack)[i]
	}
	return f
}

func callers() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	stack := pcs[0:n]
	return stack
}

func NewError(message string, code string, tag string) DomainError {
	return DomainError{
		message: message,
		code:    code,
		tag:     tag,
		stack:   callers(),
	}
}

type ValidationErr struct {
	DomainError
}

type NotFoundErr struct {
	DomainError
}

func NewValidationErr(msg string, tag string) error {
	return ValidationErr{
		DomainError: NewError(msg, BadRequestCode, tag),
	}
}

func NewUserNotFoundErr(userId UserId) error {
	msg := fmt.Sprintf("user not found with id: %s", userId)
	return NotFoundErr{
		DomainError: NewError(msg, NotFoundCode, UserErrorTag),
	}
}
