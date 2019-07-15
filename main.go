package main

import (
	"fmt"

	"github.com/paichana/algorithms/btree"
	"github.com/paichana/algorithms/linkedlist"
	"github.com/paichana/algorithms/queue"
)

func main() {
	n1 := btree.New("1")
	n2 := btree.New("2")
	n3 := btree.New("3")
	n4 := btree.New("4")
	n5 := btree.New("5")
	n6 := btree.New("6")
	n7 := btree.New("7")
	n8 := btree.New("8")

	n1.AddLeft(n2)
	n1.AddRight(n3)

	n2.AddLeft(n4)
	n2.AddRight(n5)

	n3.AddLeft(n6)
	n3.AddRight(n7)

	n4.AddLeft(n8)

	fmt.Println("DFS")
	btree.DFS(n1)
	fmt.Println("DFS recursive")
	btree.DFSRecursive(n1)
	fmt.Println("BFS")
	btree.BFS(n1)

	q := queue.New()
	q.Push("1")
	q.Push("2")
	q.Push("3")
	q.Push("4")
	q.Push("5")
	q.Push("6")

	fmt.Println(q.Pop().(string))
	fmt.Println(q.Pop().(string))
	fmt.Println(q.Pop().(string))
	q.Push("7")
	q.Print()

	ll := linkedlist.New()
	ll.AddBack("1")
	ll.AddBack("2")
	ll.AddFront("0")
	ll.AddFront("-1")

	ll.Print()

	fmt.Println(ll.GetBack().(string))

	fmt.Println(ll.GetFront().(string))

}
