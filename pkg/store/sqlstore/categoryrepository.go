package sqlstore

import (
	"database/sql"
	"fmt"
	"github.com/morozvol/money_manager/pkg/model"
	"github.com/morozvol/money_manager/pkg/model/category_tree"
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
	c.Id = lastInsertId
	return nil
}

func (r *CategoryRepository) Get(userId int) (*category_tree.CategoryTree, error) {
	c := category{}
	var tree *category_tree.CategoryTree
	var node *category_tree.Node
	res := make([]model.Category, 0)
	rows, err := r.store.db.Queryx("SELECT id, name, type, id_owner, id_parent_category, is_end FROM category WHERE (id_owner IS NULL OR id_owner = $1)", userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.StructScan(&c)
		if err != nil {
			return nil, err
		}
		res = append(res, c.toModel())
	}

	for _, rc := range res {
		if rc.Id == 1 {
			tree = category_tree.NewCategoryTree(category_tree.NewNode(rc, tree))
			node = tree.Root
		}
	}
	FillNode(node, res, tree)
	fmt.Printf("%#v", tree)
	return tree, nil
}
func FillNode(node *category_tree.Node, res []model.Category, tree *category_tree.CategoryTree) {
	for _, cc := range res {
		if cc.IdParent == node.Category.Id {
			newNode := category_tree.NewNode(cc, tree)
			node.AddChild(newNode)
			if !cc.IsEnd {
				FillNode(newNode, res, tree)
			}
		}
	}
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

func (r *CategoryRepository) Delete(id int) error {
	_, err := r.store.db.Queryx("DELETE FROM category WHERE id = $1;", id)
	if err != nil {
		return err
	}
	return nil
}
