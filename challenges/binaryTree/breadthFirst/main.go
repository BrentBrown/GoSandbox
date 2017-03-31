package main

import "fmt"

type Node struct {
	Left *Node
	Value rune
	Right *Node
}

type Queue struct {
	nodes []*Node
	head int
	tail int
	count int
}

func walk(n *Node) {

	if n == nil { return }
	q := &Queue{nodes: make([]*Node, 11)}
	q.push(n)

	for q.count != 0 {
		current := q.pop()
		fmt.Printf("%c", current.Value)
		if current.Left != nil { q.push(current.Left) }
		if current.Right != nil { q.push(current.Right) }
	}

}

func main() {
	root := createNode('F')
	root.Left = createNode('D')
	root.Left.Left = createNode('B')
	root.Left.Right = createNode('E')
	root.Left.Left.Left = createNode('A')
	root.Left.Left.Right = createNode('C')
	root.Right = createNode('J')
	root.Right.Left = createNode('G')
	root.Right.Right = createNode('K')
	root.Right.Left.Right = createNode('I')
	root.Right.Left.Right.Left = createNode('H')
	walk(root)
}

func (q *Queue) push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)*2)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

func (q *Queue) pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func createNode(value rune) *Node {
	node := new(Node)
	node.Left = nil
	node.Value = value
	node.Right = nil
	return node
}