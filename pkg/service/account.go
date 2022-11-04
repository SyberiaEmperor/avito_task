package service

import "github.com/SyberiaEmperor/avito_task/pkg/repository"

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (acc *AccountService) GetAccountInfo() error {
	return nil
}

func (acc *AccountService) Deposit() error {
	return nil
}

func (acc *AccountService) Debit() error {
	return nil
}

func (acc *AccountService) Transfer() error {
	return nil
}
