package service

import "github.com/egorrridze/payment-emulator/pkg/repository"

type Authorization interface {

}

type Payment interface {

}

type Service struct {
	Authorization
	Payment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}