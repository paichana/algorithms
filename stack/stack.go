package stack

import (
	"github.com/paichana/algorithms/linkedlist"
)

type stack struct {
	data  *linkedlist.LinkedList
	count int
}

func New() *stack {
	return &stack{
		data:  linkedlist.New(),
		count: 0,
	}
}

func (s *stack) Push(val interface{}) {
	s.data.AddFront(val)
	s.count++
}

func (s *stack) Pop() interface{} {
	val := s.data.GetFront()
	s.data.RemoveFront()
	if s.count > 0 {
		s.count--
	}
	return val
}

func (s *stack) Count() int {
	return s.count
}

func (s *stack) Print() {
	s.data.Print()
}
