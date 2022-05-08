package category_tree

import "github.com/morozvol/money_manager/pkg/model"

type Node struct {
	Category *model.Category
	children Nodes
	tree     CategoryTree
}

type Nodes []*Node

func (n Node) GetChildren() Nodes {
	return n.children
}

func (nodes Nodes) ToCategories() []model.Category {
	var cat []model.Category
	for _, n := range nodes {
		cat = append(cat, *n.Category)
	}
	return cat
}

func (n *Node) AddChild(nodes ...*Node) {
	for _, node := range nodes {
		n.children = append(n.children, node)
	}
}

func NewNode(c model.Category, tree *CategoryTree) *Node {
	n := &Node{
		Category: &c,
		children: make(Nodes, 0),
	}
	if tree != nil {
		tree.m[c.Id] = n
	}
	return n
}
