package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner	string
	balance	int
}

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// method
// (accnt Account) => receiver. receiver의 첫 스펠링은 struct의 앞글자로 해야한다
// receiver의 type을 pointer로 받지 않으면 object를 복사해서 사용함.
// 값을 변경해야할 경우 pointer를 사용하고 그 외엔 pointer 사용을 안해서 안전하게 사용하는게 좋음
// Deposit x amount on your account
func (accnt *Account) Deposit(amount int) {
	accnt.balance += amount
}

// Balance of your account
func (accnt Account) Balance() int {
	return accnt.balance
}

var errNoMoney = errors.New("can't withdraw you are poor")

// Withdraw * amount from your account
func (accnt *Account) Withdraw(amount int) error {
	if accnt.balance < amount {
		return errNoMoney
	}
	accnt.balance -= amount
	return nil	// like null
}

// ChangeOwner of the account
func (accnt *Account) ChangeOwner(newOwner string) {
	accnt.owner = newOwner
}

func (accnt Account) Owner() string {
	return accnt.owner
}

// java class의 toString()과 같다고 보면 될 듯. builtin 함수
func (accnt Account) String() string {
	return fmt.Sprint(accnt.Owner(), "'s account.\nHas: ", accnt.Balance())
}
