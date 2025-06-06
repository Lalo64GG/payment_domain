package http

import (
	"github.com/lalo64/payment_domain/internal/payment/application"
	"github.com/lalo64/payment_domain/internal/payment/application/services"
	"github.com/lalo64/payment_domain/internal/payment/domain/ports"
	"github.com/lalo64/payment_domain/internal/payment/infraestructure/adapters"
	"github.com/lalo64/payment_domain/internal/payment/infraestructure/http/controllers"
	"github.com/lalo64/payment_domain/internal/payment/infraestructure/http/controllers/helper"
)

var (
	PaymentRepository 	ports.IPaymentRepository
	UUIDGen 		services.UUIDGenerator
)

func init() {
	var err error 
	PaymentRepository, err = adapters.NewPaymentRepositoryMySql()

	if err != nil {
		panic("Error initializing PaymentRepository: " + err.Error())
	}

	UUIDGen, err = helper.NewUUID()
	if err != nil {
		panic("Error initializing UUIDGen: " + err.Error())
	}
}

func SetUpCreate() *controllers.CreateController{
	createUseCase := application.NewCreateUseCase(PaymentRepository, UUIDGen)
	return controllers.NewCreateController(createUseCase)
}

func SetUpGetById() *controllers.GetByIdController {
	getByIdUseCase := application.NewGetByIdUseCase(PaymentRepository)
	return controllers.NewGetByIdController(getByIdUseCase)
}

func SetUpUpdate() *controllers.UpdateController {
	updateUseCase := application.NewUpdateUseCase(&PaymentRepository)
	return controllers.NewUpdateController(updateUseCase)
}


func SetUpProcess() *controllers.ProcessController {
	processUseCase := application.NewProcessUseCaseRepository(PaymentRepository)
	return controllers.NewProcessController(processUseCase)
}