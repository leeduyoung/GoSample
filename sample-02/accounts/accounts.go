package accounts

import "errors"

// ERROR COMMNET
var errorWithDroaw = errors.New("Can't withdraw you are poor")

// Account function is account data structure
type Account struct {
	owner   string
	balance int
}

// NewAccount function is make account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit function is save money
// Receiver는 a Account로 할 경우 복사본을 만들어버림
// 따라서, 값을 변경할때는 *사용해야한다.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance is saved money in account
func (a Account) Balance() int {
	return a.balance
}

// WithDraw function is amount from your accout
func (a *Account) WithDraw(amount int) error {
	if a.balance < amount {
		return errorWithDroaw
	}

	a.balance -= amount

	return nil
}
