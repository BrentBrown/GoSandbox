package main

import (
	"fmt"
)

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	nodes []*TreeNode
	count int
}

func (s *Stack) Push(n *TreeNode) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *TreeNode {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	stack := new(Stack)
	current := root

	for current != nil {
		stack.Push(current)
		current = current.Left
	}

	for stack.count > 0 {
		current = stack.Pop()
		fmt.Print(current.Data, "\t")
		if current.Right != nil {
			current = current.Right
			for current != nil {
				stack.Push(current)
				current = current.Left
			}
		}
	}
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

	inOrder(root)
}

func Create(value int) *TreeNode {
	node := new(TreeNode)
	node.Data = value
	node.Left = nil
	node.Right = nil
	return node
}
