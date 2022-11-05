package repository

import (
	"github.com/SyberiaEmperor/avito_task/models"
	"gorm.io/gorm"
)

type AccountService struct {
	db *gorm.DB
}

func NewAccountService(db *gorm.DB) *AccountService {
	return &AccountService{
		db: db,
	}
}

func (acc *AccountService) GetAccountInfo(accountId int) (float64, error) {

	var account models.Account
	acc.db.First(&account, accountId)
	return 0, nil
}

func (acc *AccountService) Deposit(req models.AccountRequest) error {
	return nil
}

func (acc *AccountService) Debit(req models.AccountRequest) error {
	return nil
}

func (acc *AccountService) Transfer() error {
	return nil
}
