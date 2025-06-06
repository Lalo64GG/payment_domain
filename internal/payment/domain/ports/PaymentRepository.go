package ports

import (
	"time"

	"github.com/lalo64/payment_domain/internal/payment/domain/entities"
)

type IPaymentRepository interface {
	Create(payment *entities.Payment) error
	GetByID(id int64) (*entities.Payment, error)
	Update(id int64, status string, processAt time.Time) (*entities.Payment, error)
}