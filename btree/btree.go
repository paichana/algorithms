package btree

import (
	"fmt"

	"github.com/paichana/algorithms/queue"
	"github.com/paichana/algorithms/stack"
)

type node struct {
	value      interface{}
	leftChild  *node
	rightChild *node
}

func (n *node) GetChild() (*node, *node) {
	return n.leftChild, n.rightChild
}

func (n *node) GetValue() interface{} {
	return n.value
}

//New creates empty of a binary tree
func New(value interface{}) *node {
	return &node{
		value: value,
	}
}

//Add left child to the node
func (n *node) AddLeft(nd *node) {
	n.leftChild = nd
}

//Add right child to the node
func (n *node) AddRight(nd *node) {
	n.rightChild = nd
}

//Change value of current node
func (n *node) ChangeValue(val interface{}) {
	n.value = val
}

func DFSRecursive(n *node) {
	if n == nil {
		return
	}
	fmt.Println(n.value)
	if n.leftChild != nil {
		DFSRecursive(n.leftChild)
	}
	if n.rightChild != nil {
		DFSRecursive(n.rightChild)
	}
}

func DFS(n *node) {
	if n == nil {
		return
	}
	s := stack.New()
	s.Push(n)

	for s.Count() > 0 {
		current := s.Pop().(*node)
		fmt.Println(current.GetValue())
		left, right := current.GetChild()

		if right != nil {
			s.Push(right)
		}

		if left != nil {
			s.Push(left)
		}

	}

}

func BFS(n *node) {
	if n == nil {
		return
	}
	q := queue.New()
	q.Push(n)

	for q.Count() > 0 {
		current := q.Pop().(*node)
		fmt.Println(current.GetValue())
		left, right := current.GetChild()

		if left != nil {
			q.Push(left)
		}
		if right != nil {
			q.Push(right)
		}

	}

}
