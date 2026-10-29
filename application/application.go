package application

import (
	"fmt"

	"github.com/ReiiSky/SharingVisionBackend/packages"
	"github.com/go-playground/validator/v10"
)

type IKernel interface {
	ArticleRepository() packages.ArticleRepository
	Validator() *validator.Validate
}

type Response[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data"`
}

func NewSuccessResponse[T any](message string, data T) Response[T] {
	return Response[T]{
		"success",
		message,
		data,
	}
}

func NewInvalidBodyResponse() Response[any] {
	return Response[any]{
		Status:  "failed",
		Message: "invalid parse body",
		Data:    nil,
	}
}

func NewInvalidBodyValueResponse(message string) Response[any] {
	return Response[any]{
		Status:  "failed",
		Message: message,
		Data:    nil,
	}
}

func NewNotFoundResponse(dataname string) Response[any] {
	return Response[any]{
		Status:  "failed",
		Message: fmt.Sprintf("%s not found", dataname),
		Data:    nil,
	}
}

func NewInternalServerErrorResponse[T any](message string) Response[any] {
	return Response[any]{
		Status:  "failed",
		Message: message,
		Data:    nil,
	}
}
