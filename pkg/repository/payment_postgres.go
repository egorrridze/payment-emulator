package repository

import (
	"fmt"
	emulator "github.com/egorrridze/payment-emulator/models"
	"github.com/jmoiron/sqlx"
	"math/rand"
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
		"VALUES ($1, $2, $3, $4, NOW(), NOW(), $5) RETURNING id", paymentsTable)
	row := tx.QueryRow(createListQuery, payment.UserId, payment.UserEmail, payment.Sum, payment.Currency, status)
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

func (r *PaymentPostgres) GetStatusById(id int) (string, error)  {
	var status string
	query := fmt.Sprintf("SELECT p.status FROM %s p WHERE p.id = $1", paymentsTable)
	err := r.db.Get(&status, query, id)

	return status, err
}

func (r *PaymentPostgres) Delete(id int) (int64, error) {
	var rowsCounter int64 = -1

	query := fmt.Sprintf("DELETE FROM %s p WHERE p.id = $1 AND (p.status = 'НОВЫЙ' OR p.status = 'ОШИБКА')", paymentsTable)
	res, err := r.db.Exec(query, id)
	if err == nil {
		rowsCounter, _ = res.RowsAffected()
	}

	return rowsCounter, err
}

func (r *PaymentPostgres) UpdateStatus(id int) (int64, string,error)  {
	var rowsCounter int64 = -1
	var newStatus string

	oldStatus, err := r.GetStatusById(id)
	if err != nil {
		return 0, "", err
	}

	successChance := rand.Intn(100)
	if successChance > 80 || oldStatus == "ОШИБКА" {
		newStatus = "НЕУСПЕХ"
	} else {
		newStatus = "УСПЕХ"
	}

	query := fmt.Sprintf("UPDATE %s SET status = $1, update_time = NOW() WHERE id = $2 and status != 'УСПЕХ' and status != 'НЕУСПЕХ'", paymentsTable)

	res, err := r.db.Exec(query, newStatus, id)
	if err == nil {
		rowsCounter, _ = res.RowsAffected()
	}

	return rowsCounter, newStatus, err
}