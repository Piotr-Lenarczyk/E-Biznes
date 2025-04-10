// controllers/payment_controller.go
package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Payment model
type Payment struct {
	ID             uint   `json:"id"`
	Amount         int    `json:"amount"`
	CardNumber     string `json:"cardNumber"`
	ExpirationDate string `json:"expirationDate"`
}

type PaymentController struct {
	DB *gorm.DB
}

func RegisterPaymentRoutes(e *echo.Echo, db *gorm.DB) {
	paymentController := &PaymentController{DB: db}

	e.POST("/payments", paymentController.CreatePayment)
}

func (pc *PaymentController) CreatePayment(ctx echo.Context) error {
	var payment Payment
	if err := ctx.Bind(&payment); err != nil {
		return err
	}

	// Basic validation
	if payment.Amount <= 0 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Amount must be greater than 0",
		})
	}
	if payment.CardNumber == "" || payment.ExpirationDate == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Card number and expiration date are required",
		})
	}

	// Simulate successful payment (you can integrate a payment gateway here later)
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "Payment Successful!",
	})
}
