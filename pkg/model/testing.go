package model

func GetUser() *User {
	return &User{Id: 1, Name: "test_1", DefaultCurrencyId: 0}
}

func GetAccount100(userid int, at AccountTypes, currency Currency) *Account {
	return &Account{1, 100, currency, "test", nil, userid, AccountType{Id: at}}
}
