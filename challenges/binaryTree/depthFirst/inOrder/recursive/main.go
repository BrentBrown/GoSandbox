package main

import (
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func InOrder(node *Node) {
	if node == nil {
		return
	}
	InOrder(node.Left)
	fmt.Print(node.Data, "\t")
	InOrder(node.Right)
}

func main() {
	root := Create(42)
	root.Left = Create(6)
	root.Left.Left = Create(12)
	root.Left.Left.Left = Create(-3)
	root.Left.Left.Right = Create(20)

	root.Left.Right = Create(4)
	root.Right = Create(10)
	root.Right.Left = Create(20)
	root.Right.Right = Create(-7)

	InOrder(root)
}

func Create(value int) *Node {
	node := new(Node)
	node.Data = value
	node.Left = nil
	node.Right = nil
	return node
}
