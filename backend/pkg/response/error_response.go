package response

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Message string      `json:"message"`
	Err     interface{} `json:"-"`
	Code    int         `json:"code"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func NewErrorResponse(code int, message string, errors interface{}) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    code,
		Err:     errors,
	}
}

func NewErrorResponseHandler(c *fiber.Ctx, code int, message string, errors interface{}) error {
	return c.Status(code).JSON(
		fiber.Map{
			"message": message,
			"errors":  errors,
		},
	)
}
