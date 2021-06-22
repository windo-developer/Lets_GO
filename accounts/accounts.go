package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

var errNomoney = errors.New("Can't widthdraw")

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNomoney
	}
	a.balance -= amount

	return nil
}
