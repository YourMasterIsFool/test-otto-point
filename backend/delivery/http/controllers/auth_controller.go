package controllers

import (
	usecase "backend/applications/usecase"
	authdto "backend/dtos/auth"
	response "backend/pkg/response"
	validation "backend/pkg/validation"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	ROUTE_API string
	usecase   usecase.AuthUsecase
}

func NewAuthController(usecase usecase.AuthUsecase, ROUTE_API string) *AuthController {
	return &AuthController{
		ROUTE_API: ROUTE_API,
		usecase:   usecase,
	}
}

// login
func (controller *AuthController) Login(c *fiber.Ctx) error {

	var req authdto.LoginDto

	//validated json request
	if err := c.BodyParser(&req); err != nil {
		return response.NewErrorResponseHandler(c, 400, "failed to parse body parse", err)
	}

	//validated schem
	if err := validator.New().Struct(req); err != nil {
		validationError := validation.ErrorValidation(err)
		return response.NewErrorResponseHandler(c, 422, "Validation Error", validationError)

	}

	//login logic usecase
	data, err := controller.usecase.SignIn(req)
	if err != nil {
		errorValidation := err.(*response.ErrorResponse)
		return response.NewErrorResponseHandler(c, errorValidation.Code, errorValidation.Message, nil)
	}

	return response.NewSuccessResponse(c, data, "Successfully login", 200)

}

func (controller *AuthController) Save(c *fiber.Ctx) error {

	var req authdto.CreateAuthDto

	//validated parsing json
	if err := c.BodyParser(&req); err != nil {
		return response.NewErrorResponseHandler(c, 400, "failed to parse body parse", err)
	}

	// validated schema json create
	if err := validator.New().Struct(req); err != nil {
		validationError := validation.ErrorValidation(err)
		return response.NewErrorResponseHandler(c, 422, "Validation Error", validationError)

	}

	// created user
	data, err := controller.usecase.Save(req)
	if err != nil {

		// error when failed to create users
		errorValidation := err.(*response.ErrorResponse)
		return response.NewErrorResponseHandler(c, errorValidation.Code, errorValidation.Message, nil)
	}

	return response.NewSuccessResponse(c, data, "Successfully Created User", 200)

}

func (controller *AuthController) StartController(app *fiber.App) {
	// defifining contrller
	app.Post(fmt.Sprintf("%s/login", controller.ROUTE_API), controller.Login)
	app.Post(fmt.Sprintf("%s/register", controller.ROUTE_API), controller.Save)

}
