package model

type Category struct {
	Id       int64         `db:"id"`
	Name     string        `db:"name"`
	Type     OperationType `db:"type"`
	IdOwner  int           `db:"id_owner"`
	IdParent int           `db:"id_parent_category"`
	IsEnd    bool          `db:"is_end"`
}
