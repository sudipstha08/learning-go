package validators

import (
	"strings"

	// "github.com/go-playground/validator"
	"github.com/go-playground/validator/v10"
)

// CUSTOM VALIDATION FOR TITLE
func ValidateCoolTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "Cool")
}
