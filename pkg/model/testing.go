package model

func GetUser() *User {
	return &User{Id: 1, Name: "test_1", DefaultCurrencyId: 0}
}

func GetAccount100(userid int, at AccountTypes, currency Currency) *Account {
	return &Account{1, 100, currency, "test", nil, userid, AccountType{Id: at}}
}

func GetCategory(t OperationPaymentType, idOwner, idParent int, isEnd, isSystem bool) *Category {
	return &Category{Name: "test", Type: t, IdOwner: idOwner, IdParent: idParent, IsEnd: isEnd, IsSystem: isSystem}
}

func GetOperation(a Account, sum float32, category Category) *Operation {
	return &Operation{IdAccount: a.Id, Sum: sum, Category: category, Description: ""}
}
