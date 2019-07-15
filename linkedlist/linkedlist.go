package linkedlist

import "fmt"

type node struct {
	next  *node
	value interface{}
}

type LinkedList struct {
	head *node
	tail *node
}

func New() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) AddBack(val interface{}) {
	if ll.head == nil {
		ll.head = &node{
			value: val,
		}
		return
	}
	cn := ll.head
	for cn.next != nil {
		cn = cn.next
	}
	cn.next = &node{
		value: val,
	}
}

func (ll *LinkedList) AddFront(val interface{}) {
	if ll.head == nil {
		ll.head = &node{
			value: val,
		}
		return
	}

	head := &node{
		value: val,
		next:  ll.head,
	}
	ll.head = head
}

func (ll *LinkedList) RemoveBack() {
	if ll.head == nil {
		return
	} else if ll.head.next == nil {
		ll.head = nil
		return
	}

	cn := ll.head

	for cn.next.next != nil {
		cn = cn.next
	}
	cn.next = nil

}

func (ll *LinkedList) RemoveFront() {
	if ll.head == nil {
		return
	}
	ll.head = ll.head.next
}

func (ll *LinkedList) GetFront() interface{} {
	if ll.head != nil {
		return ll.head.value
	}
	return ll.head.value
}

func (ll *LinkedList) GetBack() interface{} {
	if ll.head == nil {
		return ll.head.value
	}

	cn := ll.head

	for cn.next != nil {
		cn = cn.next
	}
	return cn.value
}

func (ll *LinkedList) Print() {
	cn := ll.head
	for cn != nil {
		fmt.Print(cn.value, " ")
		cn = cn.next
	}
	fmt.Println()
}
