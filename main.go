package main

import (
	"fmt"

	"github.com/ppark2ya/go-basics/accounts"
	"github.com/ppark2ya/go-basics/mydict"
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

	dictionary := mydict.Dictionary{"fruit": "apple"}
	// dictionary["hello"] = "hello world"
	// fmt.Println(dictionary)
	isExists := dictionary.Add("fruit", "apple")
	fmt.Println("isExists:", isExists)

	dictionary.Add("second", "banana")
	definition, err := dictionary.Search("second")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("definition: ", definition)
	}
}
