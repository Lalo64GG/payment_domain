package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lalo64/payment_domain/internal/payment/infraestructure/http"
)

func Routes(router *gin.RouterGroup){
	createController := http.SetUpCreate()
	getByIdController := http.SetUpGetById()
	updateController := http.SetUpUpdate()
	processController := http.SetUpProcess()


	router.POST("/process/:id", processController.Run)
	router.POST("/", createController.Run)
	router.GET("/:id", getByIdController.Run)
	router.PATCH("/:id", updateController.Run)
}