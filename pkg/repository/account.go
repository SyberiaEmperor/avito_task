package repository

import (
	"fmt"

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
	return account.Balance, nil
}

func (acc *AccountService) Deposit(req models.AccountRequest) error {
	var account models.Account
	res := acc.db.First(&account, req.ID)

	if res.Error != nil {

		if req.Income <= 0 {
			return fmt.Errorf("negative deposite")
		}

		var account = models.Account{ID: uint(req.ID), Balance: req.Income}
		res = acc.db.Create(&account)

		return res.Error
	} else {
		if account.Balance+req.Income < 0 {
			return fmt.Errorf("deposite less than balance")
		}

		account.Balance += req.Income
		res = acc.db.Save(&account)

		return res.Error
	}
}

func (acc *AccountService) Debit(req models.AccountRequest) error {
	return nil
}

func (acc *AccountService) Transfer() error {
	return nil
}
