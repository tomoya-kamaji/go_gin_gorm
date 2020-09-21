package errors

import (
	"errors"
)

// Note: 各層でのエラーハンドラーを作ると、層ごとで型が生まれてどこで変換するか面倒なので
//共通のエラー型を用意

type AppError struct {
	Errors []string `json:"errors"`
}

func NewAppError(message string) AppError {
	var errorResult []string

	err := errors.New(message)
	errorResult = append(errorResult, err.Error())

	return AppError{Errors: errorResult}
}

func (appErr *AppError) HasErrors() bool {
	if appErr == nil {
		return false
	}
	return len(appErr.Errors) >= 1
}

func (appErr *AppError) Concat(other *AppError) AppError {
	if other.Errors == nil {
		return *appErr
	}

	var errors []string
	errors = append(append(errors, appErr.Errors...), other.Errors...)

	return newAppErrors(errors)
}

func newAppErrors(messages []string) AppError {
	return AppError{Errors: messages}
}
