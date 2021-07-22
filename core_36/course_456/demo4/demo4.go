package main

import "fmt"

var block = "1"

func main() {
	block := "2"
	{
		block := "3"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
}
