package teststore

import (
	"github.com/morozvol/money_manager/internal/model"
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

func (r *CategoryRepository) GetAll(userId int) ([]model.Category, error) {
	var res []model.Category

	for _, c := range r.categories {
		if c.IdOwner == userId {
			res = append(res, *c)
		}
	}
	return res, nil
}
