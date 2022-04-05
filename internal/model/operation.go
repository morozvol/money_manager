package model

import "time"

type OperationType uint8

const (
	Coming      OperationType = 1
	Consumption               = 2
)

type Operation struct {
	Id        int64     `db:"id"`
	IdAccount int64     `db:"id_account"`
	Time      time.Time `db:"time"`
	Sum       float32   `db:"sum"`
	Category  Category  `db:"category"`
}
