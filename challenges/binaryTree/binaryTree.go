package binaryTree

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func Create(value int) *Node {
	node := new(Node)
	node.Data = value
	node.Left = nil
	node.Right = nil
	return node
}
