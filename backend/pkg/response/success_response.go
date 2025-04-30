package response

import "github.com/gofiber/fiber/v2"

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

func NewSuccessResponse(c *fiber.Ctx, data interface{}, message string, code int) error {
	return c.Status(200).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}
