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
	res := acc.db.First(&account, accountId)
	return account.Balance, res.Error
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

func (acc *AccountService) Transfer(req models.TransferRequest) error {
	var senderAccount models.Account
	var receiverAccount models.Account

	res := acc.db.First(&senderAccount, req.SenderID)

	if res.Error != nil {
		return fmt.Errorf("invalid sender id")
	}

	res = acc.db.First(&receiverAccount, req.ReceiverID)

	if res.Error != nil {
		return fmt.Errorf("invalid receiver id")
	}

	if senderAccount.Balance-req.Amount < 0 {
		return fmt.Errorf("not enough money for transaction")
	}

	senderAccount.Balance -= req.Amount
	receiverAccount.Balance += req.Amount

	acc.db.Save(&senderAccount)
	acc.db.Save(&receiverAccount)

	acc.db.Save(&models.Transaction{SenderID: req.SenderID, ReceiverID: req.ReceiverID, Amount: req.Amount})

	return nil
}
