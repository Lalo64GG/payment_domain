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
	ProcessedAt   time.Time
	TransactionID string
}

// * Comportamientos del dominio
func (p *Payment) Process() ( *events.PaymentProcessed, error) {
	

	if p.ID <= 0{
		return nil, errors.New("payment must have a valid ID to be processed")
	}

	if p.BookingID <= 0 {
		return nil, errors.New("payment must be associated with a valid booking")
	}

	if p.Amount <= 0{
		return nil, errors.New("payment amount must be greater than zero")
	}

	if p.Status != StatusPending {
		return nil, fmt.Errorf("only pending payments can be processed")
	}


	now := time.Now()
	p.Status= StatusSuccess
	p.ProcessedAt = now

	event := &events.PaymentProcessed{
		PaymentID:   p.ID,
		UserID:      p.UserID,
		BookingID:   p.BookingID,
		Amount:      p.Amount,
		Currency:    p.Currency,
		ProcessedAt: now,
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
