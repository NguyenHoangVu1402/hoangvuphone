package validations

import (
	"regexp"
	

	"github.com/go-playground/validator/v10"
	"hoangvuphone/internal/dtos"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	_ = validate.RegisterValidation("slug", validateSlug)
}

func validateSlug(fl validator.FieldLevel) bool {
	slug := fl.Field().String()
	matched, _ := regexp.MatchString(`^[a-z0-9]+(?:-[a-z0-9]+)*$`, slug)
	return matched
}

func ValidateCreateRole(input dtos.CreateRoleRequest) error {
	return validate.Struct(input)
}

func ValidateUpdateRole(input dtos.UpdateRoleRequest) error {
	return validate.Struct(input)
}