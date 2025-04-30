package validation

import (
	"backend/pkg/response"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

// fomrmationg error validatio
func ErrorValidation(err error) map[string]string {
	validationsErr := err.(validator.ValidationErrors)

	errs := make(map[string]string)

	for _, fieldErr := range validationsErr {
		errs[fieldErr.Field()] = fieldErr.ActualTag()
	}

	return errs
}

func ValidationRequest(c *fiber.Ctx, req interface{}) error {

	if err := c.BodyParser(&req); err != nil {
		return response.NewErrorResponseHandler(c, 400, "Json tidak valid", nil)
	}

	if err := validator.New().Struct(req); err != nil {
		validationError := ErrorValidation(err)
		return response.NewErrorResponseHandler(c, 422, "validation error", validationError)
	}

	return nil
}
