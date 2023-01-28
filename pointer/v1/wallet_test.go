package main

import "testing"

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	// Wallet 的副本
	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10
	// got != want, error
	if got != want {
		t.Errorf("got %d, but want %d", got, want)
	}

}
