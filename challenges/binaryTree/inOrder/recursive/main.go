package main

import (
	"fmt"
	"github.com/BrentBrown/GoSandbox/challenges/binaryTree"
)

func inOrder(node *binaryTree.Node) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	fmt.Print(node.Data, "\t")
	inOrder(node.Right)
}

func main() {
	root := binaryTree.Create(42)
	root.Left = binaryTree.Create(6)
	root.Left.Left = binaryTree.Create(12)
	root.Left.Left.Left = binaryTree.Create(-3)
	root.Left.Left.Right = binaryTree.Create(20)

	root.Left.Right = binaryTree.Create(4)
	root.Right = binaryTree.Create(10)
	root.Right.Left = binaryTree.Create(20)
	root.Right.Right = binaryTree.Create(-7)

	inOrder(root)
}
