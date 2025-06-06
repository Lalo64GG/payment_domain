package request

type PaymentRequest struct {
	Amount   float64 `json:"amount" validate:"required"`
	Currency string  `json:"currency" validate:"required"`
	BookingID int64   `json:"booking_id" validate:"required"`
	UserID   int64   `json:"user_id" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
}

type UpdateStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=SUCCESS FAILED CANCELED"`
}