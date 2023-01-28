package main

import "fmt"

type Wallet struct {
	balance int
}

// 参数是指针类型
func (w *Wallet) Deposit(amount int) {
	// output the address
	fmt.Println("the address of the w.balance in Deposit is ", &w.balance)
	w.balance += amount

}

func (w *Wallet) Balance() int {
	return w.balance
}
