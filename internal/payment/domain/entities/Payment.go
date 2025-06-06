package entities

import (
	"errors"
	"fmt"
	"time"

	"github.com/lalo64/payment_domain/internal/payment/domain/events"
)

type PaymentStatus string

const (
	StatusPending  PaymentStatus = "PENDING"
	StatusSuccess  PaymentStatus = "SUCCESS"
	StatusFailed   PaymentStatus = "FAILED"
	StatusCanceled PaymentStatus = "CANCELED"
)

type Payment struct {
	ID            int64
	BookingID     int64
	UserID        int64
	Amount        float64
	Currency      string
	Status        PaymentStatus
	PaymentMethod string
	ProcessedAt   *time.Time
	TransactionID string
}

// * Comportamientos del dominio
func (p *Payment) Process() ( *events.PaymentProcessed, error) {
	fmt.Println("Processing payment with ID:", p.ID)
	if p.Status != StatusPending {
		fmt.Println("Payment is not pending, current status:", p.Status)
		return nil, errors.New("only pending payments can be processed")
	}

	now := time.Now()
	p.Status = StatusSuccess
	p.ProcessedAt = &now

	event := &events.PaymentProcessed{
		PaymentID:   p.ID,
		UserID:      p.UserID,
		BookingID:   p.BookingID,
		Amount:      p.Amount,
		Currency:    p.Currency,
	}

	return  event, nil
}

func (p *Payment) Refund() error {
	if p.Status != StatusSuccess {
		return errors.New("only successful payments can be refunded")
	}
	p.Status = StatusCanceled
	return nil
}

func (p *Payment) Retry() error {
	if p.Status != StatusFailed {
		return errors.New("only failed payments can be retried")
	}
	p.Status = StatusPending
	return nil
}

func (p *Payment) Cancel() error {
	if p.Status != StatusSuccess {
		return errors.New("only successful payments can be successfully ")
	}

	p.Status = StatusCanceled
	return nil
}
