package model

type Category struct {
	Id       int                  `db:"id"                 json:"id"`
	Name     string               `db:"name"               json:"name"                validate:"required"`
	Type     OperationPaymentType `db:"type"               json:"type"                validate:"required"`
	IdOwner  int                  `db:"id_owner"           json:"id_owner"`
	IdParent int                  `db:"id_parent_category" json:"id_parent_category"`
	IsEnd    bool                 `db:"is_end"             json:"is_end"              validate:"required"`
	IsSystem bool                 `db:"is_system"          json:"is_system"           validate:"required"`
}
