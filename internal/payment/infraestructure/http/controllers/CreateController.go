package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lalo64/payment_domain/internal/payment/application"
	"github.com/lalo64/payment_domain/internal/payment/domain/entities"
	"github.com/lalo64/payment_domain/internal/payment/infraestructure/http/request"
	"github.com/lalo64/payment_domain/internal/shared/response"
)

type CreateController struct {
	CreateUseCase *application.CreateUseCase
	Validator *validator.Validate
}

func NewCreateController(createUseCase *application.CreateUseCase) *CreateController{
	return &CreateController{
		CreateUseCase: createUseCase,
		Validator: validator.New(),
	}
}

func (ctr *CreateController) Run(ctx *gin.Context){
	var req request.PaymentRequest 

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Status:  false,
			Message: "Invalid request data",
			Data:    nil,
			Error:   err.Error(),
		})

		return
	}

	if err := ctr.Validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Status:  false,
			Message: "Validation error",
			Data:    nil,
			Error:   err.Error(),
		})

		return
	}

	newPayment := entities.Payment{
		ID:          0,
		Amount:      req.Amount,
		Currency:    req.Currency,
		Status:      entities.StatusPending,
		BookingID:   req.BookingID,
		UserID:      req.UserID,
		PaymentMethod: req.PaymentMethod,
	}

	 err := ctr.CreateUseCase.Run(newPayment)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Status:  false,
			Message: "Failed to create payment",
			Data:    nil,
			Error:   err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Status:  true,
		Message: "Payment created successfully",
		Data:    nil,
		Error:   "",
	})
}