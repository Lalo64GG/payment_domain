package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lalo64/payment_domain/internal/payment/application"
	"github.com/lalo64/payment_domain/internal/payment/infraestructure/http/request"
	"github.com/lalo64/payment_domain/internal/shared/response"
)

type UpdateController struct {
	UpdateUseCase *application.UpdateUseCase
	Validator *validator.Validate
}

func NewUpdateController(updateUseCase *application.UpdateUseCase) *UpdateController{
	return &UpdateController{
		UpdateUseCase: updateUseCase,
		Validator: validator.New(),
	}
}

func (ctr *UpdateController) Run(ctx *gin.Context){
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

	var req request.UpdateStatusRequest

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

	

	payment, err := ctr.UpdateUseCase.Run(id, req.Status, req.ProcessAt)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Status:  false,
			Message: "Error updating payment",
			Data:    nil,
			Error:   err.Error(),
		})

		return 
	}

	ctx.JSON(http.StatusOK, response.Response{
		Status:  true,
		Message: "Payment updated successfully",
		Data:    payment,
		Error:   nil,
	})
}