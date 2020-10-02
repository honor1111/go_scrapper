package accounts

import (
	"errors"
	"fmt"
)

// Account strtuct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't Withdraw! not enough deposit")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit Method x amount on your account
// a Account is receiver, write first lower case
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance would print balance
func (a Account) Balance() int {
	return a.balance
}

// Withdraw from your balance
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner change owner
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner show owner
func (a Account) Owner() string {
	return a.owner
}

func (a *Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
