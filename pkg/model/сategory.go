package model

import "fmt"

type Category struct {
	Id       int                  `db:"id"`
	Name     string               `db:"name"`
	Type     OperationPaymentType `db:"type"`
	IdOwner  int                  `db:"id_owner"`
	IdParent int                  `db:"id_parent_category"`
	IsEnd    bool                 `db:"is_end"`
	IsSystem bool                 `db:"is_system"`
}

type Categories []Category

func (cs Categories) GetCategoriesByIdParent(id int) Categories {
	var res Categories

	for _, c := range cs {
		if c.IdParent == id {
			res = append(res, c)
		}
	}
	return res
}

func (cs Categories) GetCategoryById(id int) (*Category, error) {

	for _, c := range cs {
		if c.Id == id {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("Category with id=%d does not exist ", id)
}
