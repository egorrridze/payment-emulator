package service

import (
	emulator "github.com/egorrridze/payment-emulator/models"
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

func (s *PaymentService) GetStatusById(id int) (string, error)  {
	return s.repo.GetStatusById(id)
}

func (s *PaymentService) Delete(id int) (int64, error) {
	return s.repo.Delete(id)
}

func (s *PaymentService) UpdateStatus(id int) (int64, string, error){
	return s.repo.UpdateStatus(id)
}