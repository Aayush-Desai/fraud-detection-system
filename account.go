package main

type Account struct {
	id string
	bankName string
	user *User
	isAccountBlacklisted bool
	transactionLimit float64
}

func GetNewAccount(id string, bankName string, user *User, transactionLimit float64) *Account {
	return &Account{
		id: id,
		bankName: bankName,
		user: user,
		transactionLimit: transactionLimit,
	}
}

func (a *Account) GetTransactionLimit() float64 {
	return a.transactionLimit
}

func (a *Account) IsAccountBlacklisted() bool {
	return a.isAccountBlacklisted
}

func (a *Account) SetAccountBlacklisted(isAccountBlacklisted bool) {
	a.isAccountBlacklisted = isAccountBlacklisted
}