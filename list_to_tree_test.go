package tree

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Node struct {
	ID       uint
	PID      uint
	Children []*Node
	Name     string
}

func (n *Node) GetId() uint {
	return n.ID
}

func (n *Node) GetPid() uint {
	return n.PID
}

func (n *Node) SetChildren(arr []*Node) {
	n.Children = arr
}

func TestListToTree(t *testing.T) {
	list := []*Node{
		{
			ID:   1,
			PID:  0,
			Name: "A",
		},
		{
			ID:       2,
			PID:      0,
			Name:     "B",
			Children: make([]*Node, 0), // if you want the empty slice not a nil as a result
		},
		{
			ID:   3,
			PID:  1,
			Name: "A-1",
		},
		{
			ID:   4,
			PID:  1,
			Name: "A-2",
		},
		{
			ID:   5,
			PID:  3,
			Name: "A-1-1",
		},
		{
			ID:   6,
			PID:  3,
			Name: "A-1-2",
		},
		{
			ID:   7,
			PID:  100, // not a node's id is 100, so this node will be ignored in result tree.
			Name: "A-1-2",
		},
	}

	tree := ListToTree(list)

	fmt.Println("tree ", tree)
	bytes, err := json.Marshal(tree)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
}
