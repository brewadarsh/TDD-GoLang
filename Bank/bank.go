package bank

import "errors"

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// The type defining the structure of Bank.
type Bank struct {
	// The amount of money held by the bank.
	balance int
}

// Check the balance.
func (bank *Bank) Balance() int {
	return bank.balance
}

// Withdraw money from the bank.
func (bank *Bank) Withdraw(amount int) error {
	if amount > bank.balance {
		return ErrInsufficientFunds
	}
	bank.balance -= amount
	return nil
}

// Deposit money into the bank.
func (bank *Bank) Deposit(amount int) {
	bank.balance += amount
}
