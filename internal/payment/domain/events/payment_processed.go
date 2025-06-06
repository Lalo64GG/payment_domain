package events

import "time"

type PaymentProcessed struct {
	PaymentID   int64
	UserID      int64
	BookingID   int64
	Amount      float64
	Currency    string
	ProcessedAt time.Time
}
