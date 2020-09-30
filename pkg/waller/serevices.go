package waller

import (
	"github.com/Firdavs2002/waller/pkg/types"
)

// Services представляет информацию о пользователе.
type Services struct {
	nextAccountID int64
	accounts      []types.Account
	payments      []types.Payment
}

// RegisterAccount регистрация нового прользователя.
func RegisterAccount(service *Services, phone types.Phone) {
	for _, account := range service.accounts {
		if account.Phone == phone {
			return
		}
	}
	service.nextAccountID++
	service.accounts = append(service.accounts, types.Account{
		ID:      service.nextAccountID,
		Phone:   phone,
		Balance: 0,
	})
}
