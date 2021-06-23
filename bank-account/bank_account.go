package account

import "sync"


type Account struct {
	balance int
	closed bool
	sync.Mutex
}

func Open(amount int) *Account {
	if amount < 0 {
		return nil
	}

	account := &Account{
		balance: amount,
		closed: false,
	}
	return account
}


func (acc *Account) Balance() (int, bool) {
	acc.Lock()
	defer acc.Unlock()
	
	if acc.closed {
		return 0, false
	}
	return acc.balance, true
}


func (acc *Account)Close() (int, bool) {
	acc.Lock()
	defer acc.Unlock()

	if acc.closed {
		return 0, false
	}
	acc.closed = true
	return acc.balance, true
}

func (acc *Account) Deposit(amount int) (int, bool) {
	acc.Lock()
	defer acc.Unlock()

	if acc.closed || amount+acc.balance < 0 {
		return 0, false
	}
	acc.balance += amount
	return acc.balance, true
}

