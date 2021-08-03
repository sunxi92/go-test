package main

import "fmt"

type Employee struct {
	Name string
	Age  int
}

func main() {
	ch1 := make(chan []int, 1)
	s1 := []int{1, 2, 3}
	ch1 <- s1
	s2 := <-ch1
	s1[0] = 100
	fmt.Println(s1, s2) //[100 2 3] [100 2 3]
	s2[1] = 100
	fmt.Println(s1, s2) //[100 100 3] [100 100 3]

	ch2 := make(chan [3]int, 1)
	s3 := [3]int{1, 2, 3}
	ch2 <- s3
	s4 := <-ch2
	s4[0] = 100
	fmt.Println(s3, s4) //[100 2 3] [1 2 3]

	employee1 := Employee{"sunxi", 18}
	ch3 := make(chan Employee, 1)
	ch3 <- employee1
	employee2 := <-ch3
	fmt.Printf("%p\n", &employee1)
	fmt.Printf("%p\n", &employee2)

	ch4 := make(chan *int, 2)
	i := 1
	fmt.Println("i: ", &i, i)
	ch4 <- &i
	p := <-ch4
	fmt.Println("p: ", p, *p)
	i = 2
	fmt.Println("i: ", &i, i)
}
