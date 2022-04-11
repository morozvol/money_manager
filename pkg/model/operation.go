package model

import "time"

type OperationPaymentType uint8
type OperationType uint8

const (
	Coming OperationPaymentType = iota + 1
	Consumption
)
const (
	Pay OperationType = iota + 1
	Transfer
)

type Operation struct {
	Id          int       `db:"id"`
	IdAccount   int       `db:"id_account"`
	Time        time.Time `db:"time"`
	Sum         float32   `db:"sum"`
	Category    Category  `db:"category"`
	Description string    `db:"description"`
}
