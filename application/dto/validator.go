package dto

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidateResult struct {
	msg    string
	detail error
}

func (r ValidateResult) Msg() string {
	return r.msg
}

func (r ValidateResult) Detail() error {
	return r.detail
}

func Validate(v *validator.Validate, object any) *ValidateResult {
	err := v.Struct(object)

	if validateErrs, ok := errors.AsType[validator.ValidationErrors](err); ok {
		for _, e := range validateErrs {

			return &ValidateResult{
				msg:    fmt.Sprintf("invalid value of %s, expected type: %s", e.StructNamespace(), e.Type()),
				detail: err,
			}
		}

		return &ValidateResult{
			detail: err,
		}
	}

	return nil
}
