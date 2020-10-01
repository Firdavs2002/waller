package wallet

import (
	"errors"

	"github.com/Firdavs2002/wallet/pkg/types"
)

var (
	ErrPhoneRegistered      = errors.New("phone already registered")         // ErrPhoneRegistered телефон уже регитрирован
	ErrAmountMustBePositive = errors.New("amount must be greated than zero") // счёт не может быть пустым
	ErrAccountNotFound      = errors.New("account not found")                // пользователь не найден
)

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
			return nil, ErrPhoneRegistered
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
		return ErrPhoneRegistered
	}

	account.Balance += amount
	return nil
}

// FindAccountByIdmethod ищем пользователя по ID
func (s *Services) FindAccountByIdmethod(accountID int64) (*types.Account, error) {
	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountID {
			account = acc
			break
		}
	}

	if account == nil {
		return nil, ErrAccountNotFound
	}

	return account, nil
}
