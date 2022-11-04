package repository

type AccountService struct {

}

func NewAccountService() *AccountService {
	return &AccountService{}
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