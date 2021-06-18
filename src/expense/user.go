package expense

import (
	"errors"
	"time"
)

type User struct {
	name      string
	balance   float64
	CreatedAt time.Time
}

func (w *User) Name() string {
	return w.name
}

func (w *User) SetName(name string) {
	w.name = name
}

func (w *User) Balance() float64 {
	return w.balance
}

func (w *User) SetBalance(balance float64) {
	w.balance = balance
}

func (w *User) GetCreatedAt() time.Time {
	return w.CreatedAt
}

func CreateWallet(name string, balance float64) (*User,error) {
	if balance>0 {
		return &User{name: name, balance: balance, CreatedAt: time.Now()},nil
	}
	return nil, errors.New("zero balance")
}

func (w *User) CreditBalance(balance float64) {
	w.balance += balance
}

func (w *User) DebitBalance(balance float64) {
	w.balance -= balance
}