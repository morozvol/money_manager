package sqlstore

import (
	"database/sql"
	"github.com/morozvol/money_manager/pkg/model"
)

type CategoryRepository struct {
	store *Store
}
type category struct {
	Id       int                        `db:"id"`
	Name     string                     `db:"name"`
	Type     model.OperationPaymentType `db:"type"`
	IdOwner  sql.NullInt64              `db:"id_owner"`
	IdParent sql.NullInt64              `db:"id_parent_category"`
	IsEnd    bool                       `db:"is_end"`
	IsSystem bool                       `db:"is_system"`
}

func (c category) toModel() model.Category {
	return model.Category{Id: c.Id, Name: c.Name, Type: c.Type, IdOwner: int(c.IdOwner.Int64), IdParent: int(c.IdParent.Int64), IsEnd: c.IsEnd}
}

func (r *CategoryRepository) Create(c *model.Category) error {
	var idOwner sql.NullInt64
	var idParent sql.NullInt64
	if c.IdParent != 0 {
		idParent.Int64 = int64(c.IdParent)
		idParent.Valid = true
	}
	if c.IdOwner != 0 {
		idOwner.Int64 = int64(c.IdOwner)
		idOwner.Valid = true
	}
	lastInsertId := 0
	err := r.store.db.QueryRow("INSERT INTO category (name, type, id_owner, id_parent_category, is_end, is_system) VALUES($1,$2,$3,$4,$5,$6) RETURNING id;", c.Name, c.Type, idOwner, idParent, c.IsEnd, c.IsSystem).Scan(&lastInsertId)
	if err != nil {
		return err
	}
	c.Id = int(lastInsertId)
	return nil
}

func (r *CategoryRepository) Get(userId int) ([]model.Category, error) {
	c := category{}
	res := make([]model.Category, 0)
	rows, err := r.store.db.Queryx("SELECT id, name, type, id_owner, id_parent_category, is_end FROM category WHERE (id_owner IS NULL OR id_owner = $1) AND is_system = false;", userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&c)
		if err != nil {
			return nil, err
		}
		res = append(res, c.toModel())
	}
	return res, nil
}

func (r *CategoryRepository) GetSystem() ([]model.Category, error) {
	c := category{}
	res := make([]model.Category, 0)
	rows, err := r.store.db.Queryx("SELECT id, name, type, id_owner, id_parent_category, is_end FROM category WHERE is_system = true;")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.StructScan(&c)
		if err != nil {
			return nil, err
		}
		res = append(res, c.toModel())
	}
	return res, nil
}
