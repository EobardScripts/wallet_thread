package bankClient

import (
	"fmt"
	"sync"
)

type BankClient interface {
	// Deposit deposits given amount to clients account
	Deposit(amount int)

	// Withdrawal withdraws given amount from clients account.
	// return error if clients balance less the withdrawal amount
	Withdrawal(amount int) error

	// Balance returns clients balance
	Balance() int
}

type Wallet struct {
	mutex sync.RWMutex
	money int
}

func NewWallet() *Wallet {
	return &Wallet{
		mutex: sync.RWMutex{},
		money: 0,
	}
}

func (w *Wallet) Deposit(amount int) {
	w.mutex.Lock()
	w.money += amount
	w.mutex.Unlock()
}

func (w *Wallet) Withdrawal(amount int) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if amount > w.money {
		return fmt.Errorf("Недостаточно денег на балансе")
	}

	w.money -= amount

	return nil
}

func (w *Wallet) Balance() int {
	w.mutex.RLock()
	amount := w.money
	w.mutex.RUnlock()

	return amount
}
