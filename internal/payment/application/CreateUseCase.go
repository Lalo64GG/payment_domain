package application

import (
	"github.com/lalo64/payment_domain/internal/payment/application/services"
	"github.com/lalo64/payment_domain/internal/payment/domain/entities"
	"github.com/lalo64/payment_domain/internal/payment/domain/ports"
)

type CreateUseCase struct {
	PaymentRepository ports.IPaymentRepository
	UUIDGen services.UUIDGenerator
}

func NewCreateUseCase(paymentRepository ports.IPaymentRepository, uuidGen services.UUIDGenerator) *CreateUseCase {
	return &CreateUseCase{
		PaymentRepository: paymentRepository,
		UUIDGen: uuidGen,
	}
}

func (uc *CreateUseCase) Run(newPayment entities.Payment) error {

	newPayment.TransactionID = uc.UUIDGen.GenerateUUID()

	if err := uc.PaymentRepository.Create(&newPayment); err != nil {
		return err
	}

	return nil
}