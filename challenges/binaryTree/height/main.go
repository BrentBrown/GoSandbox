package main

import (
	"math"
	"fmt"
)

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

func height(root *Node) float64 {
	if root == nil { return -1 }
	return math.Max(height(root.Left), height(root.Right)) + 1
}

func main() {
	root := CreateNode(42)
	root.Left = CreateNode(23)
	root.Left.Left = CreateNode(12)
	root.Left.Left.Left = CreateNode(9)
	root.Right = CreateNode(44)
	root.Right.Left = CreateNode(43)
	root.Right.Right = CreateNode(50)
	fmt.Print(height(root))
}
