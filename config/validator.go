package config

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CustomValidator wrapper untuk Echo
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		var messages []string
		for _, e := range err.(validator.ValidationErrors) {
			var msg string
			switch e.Tag() {
			case "required":
				msg = fmt.Sprintf("field '%s' is required", e.Field())
			case "email":
				msg = fmt.Sprintf("field '%s' must be a valid email", e.Field())
			case "alphaSpace":
				msg = fmt.Sprintf("field '%s' must contain only letters and spaces", e.Field())
			case "numeric":
				msg = fmt.Sprintf("field '%s' must be numeric", e.Field())
			case "min":
				msg = fmt.Sprintf("field '%s' must be at least %s characters long", e.Field(), e.Param())
			case "max":
				msg = fmt.Sprintf("field '%s' must be at most %s characters long", e.Field(), e.Param())
			default:
				msg = fmt.Sprintf("field '%s' is not valid", e.Field())
			}
			messages = append(messages, msg)
		}
		return fmt.Errorf(strings.Join(messages, ", "))
	}
	return nil
}

// InitValidator set validator ke echo
func InitValidator(e *echo.Echo) {
	v := validator.New()

	// contoh custom validation: name hanya huruf
	_ = v.RegisterValidation("alphaSpace", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		for _, r := range value {
			if !(r == ' ' || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
				return false
			}
		}
		return true
	})

	e.Validator = &CustomValidator{validator: v}
}
