package model

type Category struct {
	Id       int                  `db:"id"`
	Name     string               `db:"name"`
	Type     OperationPaymentType `db:"type"`
	IdOwner  int                  `db:"id_owner"`
	IdParent int                  `db:"id_parent_category"`
	IsEnd    bool                 `db:"is_end"`
	IsSystem bool                 `db:"is_system"`
}
