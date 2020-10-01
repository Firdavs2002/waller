package wallet

import (
	"fmt"
	"testing"

	"github.com/Firdavs2002/wallet/pkg/wallet"
)

func TestServices_RegisterAccount_success(t *testing.T) {
	svc := &wallet.Services{}
	account, err := svc.RegisterAccount("+9920000001")
	if err != nil {
		fmt.Println(account)
	}

	account, err = svc.FindAccountByIdmethod(1)
	if err != nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}

func TestServices_FindAccoundByIdmethod_notFound(t *testing.T) {
	svc := &wallet.Services{}
	account, err := svc.RegisterAccount("+9920000001")
	if err != nil {
		fmt.Println(account)
	}

	account, err = svc.FindAccountByIdmethod(2)
	if err == nil {
		t.Errorf("\ngot > %v \nwant > nil", err)
	}
}
