package helpers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/kulinhdev/serentyspringsedu/api/res"
)

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "evenAge":
		return "This field must be even number"
	}
	return ""
}

func CustomMessageErrors(err error, ve validator.ValidationErrors) []res.ApiError {
	if errors.As(err, &ve) {
		out := make([]res.ApiError, len(ve))
		for i, fe := range ve {
			out[i] = res.ApiError{Field: fe.Field(), Msg: msgForTag(fe.Tag())}
		}
		return out
	} else {
		return nil
	}
}
