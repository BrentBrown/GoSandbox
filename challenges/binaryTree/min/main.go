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
	node.Right = nil
	node.Value = value
	return node
}

func min(root *Node) {
	if root.Left != nil {
		min(root.Left)
	} else {
		fmt.Print(root.Value)
		return
	}
}

func main() {
	root := CreateNode(42)
	root.Left = CreateNode(23)
	root.Right = CreateNode(44)
	root.Left.Left = CreateNode(12)
	min(root)
}


