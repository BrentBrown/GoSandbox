package main

import "fmt"

//bt not bst
//assumes no duplicates and that search values exist in tree

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
	if root == nil { return nil }
	if root.Value == v1 || root.Value == v2 { return root }
	left := lca(root.Left, v1, v2)
	right := lca(root.Right, v1, v2)
	if left != nil && right != nil { return root }
	if left == nil { return right } else { return left }
}

func main() {
	root := createNode(3)
	root.Left = createNode(6)
	root.Left.Left = createNode(2)
	root.Left.Right = createNode(11)
	root.Left.Right.Left = createNode(9)
	root.Left.Right.Right = createNode(5)
	root.Right = createNode(8)
	root.Right.Right = createNode(13)
	root.Right.Right.Left = createNode(7)
	lca := lca(root, 11, 2)
	fmt.Print(lca.Value)
}

/*
          3
         / \
        6   8
       / \   \
      2  11  13
         /\  /
        9 5 7
*/