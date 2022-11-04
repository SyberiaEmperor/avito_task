package service

import "github.com/SyberiaEmperor/avito_task/pkg/repository"

type Account interface {
	GetAccountInfo() error
	Deposit() error
	Debit() error
	Transfer() error
}

type Service struct {
	Account
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Account: NewAccountService(repo.Account),
	}
}