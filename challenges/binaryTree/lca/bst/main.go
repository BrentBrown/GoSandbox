package main

import "fmt"

type Node struct {
	Left *Node
	Value int
	Right *Node
}

func createNode(value int) *Node {
	node := new(Node)
	node.Left = nil
	node.Value = value
	node.Right = nil
	return node
}

func lca(root *Node, v1 int, v2 int) *Node {
	if v1 < root.Value && v2 < root.Value { return lca(root.Left, v1, v2) }
	if v1 > root.Value && v2 > root.Value { return lca(root.Right, v1, v2) }
	return root
}

func main() {
	root := createNode(10)
	root.Left = createNode(-10)
	root.Left.Right = createNode(8)
	root.Left.Right.Left = createNode(6)
	root.Left.Right.Right = createNode(9)
	root.Right = createNode(30)
	root.Right.Left = createNode(25)
	root.Right.Right = createNode(60)
	root.Right.Right = createNode(78)

	lca := lca(root, 6, 9)
	fmt.Print(lca.Value)

}

/*
              10
             /  \
           -10   30
             \   / \
             8  25  60
            / \   \   \
           6   9   28  78
 */