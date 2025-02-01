package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Custom validation function for various fields
func UsernameValidator(fl validator.FieldLevel) bool {
	// Regex to allow alphanumeric characters and underscores in username
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return re.MatchString(fl.Field().String())
}

// Register all of the custom validators
func RegisterValidators(validate *validator.Validate) {
	validate.RegisterValidation("usernamevalid", UsernameValidator)
}
