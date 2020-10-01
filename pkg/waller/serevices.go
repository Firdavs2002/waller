package waller

import (
	"errors"

	"github.com/Firdavs2002/wallet/pkg/types"
)

var ErrPhoneRigistered = errors.New("phone already registered")
var ErrAmountMustBePositive = errors.New("amount must be greated than zero")
var ErrAccountNotFound = errors.New("account not found")

// Services представляет информацию о пользователе.
type Services struct {
	nextAccountID int64
	accounts      []*types.Account
	payments      []*types.Payment
}

// RegisterAccount - метод для регистрация нового прользователя.
func (s *Services) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.Phone == phone {
			return nil, ErrPhoneRigistered
		}
	}
	s.nextAccountID++
	account := &types.Account{
		ID:      s.nextAccountID,
		Phone:   phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)
	return account, nil
}

// Deposit - метод для пополнениā сùёта
func (s *Services) Deposit(accountID int64, amount types.Money) error {
	if amount <= 0 {
		return ErrAmountMustBePositive
	}
	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}
	if account == nil {
		return ErrPhoneRigistered
	}

	account.Balance += amount
	return nil
}

// FindAccoundByIdmethod ищем пользователя по ID
func (s *Services) FindAccoundByIdmethod(accoundID int64) (*types.Account, error) {
	var accound *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accoundID {
			accound = acc
			break
		}
	}
	return accound, nil
}
