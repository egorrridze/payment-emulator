package repository

import (
	"fmt"
	emulator "github.com/egorrridze/payment-emulator"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"time"
)

type PaymentPostgres struct {
	db *sqlx.DB
}

func NewPaymentPostgres(db *sqlx.DB) *PaymentPostgres {
	return &PaymentPostgres{db: db}
}

func (r *PaymentPostgres) Create(payment emulator.Payment) (int, string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, "", err
	}

	var id int
	var status string
	errorChance := rand.Intn(100)
	if errorChance < 80 {
		status = "НОВЫЙ"
	} else {
		status = "ОШИБКА"
	}
	createListQuery := fmt.Sprintf("INSERT INTO %s (user_id, user_email, summ, currency, creation_time, update_time, status) " +
		"VALUES ($1, $2, $3, $4, $5, $5, $6) RETURNING id", paymentsTable)
	row := tx.QueryRow(createListQuery, payment.UserId, payment.UserEmail, payment.Sum, payment.Currency, time.Now(), status)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, "", err
	}

	return id, status, tx.Commit()
}

func (r *PaymentPostgres) GetAllById(userId int) ([]emulator.Payment, error) {
	var payments []emulator.Payment
	query := fmt.Sprintf("SELECT * FROM %s p WHERE p.user_id = $1", paymentsTable)
	err := r.db.Select(&payments, query, userId)

	return payments, err
}

func (r *PaymentPostgres) GetAllByEmail(userEmail string) ([]emulator.Payment, error) {
	var payments []emulator.Payment
	query := fmt.Sprintf("SELECT * FROM %s p WHERE p.user_email = $1", paymentsTable)
	err := r.db.Select(&payments, query, userEmail)

	return payments, err
}