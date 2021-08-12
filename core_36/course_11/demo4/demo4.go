package main

import (
	"fmt"
	"time"
)

func main() {
	example1()
	example2()
}

func example1() {
	ch1 := make(chan int, 1)
	time.AfterFunc(time.Second, func() { close(ch1) })
loop:
	for {
		select {
		case _, ok := <-ch1:
			if !ok {
				break loop
			}
			fmt.Println("in")
		}
	}
	fmt.Println("out")
}

func example2() {
	ch1 := make(chan int, 1)
	time.AfterFunc(time.Second, func() { close(ch1) })

	for {
		select {
		case _, ok := <-ch1:
			if !ok {
				goto loop
			}
			fmt.Println("in")
		}
	}
loop:
	fmt.Println("out")
}
