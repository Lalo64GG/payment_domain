package application

import (
	"fmt"

	"github.com/lalo64/payment_domain/internal/payment/domain/ports"
)

type ProcessUseCaseRepository struct {
	PaymentRepository ports.IPaymentRepository
}

func NewProcessUseCaseRepository(paymentRepository ports.IPaymentRepository) *ProcessUseCaseRepository {
	return &ProcessUseCaseRepository{ PaymentRepository: paymentRepository }
}

func (uc *ProcessUseCaseRepository) Run(id int64) error {
	fmt.Println("Processing payment with ID:", id)
	
	payment, err := uc.PaymentRepository.GetByID(id)
	if err != nil {
		return err
	}

	event, err := payment.Process()
	if err != nil {
		return err
	}
	fmt.Println("Payment event processed:", event)

	updatedPayment, err := uc.PaymentRepository.Update(payment.ID, string(payment.Status), payment.ProcessedAt)
	if err != nil {
		return err
	}

	fmt.Printf("Payment processed successfully: %+v\n", event)
	fmt.Printf("Updated payment: %+v\n", updatedPayment)

	return nil
}

