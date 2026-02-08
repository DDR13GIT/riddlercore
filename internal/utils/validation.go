package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func FutureValidator(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch v := field.Interface().(type) {
	case time.Time:
		return v.After(time.Now().UTC())
	case *time.Time:
		if v == nil {
			return false
		}
		return v.After(time.Now().UTC())
	default:
		return false
	}
}

func NotBlankValidator(fl validator.FieldLevel) bool {
	s, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	return strings.TrimSpace(s) != ""
}

func MsgForTag(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "notblank":
		return fmt.Sprintf("%s must not be blank", e.Field())
	case "future":
		return fmt.Sprintf("%s must be a future date", e.Field())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", e.Field(), e.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of [%s]", e.Field(), e.Param())
	default:
		return fmt.Sprintf("%s is invalid", e.Field())
	}
}
