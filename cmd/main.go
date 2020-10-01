package main

import "github.com/Firdavs2002/wallet/pkg/wallet"

func main() {
	svc := &wallet.Services{}
	svc.RegisterAccount("+9920000001")
	svc.FindAccoundByIdmethod(1)
}
