package util

import (
	"errors"
)

func treeMain() {
}

type TreeNode struct {
	Content  interface{}
	Parent   *TreeNode
	Children []*TreeNode
}

//add a node and return the added node
func (c *TreeNode) AddChild(nc *TreeNode) *TreeNode {
	c.Children = append(c.Children, nc)
	nc.Parent = c
	return nc
}

//the depth from the current node to root node
func (c *TreeNode) Depth(root *TreeNode) (int, error) {
	counter := 0
	curNode := c
	for curNode != nil {
		if curNode == root {
			return counter, nil
		} else {
			counter++
			curNode = curNode.Parent
		}
	}
	return 0, errors.New("wrong root node")
}

//limited functionality, walk over the tree, apply functions to all the nodes
func (c *TreeNode) ForEachNode(f func(*TreeNode)) {
	for _, k := range c.Children {
		f(k)
		k.ForEachNode(f)
	}
}
