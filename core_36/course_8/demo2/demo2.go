package main

import (
	"container/ring"
	"fmt"
)

func main() {
	//initialize a ring
	r := ring.New(5)

	len := r.Len()
	fmt.Println("the length of ring is: ", len)

	//initialize the ring
	for i := 0; i < len; i++ {
		r.Value = i
		r = r.Next()
	}
	fmt.Println("ring r before move")
	//iterate through the ring and print its contents
	r.Do(func(i interface{}) {
		fmt.Println(i.(int))
	})
	//move the pointer forward by one step
	r = r.Move(1)
	fmt.Println("ring r after move")
	r.Do(func(i interface{}) {
		fmt.Println(i.(int))
	})
	//move the pointer back by two steps
	r = r.Move(-2)
	fmt.Println("ring r after move")
	r.Do(func(i interface{}) {
		fmt.Println(i.(int))
	})

	//unlink two elements from r, starting from r.Next()
	r.Unlink(2)
	fmt.Println("ring r after unlink")
	r.Do(func(i interface{}) {
		fmt.Println(i.(int))
	})

	s := ring.New(2)
	len1 := s.Len()

	for i := 0; i < len1; i++ {
		s.Value = i
		s = s.Next()
	}
	fmt.Println("ring: rs")
	//link ring r and ring s
	rs := r.Link(s)
	rs.Do(func(i interface{}) {
		fmt.Println(i.(int))
	})

}
