package controllers

import (
	usecase "backend/applications/usecase"
	voucherdto "backend/dtos/Voucher"
	response "backend/pkg/response"
	validation "backend/pkg/validation"
	"fmt"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type voucherController struct {
	ROUTE_API string
	usecase   usecase.VoucherUsecase
}

func NewVoucherController(usecase usecase.VoucherUsecase, ROUTE_API string) *voucherController {
	return &voucherController{
		ROUTE_API: ROUTE_API,
		usecase:   usecase,
	}
}

func (controller *voucherController) Save(c *fiber.Ctx) error {

	var req voucherdto.CreateVoucherDto

	//validated json request

	if err := c.BodyParser(&req); err != nil {
		return response.NewErrorResponseHandler(c, 400, "failed to parse body parse", err)
	}

	if err := validator.New().Struct(req); err != nil {
		validationError := validation.ErrorValidation(err)
		return response.NewErrorResponseHandler(c, 422, "Validation Error", validationError)

	}

	data, err := controller.usecase.Save(req)
	if err != nil {
		errorValidation := err.(*response.ErrorResponse)
		return response.NewErrorResponseHandler(c, errorValidation.Code, errorValidation.Message, nil)
	}

	return response.NewSuccessResponse(c, data, "Successfully created voucher", 200)

}

func (controller *voucherController) GetByBrand(c *fiber.Ctx) error {
	id := c.Query("id")

	// parsing id into uint
	parseId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return response.NewErrorResponseHandler(c, 400, "Masukan id dalam bentuk angka", nil)
	}

	// find data transaction from usecase
	data, err := controller.usecase.FindByBrandId(uint(parseId))
	if err != nil {
		errorValidation := err.(response.ErrorResponse)

		return response.NewErrorResponseHandler(c, errorValidation.Code, errorValidation.Message, nil)
	}
	return response.NewSuccessResponse(c, data, "Successfully get data voucher detail", 200)
}

func (controller *voucherController) GetById(c *fiber.Ctx) error {
	id := c.Query("id")

	// parsing id into uint
	parseId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return response.NewErrorResponseHandler(c, 400, "Masukan id dalam bentuk angka", nil)
	}

	// find data transaction from usecase
	data, err := controller.usecase.Detail(uint(parseId))
	if err != nil {
		errorValidation := err.(*response.ErrorResponse)

		return response.NewErrorResponseHandler(c, errorValidation.Code, errorValidation.Message, nil)
	}
	return response.NewSuccessResponse(c, data, "Successfully get data voucher detail", 200)
}

func (controller *voucherController) StartController(app *fiber.App) {

	// defifining contrller
	app.Post(controller.ROUTE_API, controller.Save)
	app.Get(controller.ROUTE_API, controller.GetById)

	app.Get(fmt.Sprintf("%s/brand", controller.ROUTE_API), controller.GetByBrand)

}
