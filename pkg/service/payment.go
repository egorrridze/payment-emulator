package service

import (
	emulator "github.com/egorrridze/payment-emulator"
	"github.com/egorrridze/payment-emulator/pkg/repository"
)

type PaymentService struct {
	repo repository.Payment
}

func NewPaymentService(repo repository.Payment) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) Create(payment emulator.Payment) (int, string, error) {
	return s.repo.Create(payment)
}

func (s *PaymentService) GetAllById(userId int) ([]emulator.Payment, error) {
	return s.repo.GetAllById(userId)
}

func (s *PaymentService) GetAllByEmail(userEmail string) ([]emulator.Payment, error) {
	return s.repo.GetAllByEmail(userEmail)
}