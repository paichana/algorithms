package queue

import (
	"github.com/paichana/algorithms/linkedlist"
)

type queue struct {
	data  *linkedlist.LinkedList
	count int
}

func New() *queue {
	return &queue{
		data: linkedlist.New(),
	}
}

func (s *queue) Count() int {
	return s.count
}
func (s *queue) Push(val interface{}) {
	s.data.AddFront(val)
	s.count++
}

func (s *queue) Pop() interface{} {
	val := s.data.GetBack()
	s.data.RemoveBack()
	if s.count > 0 {
		s.count--
	}
	return val
}

func (s *queue) Print() {
	s.data.Print()
}
