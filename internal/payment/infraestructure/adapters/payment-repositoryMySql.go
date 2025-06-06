package adapters

import (
	"database/sql"
	"log"

	"github.com/lalo64/payment_domain/cmd/db"
	"github.com/lalo64/payment_domain/internal/payment/domain/entities"
)

type PaymentRepositoryMySql struct {
	DB *sql.DB
}

func NewPaymentRepositoryMySql() (*PaymentRepositoryMySql, error) {
	db, err := db.Connect()
	if err != nil {
		panic("Error connecting to the database: " + err.Error())
	} 

	return &PaymentRepositoryMySql{DB: db}, nil
}

func (r *PaymentRepositoryMySql) Create(payment *entities.Payment) error {
	query := `INSERT INTO payments (
		id, booking_id, user_id, amount, currency, status, transaction_id, payment_method
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		log.Fatal(err, 1)
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		payment.ID,
		payment.BookingID,
		payment.UserID,
		payment.Amount,
		payment.Currency,
		payment.Status,
		payment.TransactionID,
		payment.PaymentMethod,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	payment.ID = id

	return nil
}


func (r *PaymentRepositoryMySql) GetByID(id int64) (*entities.Payment, error) {
	query := "SELECT id, booking_id, user_id, amount, currency, status, transaction_id, payment_method FROM payments WHERE id = ?"
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		log.Fatal(err, 1)
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var payment entities.Payment

	err = row.Scan(&payment.ID, &payment.BookingID, &payment.UserID, &payment.Amount, &payment.Currency, &payment.Status, &payment.TransactionID, &payment.PaymentMethod)

	if err != nil {
		return &entities.Payment{}, err
	}

	return &payment, nil
}

func (r *PaymentRepositoryMySql) Update(id int64, status string) (*entities.Payment, error) {
	query := "UPDATE payments SET status = ? WHERE id = ?"
	stmt, err := r.DB.Prepare(query)

	if err != nil {
		log.Fatal(err, 1)
	}
	defer stmt.Close()

	_, err = stmt.Exec(status, id)
	
	if err != nil {
		return &entities.Payment{}, err
	}

	payment, err := r.GetByID(id)

	if err != nil {
		return &entities.Payment{}, err
	}

	return payment, nil

}