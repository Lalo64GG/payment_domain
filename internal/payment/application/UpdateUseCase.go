package application

import (
	"github.com/lalo64/payment_domain/internal/payment/domain/entities"
	"github.com/lalo64/payment_domain/internal/payment/domain/ports"
)

type UpdateUseCase struct {
	PaymentRepository ports.IPaymentRepository
}

func NewUpdateUseCase(paymentRepository *ports.IPaymentRepository) *UpdateUseCase {
	return &UpdateUseCase{PaymentRepository: *paymentRepository}
}

func (uc *UpdateUseCase) Run(id int64, status string) (entities.Payment, error) {

	payment, err := uc.PaymentRepository.Update(id, status)

	if err != nil {
		return entities.Payment{}, err
	}

	return *payment, nil
}
