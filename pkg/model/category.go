package model

import (
	"fmt"
)

type Category struct {
	Id       int                  `db:"id" json:"id"`
	Name     string               `db:"name" json:"name"`
	Type     OperationPaymentType `db:"type" json:"type"`
	IdOwner  int                  `db:"id_owner" json:"id_owner"`
	IdParent int                  `db:"id_parent_category" json:"id_parent_category"`
	IsEnd    bool                 `db:"is_end" json:"is_end"`
	IsSystem bool                 `db:"is_system" json:"is_system"`
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

func (cs Categories) DeleteById(id int) Categories {
	index := 0
	for i, c := range cs {
		if c.Id == id {
			index = i
		}
	}
	return append(cs[:index], cs[index+1:]...)
}
