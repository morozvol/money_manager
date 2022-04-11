package teststore

import (
	"github.com/morozvol/money_manager/pkg/model"
)

type CategoryRepository struct {
	store      *Store
	categories map[int]*model.Category
}

func (r *CategoryRepository) Create(c *model.Category) error {
	c.Id = int64(len(r.categories) + 1)
	r.categories[int(c.Id)] = c
	return nil
}

func (r *CategoryRepository) Get(userId int) ([]model.Category, error) {
	var res []model.Category

	for _, c := range r.categories {
		if c.IdOwner == userId && c.IsSystem == false {
			res = append(res, *c)
		}
	}
	return res, nil
}

func (r *CategoryRepository) GetSystem() ([]model.Category, error) {
	var res []model.Category

	for _, c := range r.categories {
		if c.IsSystem == true {
			res = append(res, *c)
		}
	}
	return res, nil
}
