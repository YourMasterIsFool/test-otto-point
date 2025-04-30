package controllers

import (
	"backend/applications/usecase"
	transactiondto "backend/dtos/transaction"
	"backend/middleware"
	"backend/pkg/jwt"
	"backend/pkg/response"
	validation "backend/pkg/validation"
	"fmt"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	transactionUsecase usecase.TransactionUsecase
	ROUTE_API          string
}

func NewTransactionController(transactionUsecase usecase.TransactionUsecase, route_api string) *TransactionController {
	return &TransactionController{
		ROUTE_API:          route_api,
		transactionUsecase: transactionUsecase,
	}
}

// crate  transaction
func (controller *TransactionController) Save(c *fiber.Ctx) error {

	userID, err := jwt.ExtractUserIDFromJWT(c)

	if err != nil {
		return response.NewErrorResponseHandler(c, 400, "Failed to extract user ID from JWT", nil)
	}

	// ddto create transaction
	var dto transactiondto.CreateTransactionDto

	// validation json
	if err := c.BodyParser(&dto); err != nil {
		return response.NewErrorResponseHandler(c, 400, "Format Json Tidak Valid", nil)
	}

	// validation input
	if err := validator.New().Struct(dto); err != nil {
		//show error valdiation
		errorValidations := validation.ErrorValidation(err)
		return response.NewErrorResponseHandler(c, 422, "Validation Error", errorValidations)
	}

	// created transaction
	data, err := controller.transactionUsecase.Save(dto, userID)
	if err != nil {

		// parsing error when created into response ErrorResponse
		errorResponse := err.(*response.ErrorResponse)
		return response.NewErrorResponseHandler(c, errorResponse.Code, errorResponse.Message, errorResponse.Err)
	}

	return response.NewSuccessResponse(c, data, "Successfully created", 201)
}

func (controller *TransactionController) Get(c *fiber.Ctx) error {

	//query params transaction id
	transaction_id := c.Query("transaction_id")

	if transaction_id == "" {
		return response.NewErrorResponseHandler(c, 400, "filter transaction id tidak terisi", nil)

	}
	// parsing id into uint
	parseId, err := strconv.ParseUint(transaction_id, 10, 64)
	if err != nil {
		return response.NewErrorResponseHandler(c, 400, "Masukan id dalam bentuk angka", nil)
	}

	// find data transaction from usecase
	data, err := controller.transactionUsecase.Detail(uint(parseId))
	if err != nil {
		errorValidation := err.(*response.ErrorResponse)

		return response.NewErrorResponseHandler(c, errorValidation.Code, errorValidation.Message, nil)
	}
	return response.NewSuccessResponse(c, data, "Successfully get data transaction detail", 200)
}

// defining start controller
func (controller *TransactionController) StartController(app *fiber.App) {
	app.Post(fmt.Sprintf("%s/redemption", controller.ROUTE_API), middleware.JWTProtected, controller.Save)
	app.Get(fmt.Sprintf("%s/redemption", controller.ROUTE_API), middleware.JWTProtected, controller.Get)

}
