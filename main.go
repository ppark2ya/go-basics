package main

import (
	"fmt"

	"github.com/ppark2ya/go-basics/accounts"
)

func main() {
	account := accounts.NewAccount("jtpark")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		// Println을 실행하고 프로그램을 종료함
		// log.Fatalln(err)
		fmt.Println(err)
	}
	// fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
