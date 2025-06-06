package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/payment_domain/internal/payment/application"
	"github.com/lalo64/payment_domain/internal/shared/response"
)

type GetByIdController struct {
	GetByIdUseCase *application.GetByIdUseCase
}

func NewGetByIdController(getByIdUseCase *application.GetByIdUseCase) *GetByIdController{
	return &GetByIdController{
		GetByIdUseCase: getByIdUseCase,
	}
}

func (ctr *GetByIdController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Status:  false,
			Message: "Invalid ID format",	
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	payment, err := ctr.GetByIdUseCase.Run(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Status:  false,
			Message: "Error retrieving payment",
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Status:  true,
		Message: "Payment retrieved successfully",
		Data:    payment,
		Error:   nil,
	})
}