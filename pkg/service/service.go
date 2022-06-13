package service

import (
	emulator "github.com/egorrridze/payment-emulator"
	"github.com/egorrridze/payment-emulator/pkg/repository"
)

type Authorization interface {

}

type Payment interface {
	Create(payment emulator.Payment) (int, string, error)
	GetAllById(userId int) ([]emulator.Payment, error)
	GetAllByEmail(userEmail string) ([]emulator.Payment, error)
}

type Service struct {
	Authorization
	Payment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Payment: NewPaymentService(repos.Payment),
	}
}