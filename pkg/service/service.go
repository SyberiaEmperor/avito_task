package service

import (
	"github.com/SyberiaEmperor/avito_task/models"
	"github.com/SyberiaEmperor/avito_task/pkg/repository"
)

type Account interface {
	GetAccountInfo(accountId int) (float64,error)
	Deposit(req models.AccountRequest) error
	Debit(req models.AccountRequest) error
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