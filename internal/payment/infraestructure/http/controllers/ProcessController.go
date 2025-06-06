package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/payment_domain/internal/payment/application"
)

type ProcessController struct {
	ProcessUseCase *application.ProcessUseCaseRepository
}

func NewProcessController(processUseCase *application.ProcessUseCaseRepository) *ProcessController {
	return &ProcessController{
		ProcessUseCase: processUseCase,
	}
}

func (c *ProcessController) Run(ctx *gin.Context)  {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		 ctx.JSON(400, gin.H{
			"status":  false,
			"message": "Invalid ID format",
			"data":    nil,
			"error":   err.Error(),
		})
		return 
	}

	err = c.ProcessUseCase.Run(id)

	if err != nil {
		return  
	}
}