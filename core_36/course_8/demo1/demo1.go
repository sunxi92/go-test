package main

import (
	"container/list"
	"fmt"
)

func main() {
	//initialize a list
	l := list.New()
	fmt.Println(l)

	//add items to the list
	l.PushFront(10)
	fmt.Println("the front element: ", l.Front())
	l.PushBack(11)
	fmt.Println("the back element: ", l.Back())
	l.PushBack(12)
	l.PushBack(13)
	fmt.Println("the back element: ", l.Back())

	//traverse the list
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("Traverse the list in reverse order")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	//remove items from the list
	v := l.Back()
	l.Remove(v)
	fmt.Println("the back element: ", l.Back())

	//insert item
	v1 := l.Front()
	v2 := l.InsertBefore(8, v1)
	fmt.Println("the front element: ", v2.Value)

	v3 := l.InsertAfter(9, v2)
	fmt.Println("the second element: ", v3.Value)

	fmt.Println("the length of the list: ", l.Len())

	//move the item
	fmt.Println("before move items")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	back := l.Back()
	l.MoveAfter(v2, back)
	l.MoveBefore(v3, back)
	fmt.Println("after move item")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	l.MoveToFront(v2)
	l.MoveToBack(v3)
	fmt.Println("after move item")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	l1 := list.New()
	l1.PushFront(1)
	l1.PushBack(2)

	//append another list
	l.PushBackList(l1)
	fmt.Println("apppend")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	l.PushFrontList(l1)
	fmt.Println("apppend")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
