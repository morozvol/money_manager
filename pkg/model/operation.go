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
	Id          int       `db:"id"          json:"id"`
	IdAccount   int       `db:"id_account"  json:"id_account"     validate:"required"`
	Time        time.Time `db:"time"        json:"time,omitempty"`
	Sum         float32   `db:"sum"         json:"sum"            validate:"required"`
	Category    Category  `db:"category"    json:"category"       validate:"required"`
	Description string    `db:"description" json:"description"`
}
