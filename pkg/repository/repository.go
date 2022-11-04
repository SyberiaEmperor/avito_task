package repository

type Account interface {
	GetAccountInfo() error
	Deposit() error
	Debit() error
	Transfer() error
}

type Repository struct {
	Account
}

func NewRepository() *Repository {
	return &Repository{
		Account: NewAccountService(),
	}
}