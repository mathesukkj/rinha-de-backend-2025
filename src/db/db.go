package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/mathesukkj/rinha-de-backend-2025/src/models"
)

type DB struct {
	db *sql.DB
}

func LoadDB() (*DB, error) {
	db, err := sql.Open("postgres", "host=db port=5432 user=postgres password=postgres dbname=rinha sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}

func (db *DB) CreatePayment(payment *models.CreatePaymentRequest) error {
	query := `
		INSERT INTO payments (correlation_id, amount, created_at)
		VALUES ($1, $2, $3)
	`

	_, err := db.db.Exec(query, payment.CorrelationID, payment.Amount, payment.RequestedAt)
	if err != nil {
		return err
	}

	return nil
}
