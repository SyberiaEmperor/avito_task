package repository

import (
	"github.com/SyberiaEmperor/avito_task/models"
	"gorm.io/gorm"
)

type Account interface {
	GetAccountInfo(accountId int) (float64, error)
	Deposit(req models.AccountRequest) error
	Debit(req models.AccountRequest) error
	Transfer(req models.TransferRequest) error
}

type Repository struct {
	Account
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Account: NewAccountService(db),
	}
}
