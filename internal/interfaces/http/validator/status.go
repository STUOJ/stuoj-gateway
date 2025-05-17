package validator

import (
	"github.com/go-playground/validator/v10"
	"stuoj-common/domain/shared"
)

func StatusRangeValidator(fl validator.FieldLevel) bool {
	if status, ok := fl.Field().Interface().(shared.ValidatableStatus); ok {
		return status.IsValid()
	}

	return false
}
