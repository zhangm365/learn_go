package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}
	// 参数调用的是副本
	wallet.Deposit(10)

	fmt.Println("the address of the w.wallet in Test is ", &wallet.balance)

	got := wallet.Balance()

	want := 10
	// the address of the Deposit and Test is different.
	if got != want {

		t.Errorf("got %d, but want %d", got, want)
	}
}
