package account

import "sync"

type Account interface {
	Close() (payout int, ok bool)
	Balance() (balance int, ok bool)
	Deposit(amount int) (newBalance int, ok bool)
}

type accountDetails struct {
	mux     sync.Mutex
	balance int
	active	bool
}

func Open(amt int) Account {
	if amt < 0 {
		return nil
	}

	var account Account
	account = &accountDetails{
		balance: amt,
		active: true,
	}

	return account
}

func (a *accountDetails)Balance() (int, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if !a.active {
		return 0, false
	}
	return a.balance, true
}

func (a *accountDetails)Close() (int, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if !a.active {
		return 0, false
	}

	a.active = false
	return a.balance, true
}

func (a *accountDetails)Deposit(amt int) (int, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if !a.active {
		return 0, false
	}

	if a.balance + amt < 0 {
		return a.balance, false
	}

	a.balance = a.balance + amt
	return a.balance, true
}
