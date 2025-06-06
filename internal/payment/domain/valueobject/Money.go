package valueobject

import (
	"errors"
	"fmt"
)

type Money struct {
	Amount  float64  
	Currency  string
}


func NewMoney(amount float64, currency string) (Money, error){
	if amount < 0 {
		return Money{}, errors.New("amount cannot be negative or zero")
	}

	if currency == "" {
		return Money{}, errors.New("currency cannot be empty")
	}

	return Money{Amount: amount, Currency: currency}, nil
}

func (m Money) String() string{
	return fmt.Sprintf("%.2f %s", m.Amount, m.Currency)
}

func (m Money) Equals(other Money) bool {
	return m.Amount == other.Amount && m.Currency == other.Currency
}