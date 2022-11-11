package service

import (
	"github.com/SyberiaEmperor/avito_task/models"
	"github.com/SyberiaEmperor/avito_task/pkg/repository"
)

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (acc *AccountService) GetAccountInfo(accountId int) (float64,error) {
	return acc.repo.GetAccountInfo(accountId)
}

func (acc *AccountService) Deposit(req models.AccountRequest) error {
	return acc.repo.Deposit(req)
}

func (acc *AccountService) Debit(req models.AccountRequest) error {
	return acc.repo.Debit(req)
}

func (acc *AccountService) Transfer(req models.TransferRequest) error {
	return acc.repo.Transfer(req)
}
