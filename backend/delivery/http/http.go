package http

import (
	"backend/delivery/http/controllers"
	"backend/pkg/validation"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

func StartServerHttp() error {

	validation.InitValidator()

	app := fiber.New()

	gormService := services.NewGormService()
	// defining api

	// brand api
	brandController := controllers.NewBrandController(gormService.BrandUsecase, "/brand")
	brandController.StartController(app)

	//voucher api
	voucherController := controllers.NewVoucherController(gormService.VoucherUsecase, "/voucher")
	voucherController.StartController(app)

	// transaction api
	transactionController := controllers.NewTransactionController(gormService.TransactionUsecase, "/transaction")
	transactionController.StartController(app)

	//auth api
	authController := controllers.NewAuthController(gormService.AuthUsecase, "/auth")
	authController.StartController(app)
	// running applikasi go
	return app.Listen(":8000")
}
