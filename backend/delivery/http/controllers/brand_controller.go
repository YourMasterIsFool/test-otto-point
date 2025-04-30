package controllers

import (
	usecase "backend/applications/usecase"
	dtos "backend/dtos/brand"
	"backend/pkg/response"
	validation "backend/pkg/validation"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type BrandController struct {
	brandUsecase usecase.BrandUsecase
	ROUTE_API    string
}

func NewBrandController(uc usecase.BrandUsecase, ROUTE_API string) *BrandController {
	return &BrandController{
		brandUsecase: uc,
		ROUTE_API:    ROUTE_API,
	}
}

func (controller *BrandController) Save(c *fiber.Ctx) error {

	var req dtos.CreateBrand

	// parse CreatedBrand
	if err := c.BodyParser(&req); err != nil {
		return response.NewErrorResponseHandler(c, 400, "failed to parse body parse", err)
	}

	if err := validator.New().Struct(req); err != nil {

		// custom validation error
		validationError := validation.ErrorValidation(err)
		return response.NewErrorResponseHandler(
			c, http.StatusUnprocessableEntity, "Error validation create", validationError)
	}

	data, err := controller.brandUsecase.Save(req)

	if err != nil {

		return response.NewErrorResponseHandler(c, 400, "Error Creating", err)
	}

	return response.NewSuccessResponse(c, data, "Successfully created brand", 200)

}

func (controller *BrandController) StartController(app *fiber.App) {
	app.Post(controller.ROUTE_API, controller.Save)
}
