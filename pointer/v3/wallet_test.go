package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		// Bitcoin新类型的用法
		wallet.Deposit(Bitcoin(10))

		// fmt.Println("the address of the w.wallet in Test is ", &wallet.balance)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {

		wallet := Wallet{Bitcoin(20)}
		// 新类型的用法
		err := wallet.Withdraw(Bitcoin(10))
		// fmt.Println(err) // output <nil>
		// fmt.Println("the address of the w.wallet in Test is ", &wallet.balance)

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(10)
		wallet := Wallet{startBalance}
		// 提现大于现有金额
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startBalance)
		assertError(t, err, InsufficientFundsError)

	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	// fmt.Println("the address in assertBalance is ", &wallet.balance)
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}
}

func assertError(t *testing.T, got error, want error) {

	if got == nil {
		// t.Fatal 如果被调用，它将停止测试
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got '%s', but want '%s'", got, want)
	}
}
