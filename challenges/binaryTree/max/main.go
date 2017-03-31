package main

import "fmt"

type Node struct {
	Left *Node
	Value int
	Right *Node
}

func CreateNode(value int) *Node {
	node := new(Node)
	node.Left = nil
	node.Value = value
	node.Right = nil
	return node
}

func max(root *Node) {
	if root.Right != nil {
		max(root.Right)
	} else {
		fmt.Print(root.Value)
		return
	}
}

func main() {
	root := CreateNode(42)
	root.Left = CreateNode(23)
	root.Right = CreateNode(44)
	root.Right.Right = CreateNode(50)
	root.Left.Left = CreateNode(12)
	max(root)
}
