package category_tree

import (
	"fmt"
	"github.com/morozvol/money_manager/pkg/model"
)

type CategoryTree struct {
	Root *Node
	m    map[int]*Node
}

func NewCategoryTree(root *Node) *CategoryTree {
	ct := &CategoryTree{
		Root: root,
		m:    make(map[int]*Node),
	}
	ct.m[root.Category.Id] = root
	return ct
}

func (n CategoryTree) GetCategoryById(id int) (*model.Category, error) {
	v, ok := n.m[id]
	if ok {
		return v.Category, nil
	}
	return nil, fmt.Errorf("Category with id=%d does not exist ", id)
}

func (n CategoryTree) FindNode(id int) (*Node, error) {
	v, ok := n.m[id]
	if ok {
		return v, nil
	}
	return nil, fmt.Errorf("Category with id=%d does not exist ", id)
}

func (n CategoryTree) DeleteNode(node *Node) {
	idParent := node.Category.IdParent
	delete(n.m, node.Category.Id)
	cs := n.m[idParent].children
	index := 0
	for i, c := range cs {
		if c.Category.Id == node.Category.Id {
			index = i
		}
	}
	n.m[idParent].children = append(cs[:index], cs[index+1:]...)
}
