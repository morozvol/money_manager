package teststore

import (
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/store"
)

type CategoryRepository struct {
	store      *Store
	categories map[int]*model.Category
}

func (r *CategoryRepository) Create(c *model.Category) error {
	c.Id = len(r.categories) + 1
	r.categories[c.Id] = c
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

func (r *CategoryRepository) Delete(id int) error {
	_, ok := r.categories[id]
	if !ok {
		return store.ErrRecordNotFound
	}
	delete(r.categories, id)
	return nil
}
