package application

import (
	"github.com/lalo64/payment_domain/internal/payment/domain/entities"
	"github.com/lalo64/payment_domain/internal/payment/domain/ports"
)

type GetByIdUseCase struct {
	PaymentRepository ports.IPaymentRepository
}

func NewGetByIdUseCase(paymentRepositort ports.IPaymentRepository) *GetByIdUseCase{
	return &GetByIdUseCase{PaymentRepository: paymentRepositort}
}

func (uc *GetByIdUseCase) Run(id int64) (entities.Payment ,error) {

	payment, err := uc.PaymentRepository.GetByID(id)

	if err != nil {
		return entities.Payment{}, err
	}


	return *payment, nil
}