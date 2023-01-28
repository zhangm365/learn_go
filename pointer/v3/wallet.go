package main

import (
	"errors"
	"fmt"
)

// 定义一个新类型
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// 参数是指针类型
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Println("the address of the w.balance in Deposit is ", &w.balance)
	w.balance += amount
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

// 返回值是 error
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

/*
	// https://pkg.go.dev/fmt#Stringer
type Stringer interface {
	String() string
}*/
// satisfy the Stringer interface
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
