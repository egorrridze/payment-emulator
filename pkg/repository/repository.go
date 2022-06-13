package repository

import (
	emulator "github.com/egorrridze/payment-emulator"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {

}

type Payment interface {
	Create(payment emulator.Payment) (int, string, error)
	GetAllById(userId int) ([]emulator.Payment, error)
	GetAllByEmail(userEmail string) ([]emulator.Payment, error)
}

type Repository struct {
	Authorization
	Payment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Payment: NewPaymentPostgres(db),
	}
}
